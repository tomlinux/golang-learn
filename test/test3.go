package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

var (
	NotPV []string = []string{"css", "js", "class", "gif", "jpg", "jpeg", "png", "bmp", "ico", "rss", "xml", "swf"}
)

const big = 0xFFFFFF

func LogPV(urls string) bool {
	ext := filepath.Ext(urls)
	if ext == "" {
		return true
	}
	for _, v := range NotPV {
		if v == strings.ToLower(ext) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(LogPV("https://zcmops.oss-cn-hangzhou.aliyuncs.com/ops/C29FA2642ED5E0D7C062C80B900CEC5E.png"))
}
