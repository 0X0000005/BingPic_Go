package tool

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetRequest(url string) ([]byte) {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if nil != err {
		fmt.Printf("get connection error:%v\n", err)
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("http status not ok:%v\n", resp.StatusCode)
	}
	fmt.Printf("http status code:%v\n", resp.StatusCode)
	bytes, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		fmt.Printf("http read body error:%v\n", err)
	}
	return bytes
}


