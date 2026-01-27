package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"WallpaperManager/internal/config"
)

func TestHandleListImages(t *testing.T) {
	// 设置测试目录
	tmpDir, err := os.MkdirTemp("", "images_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	// 创建虚拟图片
	os.Create(filepath.Join(tmpDir, "test1.jpg"))
	os.Create(filepath.Join(tmpDir, "test2.png"))
	os.Create(filepath.Join(tmpDir, "ignore.txt"))

	cfg := &config.Config{
		Download: config.DownloadConfig{
			Path: tmpDir,
		},
	}
	srv := NewServer(cfg, "")

	req, _ := http.NewRequest("GET", "/api/images", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(srv.handleListImages)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("处理程序返回错误状态码: 实际 %v 预期 %v",
			status, http.StatusOK)
	}

	var images []string
	err = json.Unmarshal(rr.Body.Bytes(), &images)
	if err != nil {
		t.Fatalf("解析 JSON 失败: %v", err)
	}

	if len(images) != 2 {
		t.Errorf("预期 2 张图片, 实际 %d", len(images))
	}
}
