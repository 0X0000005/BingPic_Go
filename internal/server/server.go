package server

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"WallpaperManager/internal/config"
)

type Server struct {
	Config *config.Config
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		Config: cfg,
	}
}

func (s *Server) Start() error {
	fs := http.FileServer(http.Dir(s.Config.Download.Path))
	http.Handle("/images/", http.StripPrefix("/images/", fs))

	http.HandleFunc("/", s.handleIndex)
	http.HandleFunc("/api/images", s.handleListImages)

	addr := fmt.Sprintf(":%d", s.Config.Server.Port)
	fmt.Printf("启动 Web 服务器: http://localhost%s\n", addr)
	return http.ListenAndServe(addr, nil)
}

func (s *Server) handleIndex(w http.ResponseWriter, r *http.Request) {
	// 加载首页模板
	// 这里依赖 'web/templates/index.html' 文件
	tmplPath := filepath.Join("web", "templates", "index.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "加载模板失败", http.StatusInternalServerError)
		fmt.Printf("模板错误: %v\n", err)
		return
	}

	// 渲染页面
	tmpl.Execute(w, nil)
}

func (s *Server) handleListImages(w http.ResponseWriter, r *http.Request) {
	var images []string

	files, err := os.ReadDir(s.Config.Download.Path)
	if err != nil {
		http.Error(w, "读取图片目录失败", http.StatusInternalServerError)
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
