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
	mode := flag.String("mode", "cli", "运行模式: cli (默认) 或 server (API 服务)")
	flag.Parse()

	// 1. 加载配置
	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	// 初始化 Bing 下载器
	bingDL := downloader.NewBingDownloader(cfg.Download.Path)

	if *mode == "cli" {
		log.Println("以 CLI 模式运行...")
		if errors := bingDL.DownloadRecentWallpapers(); errors != nil {
			log.Fatalf("Bing 壁纸下载失败: %v", errors)
		}
		log.Println("Bing 壁纸下载完成")
		return
	}

	if *mode == "server" {
		// 2. 初始化调度器 (Server 模式)
		sched := scheduler.NewScheduler()

		// 3. 注册 Bing 任务
		if cfg.Download.Bing.Enabled {
			// 立即执行一次下载 (异步执行，避免阻塞启动)
			go func() {
				log.Println("正在执行启动时 Bing 壁纸下载...")
				if err := bingDL.DownloadRecentWallpapers(); err != nil {
					log.Printf("启动时 Bing 壁纸下载失败: %v", err)
				}
			}()

			// 注册定时任务
			err := sched.AddJob(cfg.Download.Bing.Schedule, func() {
				log.Println("执行 Bing 壁纸下载任务...")
				if err := bingDL.DownloadRecentWallpapers(); err != nil {
					log.Printf("Bing 壁纸下载失败: %v", err)
				}
			})
			if err != nil {
				log.Fatalf("调度 Bing 任务失败: %v", err)
			}
		}

		// 启动调度器
		sched.Start()
		defer sched.Stop()

		// 4. 启动 API 服务器
		srv := server.NewServer(cfg, *configPath)
		if err := srv.Start(); err != nil {
			log.Fatalf("API 服务器失败: %v", err)
		}
	} else {
		log.Fatalf("未知模式: %s. 请使用 'cli' 或 'server'", *mode)
	}
}
