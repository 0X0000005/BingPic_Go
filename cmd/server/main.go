package main

import (
	"flag"
	"log"

	"WallpaperManager/internal/config"
	"WallpaperManager/internal/downloader"
	"WallpaperManager/internal/scheduler"
	"WallpaperManager/internal/server"
)

func main() {
	configPath := flag.String("config", "config.yaml", "配置文件路径")
	flag.Parse()

	// 1. 加载配置
	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	// 2. 初始化调度器
	sched := scheduler.NewScheduler()

	// 3. 注册 Bing 任务
	if cfg.Download.Bing.Enabled {
		bingDL := downloader.NewBingDownloader(cfg.Download.Path)
		// 启动时立即运行一次
		go func() {
			if err := bingDL.DownloadTodayWallpaper(); err != nil {
				log.Printf("初始 Bing 壁纸下载失败: %v", err)
			}
		}()

		err := sched.AddJob(cfg.Download.Bing.Schedule, func() {
			log.Println("执行 Bing 壁纸下载任务...")
			if err := bingDL.DownloadTodayWallpaper(); err != nil {
				log.Printf("Bing 壁纸下载失败: %v", err)
			}
		})
		if err != nil {
			log.Fatalf("调度 Bing 任务失败: %v", err)
		}
	}

	// 4. 注册 Spotlight 任务
	if cfg.Download.Spotlight.Enabled {
		spotlightDL := downloader.NewSpotlightExtractor(cfg.Download.Path)
		// 启动时立即运行一次
		go func() {
			if err := spotlightDL.Extract(); err != nil {
				log.Printf("初始 Spotlight 提取失败: %v", err)
			}
		}()

		err := sched.AddJob(cfg.Download.Spotlight.Schedule, func() {
			log.Println("执行 Spotlight 提取任务...")
			if err := spotlightDL.Extract(); err != nil {
				log.Printf("Spotlight 提取失败: %v", err)
			}
		})
		if err != nil {
			log.Fatalf("调度 Spotlight 任务失败: %v", err)
		}
	}

	// 启动调度器
	sched.Start()
	defer sched.Stop()

	// 5. 启动 Web 服务器
	srv := server.NewServer(cfg)
	if err := srv.Start(); err != nil {
		log.Fatalf("Web 服务器失败: %v", err)
	}
}
