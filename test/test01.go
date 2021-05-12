package main

import (
	"fmt"
	"net/http"
	"net/url"
)

type Hander interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

func main() {
	resp, err := http.Get("https://www.baidu.com")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()
	//fmt.Println(resp.Body)
	fmt.Printf("类型是%T，值是%v", resp, resp)
	//http.DefaultServeMux.HandleFunc()

	fmt.Println("#######")
	Header := map[string][]string{
		"Accept-Encoding": {"gzip, deflate"},
		"Accept-Language": {"en-us"},
		"Foo":             {"Bar", "two"},
	}
	fmt.Println(Header)

	tester := map[string]int{"key01": 2}
	fmt.Println(tester["key01"])

	//http.Handle
	//http.FileServer()
	fmt.Println(url.URL{})

}

