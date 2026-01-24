#!/bin/bash

# 默认值
PLATFORMS="windows/amd64 linux/amd64"
OUTPUT_DIR="dist"
APP_NAME="WallpaperManager"
USE_UPX=true

# 解析参数 (简单实现)
if [ "$1" == "linux" ]; then
    PLATFORMS="linux/amd64"
elif [ "$1" == "windows" ]; then
    PLATFORMS="windows/amd64"
elif [ "$1" == "all" ]; then
    PLATFORMS="windows/amd64 linux/amd64"
fi

# 清理输出目录
rm -rf $OUTPUT_DIR
mkdir -p $OUTPUT_DIR

# 检查 UPX
if ! command -v upx &> /dev/null; then
    echo "未找到 UPX。跳过压缩。"
    USE_UPX=false
fi

echo "下载依赖..."
go mod tidy

for PLATFORM in $PLATFORMS; do
    GOOS=${PLATFORM%/*}
    GOARCH=${PLATFORM#*/}
    OUTPUT_NAME=$APP_NAME
    
    if [ "$GOOS" == "windows" ]; then
        OUTPUT_NAME+=".exe"
    fi

    echo "正在构建 $GOOS/$GOARCH..."
    env GOOS=$GOOS GOARCH=$GOARCH go build -o $OUTPUT_DIR/$OUTPUT_NAME-$GOOS-$GOARCH cmd/server/main.go

    if [ $? -ne 0 ]; then
        echo "发生错误！终止脚本执行..."
        exit 1
    fi

    if [ "$USE_UPX" = true ]; then
        echo "正在使用 UPX 压缩..."
        upx $OUTPUT_DIR/$OUTPUT_NAME-$GOOS-$GOARCH
    fi
done

# 复制配置和 Web 资源
echo "复制配置和资源..."
cp config.yaml $OUTPUT_DIR/
cp -r web $OUTPUT_DIR/

echo "构建完成！构建产物在 $OUTPUT_DIR"
