package main

import (
	"fmt"
	"null/BingPic/src/service"
	"null/BingPic/src/tool"
	"time"
)

const testurl = "http://cn.bing.com/HPImageArchive.aspx?format=js&idx=0&n=8"

func main() {
	start := time.Now()
	defer func() {
		cost := time.Since(start)
		fmt.Printf("program run time:%s\n", cost)
	}()

	tool.IsExistsAndCreate(service.WALLPAPER, true)
	images := service.GetWeekBingInfo()
	imageInfos := service.ImageInfoHandler(images)
	//fmt.Println(imageInfos)
	service.DownloadImages(&imageInfos)
	//service.DownloadFirst(&(imageInfos[0]))
	downloadSuccess, downloadFail, downloadSkip := 0, 0, 0
	fmt.Printf("download end.DOWNLOADSUCCESS=%v,DOWNLOADFAIL=%vDOWNLOADSKIP=%v\n", downloadSuccess, downloadFail, downloadSkip)
}
