package service

import (
	"encoding/json"
	"strconv"
)

func GetUrl(day,num int) string {
	return BINGIMAGEINFOURL + strconv.Itoa(day) + "&n=" + strconv.Itoa(num)
}

func GetBingInfo(data []byte, bing *Bing) error {
	err := json.Unmarshal(data, bing)
	return err
}

func GetImageInfo(bing Bing) {

}
