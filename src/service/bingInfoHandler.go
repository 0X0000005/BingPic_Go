package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"null/BingPic/src/tool"
	"os"
	"strconv"
)

func GetUrl(day, num int) string {
	return BINGIMAGEINFOURL + strconv.Itoa(day) + "&n=" + strconv.Itoa(num)
}

func GetBingInfo(data []byte, bing *Bing){
	err := json.Unmarshal(data, bing)
	if nil != err {
		fmt.Printf("parse json error:%v\n", err)
	}
}

func GetWeekBingInfo()[]Image{
	bytes1 := tool.GetRequest(GetUrl(0,8))
	bytes2 := tool.GetRequest(GetUrl(8,8))
	var bing1 Bing
	var bing2 Bing
	GetBingInfo(bytes1,&bing1)
	GetBingInfo(bytes2,&bing2)
	images := append(bing1.Images,bing2.Images...)
	return images
}

//下载图片
func Download(img Image) {
	enddate := img.Enddate
	url := img.Url
	//hash := img.Hsh
	//copyright := img.Copyright
	imgUrl := BINGIMAGEBASE + url
	imgPath := WALLPAPER + "\\" + enddate + ".jpg"
	downloadPic(imgUrl, imgPath)
}

func downloadPic(imgUrl string, imgPath string) {
	resp, err := http.Get(imgUrl)
	defer resp.Body.Close()
	if nil != err {
		fmt.Printf("get connection error:%v\n", err)
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("http status not ok:%v\n", resp.StatusCode)
	}
	var bytes []byte
	if tool.IsExists(imgPath){
		f,_ := os.Stat(imgPath)
		headSize,_ := strconv.ParseInt(resp.Header.Get("Content-Length"), 10, 64)
		fileSize := f.Size()
		if headSize != fileSize {
			fmt.Printf("%v exists,head size:%v,file size:%v,re download\n",f.Name(),headSize,fileSize)
			bytes = readResponseBody(resp)
		}else {
			fmt.Printf("%v exists\n",f.Name())
		}
	} else {
		fmt.Printf("%v not exists,downliad\n",imgPath)
		bytes = readResponseBody(resp)
	}
	if len(bytes)>0 {
		tool.WriteFile(imgPath, bytes)
	}
}

func readResponseBody(resp *http.Response) []byte {
	bytes, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		fmt.Printf("http read body error:%v\n", err)
	}
	return bytes
}
