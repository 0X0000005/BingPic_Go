# 壁纸管理工具 (Wallpaper Manager)

一个 Go 语言编写的跨平台应用，用于管理 Bing 壁纸。

## 功能
- **Bing 壁纸**: 自动下载每日 Bing 壁纸。
- **Web 界面**: 在画廊中查看已下载的壁纸。
- **定时任务**: 可配置的定时计划。
- **CLI/Web 模式**: 支持命令行一次性运行或 Web 服务常驻运行。
- **Docker 支持**: 支持容器化部署。
- **跨平台**: 支持 Windows 和 Linux。

## 使用说明

### 运行模式

程序支持两种运行模式：

1.  **Web 模式 (默认)**
    启动 Web 服务器和定时任务调度器。
    ```bash
    ./WM -mode web
    # 或直接运行
    ./WM
    ```

2.  **CLI 模式**
    立即执行一次下载任务并退出。
    ```bash
    ./WM -mode cli
    ```

### Web 界面
启动后访问 `http://localhost:8080` (默认端口) 查看壁纸画廊和进行设置。

### 配置
配置文件 `config.yaml` 支持通过 Web 界面进行修改。
```yaml
server:
  port: 8080
download:
  path: "./wallpapers"
  bing:
    enabled: true
    schedule: "0 10 * * *"
```

### 运行
```bash
go run cmd/server/main.go
```
访问 `http://localhost:8080` 查看画廊。

### 构建
使用提供的构建脚本:

**Windows (cmd/powershell):**
```bat
build.bat
```

**Linux/Mac (bash):**
```bash
./build.sh
```
输出位于项目根目录。

### Docker
```bash
docker-compose up -d
```

## 开发
- **测试**: `go test ./...`