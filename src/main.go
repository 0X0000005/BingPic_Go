package main

import (
	"fmt"
	"null/BingPic/src/imageinfo"
	"null/BingPic/src/service"
	"null/BingPic/src/tool"
	"time"
)

func main() {
	fmt.Println("*****按下回车开始*****")
	fmt.Scanln()
	start := time.Now()
	serviceRun()
	defer func() {
		cost := time.Since(start)
		fmt.Printf("Program Run Time:%s\n", cost)
		fmt.Println("*****按下回车退出*****")
		fmt.Scanln()
	}()

}

func serviceRun(){
	tool.IsExistsAndCreate(service.WALLPAPER, true)
	images := service.GetWeekBingInfo()
	imageInfos := service.ImageInfoHandler(images)
	result := service.DownloadImages(&imageInfos)
	//result := service.DownloadImage(&imageInfos)
	count(result)
}



func count(images []imageinfo.ImageInfo){
	downloadSuccess, downloadFail, downloadSkip := 0, 0, 0
	for _,imgResult := range images{
		switch imgResult.DownloadResult {
		case service.DOWNLOADFAIL:
			downloadFail++
		case service.DOWNLOADSUCCESS :
			downloadSuccess++
		case service.DOWNLOADSKIP:
			downloadSkip++
		}
	}
	fmt.Printf("Download end.\nDOWNLOADSUCCESS=%v\nDOWNLOADFAIL=%v\nDOWNLOADSKIP=%v\n", downloadSuccess, downloadFail, downloadSkip)
}
