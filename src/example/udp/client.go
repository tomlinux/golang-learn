package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	//请求连接服务器，得到连接对象
	conn, err := net.Dial("udp", "127.0.0.1:8848")
	defer conn.Close()
	if err != nil {
		fmt.Println("网络连接出错")
		os.Exit(1)
	}

	//向连接中写入消息
	conn.Write([]byte("hello nimei"))
	fmt.Println("发送消息", "hello nimei")

	//读取代表收取消息(阻塞)
	buffer := make([]byte, 30)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("读取消息错误：err=", err)
		os.Exit(1)
	}

	fmt.Println("收到消息", string(buffer[:n]))

}
