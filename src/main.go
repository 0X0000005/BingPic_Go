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
	images := service.GetWeekBingInfo()
	tool.IsExistsAndCreate(service.WALLPAPER, true)
	downloadSuccess, downloadFail, downloadSkip := 0, 0, 0
	ch := make(chan struct{})
	for _, imgInfo := range images {
		//result := service.Download(imgInfo)
		//fmt.Println(imgInfo)
		/*if service.DOWNLOADSUCCESS == result {
			downloadSuccess++
		} else if service.DOWNLOADFAIL == result {
			downloadFail++
		} else if service.DOWNLOADSKIP == result {
			downloadSkip++
		}*/
		go func(imgInfo service.Image) {
			service.Download(imgInfo)
			ch<- struct{}{}
		}(imgInfo)
	}
	for range images{
		<-ch
	}
	fmt.Printf("download end.DOWNLOADSUCCESS=%v,DOWNLOADFAIL=%vDOWNLOADSKIP=%v\n", downloadSuccess, downloadFail, downloadSkip)
}
