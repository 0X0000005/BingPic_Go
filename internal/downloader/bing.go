package downloader

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

const (
	BingAPIURLTemplate = "https://www.bing.com/HPImageArchive.aspx?format=js&idx=%d&n=8&mkt=zh-CN"
	BingBaseURL        = "https://www.bing.com"
)

type BingDownloader struct {
	DownloadPath string
	client       *http.Client
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
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (d *BingDownloader) DownloadRecentWallpapers() error {
	// 确保下载目录存在
	if err := os.MkdirAll(d.DownloadPath, 0755); err != nil {
		return fmt.Errorf("创建下载目录失败: %w", err)
	}

	// Fetch 2 batches: idx=0 (days 0-7) and idx=8 (days 8-15)
	for idx := 0; idx <= 8; idx += 8 {
		if err := d.downloadBatch(idx); err != nil {
			fmt.Printf("Batch download failed for idx=%d: %v\n", idx, err)
			continue
		}
	}

	return nil
}

func (d *BingDownloader) downloadBatch(idx int) error {
	url := fmt.Sprintf(BingAPIURLTemplate, idx)

	// 获取 Bing 数据
	resp, err := d.client.Get(url)
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
		return nil // No more images
	}

	for _, imgData := range bingResp.Images {
		imageURL := BingBaseURL + imgData.Url

		// 文件名格式: YYYYMMDD.jpg
		dateStr := imgData.Enddate // YYYYMMDD
		filename := fmt.Sprintf("%s.jpg", dateStr)

		filePath := filepath.Join(d.DownloadPath, filename)

		// 检查文件是否存在
		if _, err := os.Stat(filePath); err == nil {
			fmt.Printf("文件已存在，跳过: %s\n", filePath)
			continue
		}

		// Download image
		fmt.Printf("正在下载: %s\n", filename)
		imgResp, err := d.client.Get(imageURL)
		if err != nil {
			fmt.Printf("下载图片失败 [%s]: %v\n", filename, err)
			continue
		}

		if imgResp.StatusCode != http.StatusOK {
			fmt.Printf("图片下载请求返回状态码非 200 [%s]: %d\n", filename, imgResp.StatusCode)
			imgResp.Body.Close()
			continue
		}

		out, err := os.Create(filePath)
		if err != nil {
			fmt.Printf("创建文件失败 [%s]: %v\n", filename, err)
			imgResp.Body.Close()
			continue
		}

		_, err = io.Copy(out, imgResp.Body)
		out.Close()
		imgResp.Body.Close()

		if err != nil {
			fmt.Printf("保存文件内容失败 [%s]: %v\n", filename, err)
		} else {
			fmt.Printf("Bing 壁纸下载成功: %s\n", filePath)
		}
	}
	return nil
}
