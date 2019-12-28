package tool

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func GetFileMD5(path string)string {
	data,err := ioutil.ReadFile(path)
	if err != nil {
		return ""
	}
	value := md5.Sum(data)
	result := hex.EncodeToString(value[:])
	fmt.Printf("md5 value:%v\n",result)
	return result
}

func GetFileHash(path string) string {
	file,err := os.Open(path)
	if err != nil {
		return ""
	}
	defer file.Close()
	hash := sha256.New()
	_,err2 := io.Copy(hash,file)
	if err2 != nil {
		return ""
	}
	value := hash.Sum(nil)
	result := hex.EncodeToString(value)
	fmt.Printf("hash value:%v\n",result)
	return result
}

func WriteFile(path string,data []byte) bool{
	err := ioutil.WriteFile(path,data,0666)
	if err != nil {
		fmt.Printf("write file error:%v\n",err)
		return false
	}
	return true
}

func IsExists(path string)(b bool){
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

func GetFileSize(path string)int64{
	fileInfo,err := os.Stat(path)
	if err != nil {
		return 0
	}
	return fileInfo.Size()
}
