package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"WallpaperManager/internal/config"
	"WallpaperManager/internal/downloader"
)

type Server struct {
	Config     *config.Config
	ConfigPath string
}

func NewServer(cfg *config.Config, configPath string) *Server {
	return &Server{
		Config:     cfg,
		ConfigPath: configPath,
	}
}

func (s *Server) Start() error {
	// API Endpoints
	http.HandleFunc("/api/images", s.handleListImages)
	http.HandleFunc("/api/config", s.handleConfig)
	http.HandleFunc("/api/download", s.handleDownload)

	// Serve downloaded images
	fsImages := http.FileServer(http.Dir(s.Config.Download.Path))
	http.Handle("/images/", http.StripPrefix("/images/", fsImages))

	addr := fmt.Sprintf(":%d", s.Config.Server.Port)
	fmt.Printf("Starting API Server: http://localhost%s\n", addr)
	return http.ListenAndServe(addr, nil)
}

func (s *Server) handleDownload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Re-initialize downloader with current config path
	dl := downloader.NewBingDownloader(s.Config.Download.Path)

	if err := dl.DownloadRecentWallpapers(); err != nil {
		http.Error(w, fmt.Sprintf("Download failed: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Download started and completed."))
}

func (s *Server) handleConfig(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(s.Config)
	case http.MethodPost:
		var newConfig config.Config
		if err := json.NewDecoder(r.Body).Decode(&newConfig); err != nil {
			http.Error(w, "Invalid JSON config", http.StatusBadRequest)
			return
		}

		// Update config
		s.Config.Server = newConfig.Server
		s.Config.Download = newConfig.Download

		// Save to file
		if err := config.SaveConfig(s.ConfigPath, s.Config); err != nil {
			http.Error(w, "Failed to save config", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Config updated"))
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (s *Server) handleListImages(w http.ResponseWriter, r *http.Request) {
	var images []string

	// Ensure download directory exists
	if _, err := os.Stat(s.Config.Download.Path); os.IsNotExist(err) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(images)
		return
	}

	files, err := os.ReadDir(s.Config.Download.Path)
	if err != nil {
		http.Error(w, "Failed to read image directory", http.StatusInternalServerError)
		return
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		ext := strings.ToLower(filepath.Ext(file.Name()))
		if ext == ".jpg" || ext == ".png" || ext == ".jpeg" {
			images = append(images, "/images/"+file.Name())
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(images)
}
