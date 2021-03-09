package main

import (
	"io/ioutil"
	"net/http"
	"os"
)

/*
http 请求原理
1. 客户机通过TCP/IP协议建立到服务器的TCP连接
2. 客户端向服务器发送HTTP协议请求包，请求服务器里的资源文档
3. 服务器向客户机发送HTTP协议应答包，如果请求的资源包含有动态语言的内容，那么服务器会调用动态语言的解释引擎负责处理“动态内容”，
并将处理得到的数据返回给客户端
4. 客户机与服务器断开。由客户端解释HTML文档，在客户端屏幕上渲染图形结果

* Request：用户请求的信息，用来解析用户的请求信息，包括post、get、cookie、url等信息
* Response：服务器需要反馈给客户端的信息
* Conn：用户的每次请求链接
* Handler：处理请求和生成返回信息的处理逻辑
*/

//1、创建Listen Socket, 监听指定的端口, 等待客户端请求到来。
//2、Listen Socket接受客户端的请求, 得到Client Socket, 接下来通过Client Socket与 客户端通信。
//3、处理客户端的请求, 首先从Client Socket读取HTTP请求的协议头, 如果是POST 方法, 还可能要读取客户端提交的数据,然后交给相应的handler处理请求, handler处理完 毕准备好客户端需要的数据, 通过Client Socket写给客户端。
//原文链接：https://blog.csdn.net/u010986776/article/details/82356712

func  main()  {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte(request.RemoteAddr+";"))
		writer.Write([]byte(request.Method+";"))
		writer.Write([]byte("hello, world"))
	})

	http.HandleFunc("/demo", func(writer http.ResponseWriter, request *http.Request) {
		//从本地html文件中读入HTML页面的原始字节
		pathDir, _ := os.Getwd()
		contentBytes, _ := ioutil.ReadFile(pathDir+ "/src/example/template/demo.html")
		//向客户端写出响应
		writer.Write(contentBytes)
	})



	http.ListenAndServe("127.0.0.1:8080", nil)

}

