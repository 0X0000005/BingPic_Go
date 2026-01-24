//go:build !windows

package downloader

import "fmt"

type SpotlightExtractor struct {
	DownloadPath string
}

func NewSpotlightExtractor(downloadPath string) *SpotlightExtractor {
	return &SpotlightExtractor{
		DownloadPath: downloadPath,
	}
}

func (s *SpotlightExtractor) Extract() error {
	fmt.Println("Spotlight 提取仅支持 Windows 系统。")
	return nil
}
