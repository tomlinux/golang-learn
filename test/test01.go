package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("https://www.baidu.com")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()
	//fmt.Println(resp.Body)
	fmt.Printf("类型是%T，值是%v",resp,resp)
}
