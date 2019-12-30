package main

import (
	"null/BingPic/src/service"
	"null/BingPic/src/tool"
)

const testurl = "http://cn.bing.com/HPImageArchive.aspx?format=js&idx=0&n=8"

func main() {
	images := service.GetWeekBingInfo()
	tool.IsExistsAndCreate(service.WALLPAPER,true)
	for _,imgInfo := range images {
		service.Download(imgInfo)
		//fmt.Println(imgInfo)
	}
}
