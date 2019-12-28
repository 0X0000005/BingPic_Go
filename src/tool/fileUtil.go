package tool

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
)

func getFileMD5(path string)(string,error) {
	data,err := ioutil.ReadFile(path)
	if err != nil {
		return "",err
	}
	value := md5.Sum(data)
	result := hex.EncodeToString(value[:])
	fmt.Printf("md5 value:%v\n",result)
	return result,err
}

func CheckFileAndWrite(path string,data []byte) bool{
	if isExists(path) {
		return true
	}
	err := ioutil.WriteFile(path,data,0666)
	if err != nil {
		fmt.Printf("write file error:%v\n")
		return false
	}
	return true
}

func isExists(path string)(b bool){
	b = true
	_,err := os.Stat(path)
	if err != nil {
		if os.IsExist(err){
			return
		}
		b = false
		return
	}
	return
}

func isDir(path string) bool {
	fileInfo,err := os.Stat(path)
	if err != nil {
		return false
	}
	return fileInfo.IsDir()
}

func IsExistsAndCreate (path string,isDirectory bool) error {
	var err error
	if isDirectory {
		if !isDir(path){
			err = os.Mkdir(path,os.ModePerm)
		}
	} else {
		if !isFile(path){
			_,err = os.Create(path)
		}
	}
	return err
}

func isFile(path string) bool {
	fileInfo,err := os.Stat(path)
	if err != nil {
		return false
	}
	return !fileInfo.IsDir()
}