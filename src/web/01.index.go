package main

import (
	"fmt"
	"net/http"
)

/*
web server
方法：http.ListenAndServer() ,http.HandleFunc()

*/
func demo02(w http.ResponseWriter,r *http.Request)  {
	fmt.Println("hello world" )
	w.Write([]byte("hello world"))
}

func main() {
	//http.ListenAndServe("localhost:9999",nil)
	////error := errors.New("asfsfss")
	////fmt.Println(error)
	server := http.Server{
		Addr: "localhost:8888",
	}
	http.HandleFunc("/test02",demo02)
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		fmt.Println(w,r.Form)
	})




	server.ListenAndServe()


}





