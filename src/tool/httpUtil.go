package tool

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

func GetRequest(url string) ([]byte,error) {
	resp, err := http.Get(url)
	defer func() {
		if nil != resp {
			resp.Body.Close()
		}
	}()
	if nil != err {
		return nil,err
	}
	if resp.StatusCode != http.StatusOK {
		errorInfo := fmt.Sprintf("http status not ok:%v\n",resp.StatusCode)
		return nil,errors.New(errorInfo)
	}
	fmt.Printf("http status code:%v\n", resp.StatusCode)
	bytes, err2 := ioutil.ReadAll(resp.Body)
	if nil != err2 {
		return nil,err2
	}
	return bytes,nil
}



func startServer(url string){
	http.HandleFunc("/",handlerServer)
	log.Fatal(http.ListenAndServe(url,nil))
}

func openBrowser(url string){
	if !strings.HasPrefix(url,"http://"){
		url = "http://" + url
	}
	cmd := exec.Command("cmd", "/C", "start "+url)
	cmd.Run()
}

func handlerServer(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w, "URL.Path = %v\n", "aaaa")
}

