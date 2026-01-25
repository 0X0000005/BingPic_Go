# 壁纸管理工具 (Wallpaper Manager)

一个 Go 语言编写的跨平台应用，用于管理 Bing 和 Windows 聚焦 (Spotlight) 壁纸。

## 功能
- **Bing 壁纸**: 自动下载每日 Bing 壁纸。
- **Windows 聚焦**: 提取 Windows 聚焦图片 (仅限 Windows)。
- **Web 界面**: 在画廊中查看已下载的壁纸。
- **定时任务**: 可配置的定时计划。
- **Docker 支持**: 支持容器化部署。
- **跨平台**: 支持 Windows 和 Linux。

## 使用方法

### 配置
编辑 `config.yaml`:
```yaml
server:
  port: 8080
download:
  path: "./wallpapers"
  bing:
    enabled: true
    schedule: "0 10 * * *"
  spotlight:
    enabled: true
    schedule: "0 */1 * * *"
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
build.bat all
```

**Linux/Mac (bash):**
```bash
# 构建所有支持的平台
./build.sh all
```
输出位于 `dist/` 目录。
输出位于 `dist/` 目录。

### Docker
```bash
docker-compose up -d
```

## 开发
- **测试**: `go test ./...`