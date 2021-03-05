package main

import (
	"fmt"
	"net"
)

func serverHandleError(err error, when string) {
	if err != nil {
		fmt.Println(when, err)
	}

}

func receiveUDPMsg(updConn *net.UDPConn){
	buffer := make([]byte, 30)
	//从udpConn读取客户端发送过来的数据，放在缓冲区中(阻塞方法)
	//返回值：n=读到的字节长度,remoteAddr=客户端地址,err=错误
	n, remoteAddr, err := updConn.ReadFromUDP(buffer)
	if err != nil{
		fmt.Println("Error",err.Error())
		return
	}

	fmt.Printf("接收到来自%v的消息：%s", remoteAddr,string(buffer[0:n]))

	//向远程地址中写入数据
	_, err = updConn.WriteToUDP([]byte("hao nimei"), remoteAddr)
	serverHandleError(err,"WriteToUDP 像远程地址写人数据的时候")


}


func main() {
	//listener, err := net.Listen("udp", "127.0.0.1")
	udpAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8866")
	serverHandleError(err, "net.ResolveIPAddr发生错误")

	//在解析的地址上面简历udp 监听
	fmt.Println("开始在udp监听8866端口....")
	udpConn, err := net.ListenUDP("udp", udpAddr)

	serverHandleError(err, "net.ListenUPD监听发送错误")

	defer udpConn.Close() // 关闭链接

	receiveUDPMsg(udpConn)

}
