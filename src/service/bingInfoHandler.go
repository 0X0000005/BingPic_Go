package service

import (
"encoding/json"
"errors"
"fmt"
"io/ioutil"
"net/http"
"null/BingPic/src/imageinfo"
"null/BingPic/src/tool"
"os"
"strconv"
)

func GetUrl(day, num int) string {
	return BINGIMAGEINFOURL + strconv.Itoa(day) + "&n=" + strconv.Itoa(num)
}

func GetBingInfo(data []byte, bing *imageinfo.Bing) {
	err := json.Unmarshal(data, bing)
	if nil != err {
		fmt.Printf("parse json error:%v\n", err)
	}
}

func getBingInfo(day int,num int) ([]imageinfo.Image,error){
	bytes,err := tool.GetRequest(GetUrl(day, num))
	if err != nil {
		return nil,err
	}
	var bing imageinfo.Bing
	err2 := json.Unmarshal(bytes, &bing)
	if err2 != nil {
		return nil,err2
	}
	return bing.Images,nil
}

func GetWeekBingInfo() []imageinfo.Image {
	images1,err1 := getBingInfo(0,8)
	if err1 != nil {
		fmt.Printf("get bing info error,day=0,num=8,error message:=%v\n",err1)
		return nil
	}
	images2,err2 := getBingInfo(8,8)
	if err2 != nil {
		fmt.Printf("get bing info error,day=8,num=8,error message:=%v\n",err2)
		return nil
	}
	images := append(images1, images2...)
	return images
}

func ImageInfoHandler(images []imageinfo.Image) []imageinfo.ImageInfo {
	length := len(images)
	imageInfos := make([]imageinfo.ImageInfo, length, length)
	for i, imgInfo := range images {
		imageInfo := imageinfo.ImageInfo{}
		imageInfo.Desc = imgInfo.Copyright
		imageInfo.DownloadUrl = BINGIMAGEBASE + imgInfo.Url
		imageInfo.ImageName = imgInfo.Enddate + Symbol_dot + Extension
		imageInfo.ImagePath = WALLPAPER + Symbol_backslashs + imageInfo.ImageName
		imageInfos[i] = imageInfo
	}
	return imageInfos
}

//下载图片
//1下载成功 -1下载失败 0不下载
func download(imgInfo *imageinfo.ImageInfo) {
	resp, err := http.Get(imgInfo.DownloadUrl)
	defer func() {
		if nil != resp {
			resp.Body.Close()
		}
	}()
	if nil != err {
		fmt.Printf("get connection error:%v\n", err)
		imgInfo.Err = err
		imgInfo.DownloadResult = DOWNLOADFAIL
		return
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("http status not ok:%v\n", resp.StatusCode)
		imgInfo.Err = errors.New("http status not ok:" + string(resp.StatusCode))
		imgInfo.DownloadResult = DOWNLOADFAIL
		return
	}
	var bytes []byte
	if tool.IsExists(imgInfo.ImagePath) {
		f, err_Stat := os.Stat(imgInfo.ImagePath)
		if err_Stat != nil {
			imgInfo.Err = errors.New("http status not ok:" + string(resp.StatusCode))
			imgInfo.DownloadResult = DOWNLOADFAIL
			return
		}
		fileSize := f.Size()
		imgInfo.LocalSize = fileSize
		headSize, err_Head := strconv.ParseInt(resp.Header.Get("Content-Length"), 10, 64)
		if err_Head != nil {
			fmt.Printf("get resp head [Content-Length] error:%v\n", err_Head)
		}
		imgInfo.ServerSize = headSize
		if headSize != fileSize {
			fmt.Printf("%v exists,head size:%v,file size:%v,re download\n", f.Name(), headSize, fileSize)
			b, err_read := ioutil.ReadAll(resp.Body)
			bytes = b
			if err_read != nil {
				imgInfo.Err = err_read
				imgInfo.DownloadResult = DOWNLOADFAIL
				return
			}
		} else {
			fmt.Printf("%v exists\n", f.Name())
			imgInfo.DownloadResult = DOWNLOADSKIP
		}
	} else {
		fmt.Printf("%v not exists,download\n", imgInfo.ImageName)
		b, err_read := ioutil.ReadAll(resp.Body)
		bytes = b
		if err_read != nil {
			imgInfo.Err = err_read
			imgInfo.DownloadResult = DOWNLOADFAIL
			return
		}
	}
	if len(bytes) > 0 {
		err_write := ioutil.WriteFile(imgInfo.ImagePath, bytes, 0666)
		if err_write != nil {
			imgInfo.Err = err_write
			imgInfo.DownloadResult = DOWNLOADFAIL
		} else {
			imgInfo.DownloadResult = DOWNLOADSUCCESS
		}
	}
}

func DownloadFirst(imgInfo *imageinfo.ImageInfo) {
	download(imgInfo)
	fmt.Println(imgInfo)
}

func DownloadImages(imgInfos *[]imageinfo.ImageInfo) []imageinfo.ImageInfo {
	ch := make(chan imageinfo.ImageInfo)
	for _, imgInfo := range *imgInfos {
		info := imageinfo.New(imgInfo)
		go func(imgInfo *imageinfo.ImageInfo) {
			download(imgInfo)
			ch <- *imgInfo
		}(&info)
	}
	var results []imageinfo.ImageInfo
	for range *imgInfos {
		result := <-ch
		results = append(results, result)
	}
	//fmt.Println(results)
	return results
}

var taskMap = make(map[string]string)

func isLoadTask(taskName string) bool {
	_, ok := taskMap[taskName]
	if !ok {
		taskMap[taskName] = taskName
	}
	return ok
}

func DownloadImage(imgInfos *[]imageinfo.ImageInfo) []imageinfo.ImageInfo {
	var results []imageinfo.ImageInfo
	for _, imgInfo := range *imgInfos {
		info := imageinfo.New(imgInfo)
		download(&info)
		results = append(results, info)
	}
	//fmt.Println(results)
	return results
}
