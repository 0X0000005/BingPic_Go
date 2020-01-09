package tool

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetRequest(t *testing.T) {
	type args struct {
		url string
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
			if got := GetRequest(tt.args.url); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_handlerServer(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
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

func Test_startServer(t *testing.T) {
	go startServer("localhost:8080")
	openBrowser("localhost:8080")
	select{}
}