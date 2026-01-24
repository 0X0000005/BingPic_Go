package downloader

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const (
	BingAPIURL  = "https://www.bing.com/HPImageArchive.aspx?format=js&idx=0&n=1&mkt=zh-CN"
	BingBaseURL = "https://www.bing.com"
)

type BingDownloader struct {
	DownloadPath string
}

type BingResponse struct {
	Images []struct {
		Startdate     string `json:"startdate"`
		Fullstartdate string `json:"fullstartdate"`
		Enddate       string `json:"enddate"`
		Url           string `json:"url"`
		Urlbase       string `json:"urlbase"`
		Copyright     string `json:"copyright"`
		Title         string `json:"title"`
	} `json:"images"`
}

func NewBingDownloader(downloadPath string) *BingDownloader {
	return &BingDownloader{
		DownloadPath: downloadPath,
	}
}

func (d *BingDownloader) DownloadTodayWallpaper() error {
	// 确保下载目录存在
	if err := os.MkdirAll(d.DownloadPath, 0755); err != nil {
		return fmt.Errorf("创建下载目录失败: %w", err)
	}

	// 获取 Bing 数据
	resp, err := http.Get(BingAPIURL)
	if err != nil {
		return fmt.Errorf("请求 Bing API 失败: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Bing API 返回状态码非 200: %d", resp.StatusCode)
	}

	var bingResp BingResponse
	if err := json.NewDecoder(resp.Body).Decode(&bingResp); err != nil {
		return fmt.Errorf("解析 JSON 失败: %w", err)
	}

	if len(bingResp.Images) == 0 {
		return fmt.Errorf("Bing 响应中未找到图片数据")
	}

	imgData := bingResp.Images[0]
	imageURL := BingBaseURL + imgData.Url

	// 文件名格式: Bing_YYYYMMDD.jpg
	dateStr := imgData.Enddate // YYYYMMDD
	filename := fmt.Sprintf("Bing_%s.jpg", dateStr)

	filePath := filepath.Join(d.DownloadPath, filename)

	// 检查文件是否存在
	if _, err := os.Stat(filePath); err == nil {
		fmt.Printf("文件已存在，跳过: %s\n", filePath)
		return nil
	}

	// 下载图片
	imgResp, err := http.Get(imageURL)
	if err != nil {
		return fmt.Errorf("下载图片失败: %w", err)
	}
	defer imgResp.Body.Close()

	if imgResp.StatusCode != http.StatusOK {
		return fmt.Errorf("图片下载请求返回状态码非 200: %d", imgResp.StatusCode)
	}

	out, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("创建文件失败: %w", err)
	}
	defer out.Close()

	_, err = io.Copy(out, imgResp.Body)
	if err != nil {
		return fmt.Errorf("保存文件内容失败: %w", err)
	}

	fmt.Printf("Bing 壁纸下载成功: %s\n", filePath)
	return nil
}
