package main

import "net/http"

/*
web server
方法：http.ListenAndServer() ,http.HandleFunc()

 */

func main()  {
	http.ListenAndServe("localhost:9999",nil)
	//error := errors.New("asfsfss")
	//fmt.Println(error)

}



