//go:build windows

package downloader

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"os"
	"path/filepath"
)

type SpotlightExtractor struct {
	DownloadPath string
}

func NewSpotlightExtractor(downloadPath string) *SpotlightExtractor {
	return &SpotlightExtractor{
		DownloadPath: downloadPath,
	}
}

func (s *SpotlightExtractor) Extract() error {
	// Spotlight 资源路径
	localAppData := os.Getenv("LocalStorage") // 通常由构建环境传递或推导
	if localAppData == "" {
		localAppData = os.Getenv("LOCALAPPDATA")
	}

	spotlightPath := filepath.Join(localAppData, "Packages", "Microsoft.Windows.ContentDeliveryManager_cw5n1h2txyewy", "LocalState", "Assets")

	if _, err := os.Stat(spotlightPath); os.IsNotExist(err) {
		return fmt.Errorf("Spotlight 目录未找到: %s", spotlightPath)
	}

	if err := os.MkdirAll(s.DownloadPath, 0755); err != nil {
		return fmt.Errorf("创建下载目录失败: %w", err)
	}

	files, err := os.ReadDir(spotlightPath)
	if err != nil {
		return fmt.Errorf("读取 Spotlight 目录失败: %w", err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		srcPath := filepath.Join(spotlightPath, file.Name())
		info, err := file.Info()
		if err != nil {
			continue
		}

		// 过滤小文件 (例如 < 100KB)
		if info.Size() < 100*1024 {
			continue
		}

		// 仅提取横屏图片
		if s.isLandscapeImage(srcPath) {
			destName := fmt.Sprintf("Spotlight_%s.jpg", file.Name())
			destPath := filepath.Join(s.DownloadPath, destName)

			if _, err := os.Stat(destPath); err == nil {
				continue // 文件已存在
			}

			if err := copyFile(srcPath, destPath); err != nil {
				fmt.Printf("复制 Spotlight 图片 %s 失败: %v\n", file.Name(), err)
			} else {
				fmt.Printf("提取 Spotlight 图片成功: %s\n", destName)
			}
		}
	}

	return nil
}

func (s *SpotlightExtractor) isLandscapeImage(path string) bool {
	f, err := os.Open(path)
	if err != nil {
		return false
	}
	defer f.Close()

	cfg, _, err := image.DecodeConfig(f)
	if err != nil {
		return false // 不是有效图片
	}

	return cfg.Width > cfg.Height
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	return err
}
