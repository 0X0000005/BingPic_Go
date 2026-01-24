package config

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	// 创建临时配置文件
	content := []byte(`
server:
  port: 9090
download:
  path: "./test_wallpapers"
  bing:
    enabled: true
    schedule: "0 10 * * *"
  spotlight:
    enabled: false
    schedule: "0 */1 * * *"
`)
	tmpfile, err := os.CreateTemp("", "config_*.yaml")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name()) // 清理

	if _, err := tmpfile.Write(content); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	// 测试加载
	cfg, err := LoadConfig(tmpfile.Name())
	if err != nil {
		t.Fatalf("LoadConfig 失败: %v", err)
	}

	if cfg.Server.Port != 9090 {
		t.Errorf("预期端口 9090, 实际 %d", cfg.Server.Port)
	}
	if cfg.Download.Path != "./test_wallpapers" {
		t.Errorf("预期路径 ./test_wallpapers, 实际 %s", cfg.Download.Path)
	}
	if !cfg.Download.Bing.Enabled {
		t.Error("预期 Bing 启用")
	}
	if cfg.Download.Spotlight.Enabled {
		t.Error("预期 Spotlight 禁用")
	}
}
