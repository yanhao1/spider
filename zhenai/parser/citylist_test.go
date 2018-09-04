package parser

import (
	"testing"
	"fmt"
	"io/ioutil"
)

func TestParseCityList(t *testing.T)  {
	//contents,err := fetcher.Fetch("http://www.zhenai.com/zhenghun")
	contents,err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil{
		panic(err)
	}
	fmt.Printf("%s\n",contents)
	//ParseCityList(contents)
}
