package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"
)

/*
1.打开和关闭文件
*/
func demo410() {
	//打开文件
	filePtr, err := os.Open("d:/小黑子.avi")
	if err != nil {
		fmt.Println("文件打开失败,err=", err)
	} else {
		fmt.Println("文件打开成功")
	}

	//关闭
	//defer filePtr.Close()
	defer func() {
		filePtr.Close()
		fmt.Println("文件已关闭")
	}()

	fmt.Println("愉快滴欣赏avi艺术")
	time.Sleep(3 * time.Second)

}

/*
2.读取数据

*/
func main() {
	//打开和关闭文件
	//os.O_RDONLY 以只读模式打开
	//0765 文件主人的权限7（4+2+1=可读+可写+可执行），文件的用户组下的人的权限6（4+2=可读+可写），其他人的权限5（4+1=可读+可执行）
	filePtr, err := os.OpenFile("d:/abc.txt", os.O_RDONLY, 0765)
	if err != nil {
		fmt.Println("文件打开失败,err=", err)
		return
	} else {
		fmt.Println("打开文件成功")
	}

	//打开时立即挂起关闭程序
	defer func() {
		defer filePtr.Close()
		fmt.Println("文件已关闭")
	}()

	//创建文件的读取器
	reader := bufio.NewReader(filePtr)

	//循环读取
	for {
		//以换行符为界，分批次读取数据，得到str（str中包括一个换行符）
		str, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("读取失败,err=", err)

			//如果已到文件末尾，就结束循环
			if err == io.EOF {
				break
			}
		}

		//打印读到的内容
		fmt.Print(str)
	}

	fmt.Println("读取结束")
}

/*
3.读取数据3
*/
func main43() {
	bytes, err := ioutil.ReadFile("d:/abc.txt")
	if err != nil {
		fmt.Println("读取失败，err=", err)
	} else {
		contentStr := string(bytes)
		fmt.Println(contentStr)
	}
}

/*
4.写入数据
写入数据分三种模式
- 创写模式
- 覆写模式
- 追加模式
*/
func main44() {

	//打开并延时关闭文件
	//创写模式
	//file, err := os.OpenFile("d:/大黑子.avi", os.O_CREATE|os.O_WRONLY, 0754)
	//覆写模式
	//file, err := os.OpenFile("d:/兄弟连.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0754)
	//追加模式
	file, err := os.OpenFile("d:/兄弟连.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0754)

	if err != nil {
		fmt.Println("文件打开失败,err=", err)
		return
	} else {
		fmt.Println("打开文件成功")
	}
	defer file.Close()

	//创建写入器
	writer := bufio.NewWriter(file)

	//执行带缓冲的写入
	writer.WriteString("大海啊你全他妈是水\n")
	writer.WriteString("蜘蛛啊你全他妈是腿\n")
	writer.WriteString("辣椒啊真他妈辣嘴\n")
	writer.WriteString("不好好学习真他妈后悔\n")

	//强制将缓冲区中的内容写入文件
	writer.Flush()

	fmt.Println("写入成功！")
}

/*
5.判断文件是否存在
*/

func CheckIfFileExist(filename string) (exists bool, info string) {
	fileInfo, err := os.Stat(filename)
	fmt.Println(fileInfo, err)
	if fileInfo != nil && err == nil {
		fmt.Printf("%s文件存在\n", filename)
		exists = true
	} else if os.IsNotExist(err) {
		fmt.Printf("%s文件不存在\n", filename)
		exists = false
	} else {
		fmt.Println("搞不清存不存在!")
		exists = false
		info = "发生了一些奇奇怪怪的事情..."
	}
	return
}
func main45() {
	CheckIfFileExist("d:/小黑子.avi")
	CheckIfFileExist("d:/小白子.avi")
}

/*
5.快捷写入
*/

func main46() {
	//默认模式：os.O_WRONLY|os.O_CREATE|os.O_TRUNC=只写+不存在就创建+【覆写模式】
	err := ioutil.WriteFile("d:/abc.txt", []byte("再看我\n再看我就把你喝掉"), 0754)
	if err != nil {
		fmt.Println("写入失败")
	} else {
		fmt.Println("写入成功")
	}
	fmt.Println("写入结束")
}

