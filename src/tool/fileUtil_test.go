package tool

import (
	"fmt"
	"testing"
)

const path = "C://Users//wzy//Desktop//menu.json"

const dir = "C://Users//wzy//Desktop"

func TestGetFileMD5(t *testing.T) {
	_,err := getFileMD5(path)
	if err != nil {
		panic(err)
	}

}

func Test_checkFileAndWrite(t *testing.T) {

}

func Test_fileIsExists(t *testing.T) {
}

func Test_getFileMD5(t *testing.T) {
}

func Test_isDir(t *testing.T) {
	fmt.Println(isDir(path))
	fmt.Println(isDir(dir))
}

func Test_isDirAndCreate(t *testing.T) {
	err := isExistsAndCreate(path+"dd",false)
	fmt.Println(err)
	if err != nil {
		panic(err)
	}
}

func Test_isFile(t *testing.T) {
	fmt.Println(isFile(path))
	fmt.Println(isFile(dir))
}