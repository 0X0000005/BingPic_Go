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
	"WallpaperManager/web"
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
	// 挂载静态资源 (如果有的话，目前只有 templates，但为了将来扩展 supported)
	// 如果 static 文件夹为空，embed 可能会报错或忽略，这里假设 static 存在或仅使用 templates
	// 为了安全起见，我们仅处理 templates

	// 处理 API
	http.HandleFunc("/api/images", s.handleListImages)
	http.HandleFunc("/api/config", s.handleConfig)

	// 处理图片文件 (动态文件，仍在磁盘上)
	fsImages := http.FileServer(http.Dir(s.Config.Download.Path))
	http.Handle("/images/", http.StripPrefix("/images/", fsImages))

	// 处理首页
	http.HandleFunc("/", s.handleIndex)

	addr := fmt.Sprintf(":%d", s.Config.Server.Port)
	fmt.Printf("启动 Web 服务器: http://localhost%s\n", addr)
	return http.ListenAndServe(addr, nil)
}

func (s *Server) handleIndex(w http.ResponseWriter, r *http.Request) {
	// 从嵌入式 FS 中读取模板
	tmpl, err := template.ParseFS(web.Content, "templates/index.html")
	if err != nil {
		http.Error(w, "加载模板失败", http.StatusInternalServerError)
		fmt.Printf("模板错误: %v\n", err)
		return
	}

	// 渲染页面
	tmpl.Execute(w, nil)
}

func (s *Server) handleConfig(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(s.Config)
	case http.MethodPost:
		var newConfig config.Config
		if err := json.NewDecoder(r.Body).Decode(&newConfig); err != nil {
			http.Error(w, "无效的 JSON 配置", http.StatusBadRequest)
			return
		}

		// Update config
		s.Config.Server = newConfig.Server
		s.Config.Download = newConfig.Download

		// Save to file
		if err := config.SaveConfig(s.ConfigPath, s.Config); err != nil {
			http.Error(w, "保存配置失败", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("配置已更新"))
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (s *Server) handleListImages(w http.ResponseWriter, r *http.Request) {
	var images []string

	// 确保下载目录存在，避免读取报错
	if _, err := os.Stat(s.Config.Download.Path); os.IsNotExist(err) {
		// 还没下载图片，返回空列表
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(images)
		return
	}

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
