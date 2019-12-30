package service

import (
	"net/http"
	"reflect"
	"testing"
)

func TestDownload(t *testing.T) {
	type args struct {
		img Image
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestGetBingInfo(t *testing.T) {
	type args struct {
		data []byte
		bing *Bing
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
			if err := GetBingInfo(tt.args.data, tt.args.bing); (err != nil) != tt.wantErr {
				t.Errorf("GetBingInfo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetUrl(t *testing.T) {
	type args struct {
		day int
		num int
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
			if got := GetUrl(tt.args.day, tt.args.num); got != tt.want {
				t.Errorf("GetUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

const imgUrl = "http://cn.bing.com/th?id=OHR.ReindeerNorway_ZH-CN5913190372_1920x1080.jpg&rf=LaDigue_1920x1080.jpg&pid=hp"

const imgPath = "D://BingWallpapers//20191224.jpg"

func Test_downloadPic(t *testing.T) {
	downloadPic(imgUrl,imgPath)
}

func Test_readResponseBody(t *testing.T) {
	type args struct {
		resp *http.Response
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := readResponseBody(tt.args.resp); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readResponseBody() = %v, want %v", got, tt.want)
			}
		})
	}
}