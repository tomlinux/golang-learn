package main

import (
	"fmt"
	"net"
)

func checkError(err error, when string) {
	if err != nil {
		fmt.Println(when, err.Error())
	}
}

func main() {
	//net.UDPAddr{}
	Conn, err := net.Dial("udp", "127.0.0.1:8866")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer Conn.Close()

	Conn.Write([]byte("client hello world"))
	fmt.Println("发送成功","[client hello world]")
	//bytes := []byte("2332")
	buffer := make([]byte, 30)
	n, err := Conn.Read(buffer)
	checkError(err,"con.read")
	fmt.Println("收到消息",string(buffer[:n]))


}
