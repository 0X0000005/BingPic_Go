package main

import (
	"fmt"
	"null/BingPic/src/service"
	"null/BingPic/src/tool"
)

const testurl = "http://cn.bing.com/HPImageArchive.aspx?format=js&idx=0&n=8"

func main() {

	b := tool.GetRequest(service.GetUrl(0,8))
	var bing service.Bing
	//image := tool.GetBingInfo(b,bing)
	//fmt.Println(image)
	err := service.GetBingInfo(b,  &bing)
	if nil != err {
		fmt.Println("err:%v\n", err)
	}
	tool.IsExistsAndCreate(service.WALLPAPER,true)
	for _,imgInfo := range bing.Images {
		enddate := imgInfo.Enddate
		url := imgInfo.Url
		//hash := imgInfo.Hsh
		//copyright := imgInfo.Copyright
		imgUrl := service.BINGIMAGEBASE+url
		fmt.Println(imgUrl)
		bytes := tool.GetRequest(imgUrl)
		fmt.Println(len(bytes))
		tool.CheckFileAndWrite(service.WALLPAPER+"\\"+enddate+".jpg",bytes)
	}
}
