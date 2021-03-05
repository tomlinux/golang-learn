package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func checkErrors(err error, when string) {
	if err != nil {
		fmt.Println(when, err.Error())
	}
}

func MessageSend(conn net.Conn) {
	var msg string
	reader := bufio.NewReader(os.Stdin)
	for {
		lineBytes, _, _ := reader.ReadLine() //读取一行
		msg = string(lineBytes)              //键盘输入转化为字符串
		if msg == "exit" {
			conn.Close()
			fmt.Println("客户端关闭")
			break
		}

		_, err := conn.Write([]byte(msg)) //输入写入字符串
		if err != nil {
			conn.Close()
			fmt.Println("客户端关闭")
			break
		}

	}

}

func main() {
	// 1.开始网络拨号，建立链接
	ln, err := net.Dial("tcp", "127.0.0.1:9999")
	checkErrors(err, "开始网络拨号的时候：")
	defer ln.Close()

	//发送消息中有阻塞读取标准输入的代码
	//为了避免阻塞住消息的接收，所以把它独立的协程中
	go MessageSend(ln)

	buffer := make([]byte, 1024)
	for {
		n, err := ln.Read(buffer)
		checkErrors(err,"读取消息的时候")

		msg := string(buffer[:n])
		fmt.Println("收到服务器消息", msg)

		if msg == "fuckoff" {
			break
		}

	}

	//// 2. 开始往链接写数据
	//ln.Write([]byte("【client】-> 【server】：hello,服务器端"))
	//fmt.Println("【clent】已经发送消息：【client】-> 【server】：hello,服务器端")
	//// 3. 开始读链接接受到的消息【阻塞】
	//buffer := make([]byte, 1024)
	//n, err := ln.Read(buffer)
	//if err != nil {
	//	fmt.Println("网络错误：",err.Error())
	//}
	//fmt.Println("已经接收到server发送的消息：",string(buffer[:n]))
}
