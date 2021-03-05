package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

/*
文件拷贝1
*/
func copyFile(srcFilename string,dstFilename string) {
	bytes, _ := ioutil.ReadFile(srcFilename)
	err := ioutil.WriteFile(dstFilename, bytes, 0754)
	if err != nil {
		fmt.Println("拷贝失败，err=", err)
	} else {
		fmt.Println("拷贝成功!")
	}
}
func main51() {
	copyFile("d:/找你妹.txt","d:/找自己.txt")
}

/*
文件拷贝2

*/
func main52() {
	//打开来源文件和目标文件（目标文件可能并不存在）
	srcFile, _ := os.OpenFile("d:/找你妹.txt", os.O_RDONLY, 0666)
	dstFile, _ := os.OpenFile("d:/找我妹.txt", os.O_WRONLY|os.O_CREATE, 0666)

	//程序的最后关闭这两个文件
	defer srcFile.Close()
	defer dstFile.Close()

	//使用标准库提供的API实现文件拷贝，返回的是拷贝的字节数
	written, err := io.Copy(dstFile, srcFile)
	if err!=nil{
		fmt.Println("拷贝失败,err=",err)
	}else{
		fmt.Println("拷贝成功,字节数=",written)
	}
}


/*
实现文件字符数统计
*/
func CountCharsOfFile(path string)map[string]int {
	bytes, _ := ioutil.ReadFile(path)
	str := string(bytes)

	var numberCount,letterCount,spaceCount,othersCount int

	//遍历字符串中的每个字符
	for i,c := range str{
		fmt.Printf("No%d,%c\n",i,c)

		//逐个判断每个字符在字符集中的序号
		switch {
		case c >= '0' && c<='9':
			numberCount++
		case (c >= 'a' && c<='z') || (c >= 'A' && c<='Z'):
			letterCount++
		case c == ' ' || c=='\t':
			spaceCount++
		default:
			othersCount++
		}
	}
	// 返回map格式的
	return map[string]int{"数字":numberCount,"字母":letterCount,"空白":spaceCount,"其它":othersCount}
}
func main() {
	retMap := CountCharsOfFile("d:/temp/test.txt")
	fmt.Println(retMap)
}
