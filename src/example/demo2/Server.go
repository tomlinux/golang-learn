package main

import (
	"fmt"
	"net"
	"strings"
)

func checkError(err error, when string) {
	if err != nil {
		fmt.Println(when, err.Error())
	}
}

func doClientRequest(conn net.Conn) {
	// 创建缓冲区的buffer空间
	buffer := make([]byte, 1024)
	for {
		d, err := conn.Read(buffer)
		checkError(err, "网络出错的错误是：")
		//fmt.Println("收到的消息长度是",d)
		if d > 0 {
			remoteAddr := conn.RemoteAddr()
			msg := string(buffer[:d])
			fmt.Println("收到消息", msg, "来自", remoteAddr)
			if strings.Contains(msg, "钱") {
				conn.Write([]byte("fuckoff"))
				break
			}
			msg += "消息已经预定！"
			conn.Write([]byte(msg))
		}

	}

}

//make(map[string]net.Conn)
func main() {
	// 1. 服务器监听端口 打开监听器listener
	listener, err := net.Listen("tcp", "127.0.0.1:9999")
	checkError(err, "listener打开的时候")
	defer listener.Close()
	fmt.Println("服务端正在等待中.......")
	//2. 处理每个客户端
	for {
		conn, err := listener.Accept()
		checkError(err, "接受客户端请求")

		// 并行处理每个客户端请求
		go doClientRequest(conn)

	}

}
