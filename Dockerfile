# 构建阶段
FROM golang:1.20-alpine AS builder

WORKDIR /app

# 安装 git 和 upx
RUN apk add --no-cache git upx

COPY . .

# 构建 (直接使用 go build 以简化 Dockerfile)
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o WallpaperManager cmd/server/main.go
# 可选 UPX
# RUN upx WallpaperManager

# 最终阶段
FROM alpine:latest

WORKDIR /app

# 安装 CA 证书 -> HTTPS 需要
RUN apk --no-cache add ca-certificates tzdata

COPY --from=builder /app/WallpaperManager .
COPY --from=builder /app/config.yaml .
COPY --from=builder /app/web ./web

# 创建下载目录
RUN mkdir -p wallpapers

# 暴露端口
EXPOSE 8080

CMD ["./WallpaperManager"]
