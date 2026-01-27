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
	mode := flag.String("mode", "web", "运行模式: web 或 cli")
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

	// 2. 初始化调度器 (Web 模式)
	sched := scheduler.NewScheduler()

	// 3. 注册 Bing 任务
	if cfg.Download.Bing.Enabled {
		// Web 模式下不再自动立即下载，只注册定时任务
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

	// 4. 启动 Web 服务器
	srv := server.NewServer(cfg, *configPath)
	if err := srv.Start(); err != nil {
		log.Fatalf("Web 服务器失败: %v", err)
	}
}
