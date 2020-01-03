package tool

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"testing"
)

func TestGetFileHash(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFileHash(tt.args.path); got != tt.want {
				t.Errorf("GetFileHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetFileMD5(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFileMD5(tt.args.path); got != tt.want {
				t.Errorf("GetFileMD5() = %v, want %v", got, tt.want)
			}
		})
	}
}

const path = "D:\\BingWallpapers\\20191221.jpg"

func TestGetFileSize(t *testing.T) {
	fmt.Println(GetFileSize(path))
}

func TestIsExists(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name  string
		args  args
		wantB bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotB := IsExists(tt.args.path); gotB != tt.wantB {
				t.Errorf("IsExists() = %v, want %v", gotB, tt.wantB)
			}
		})
	}
}

func TestIsExistsAndCreate(t *testing.T) {
	type args struct {
		path        string
		isDirectory bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := IsExistsAndCreate(tt.args.path, tt.args.isDirectory); (err != nil) != tt.wantErr {
				t.Errorf("IsExistsAndCreate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWriteFile(t *testing.T) {
	v,err := strconv.ParseInt("",10,64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(v)
}

func Test_isDir(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isDir(tt.args.path); got != tt.want {
				t.Errorf("isDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isFile(t *testing.T) {
	files, _ := ioutil.ReadDir("D://BingWallpapers")
	for _, f := range files {
		fmt.Printf("fileName:%v,size:%v\n",f.Name(),f.Size())
	}

}