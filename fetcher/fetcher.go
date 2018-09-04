package fetcher

import (
	"fmt"
	"golang.org/x/text/transform"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io/ioutil"
	"net/http"
)

func Fetch(url string)([]byte,error){
	resp , err  := http.Get(url)

	if err != nil{
		return nil ,err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		//fmt.Println("Error:status code",resp.StatusCode)
		return nil,fmt.Errorf("wrong status code %d" , resp.StatusCode)
	}
	//bodyReader := bufio.NewReader(resp.Body)
	utf8Reader := transform.NewReader(resp.Body,simplifiedchinese.GBK.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
	//printCityList(all)
	//if err != nil {
	//	panic(err)
	//}
}