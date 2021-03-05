package main

import (
	"fmt"
	"time"
)

/*
不定长函数
*/
func function001(x ...int) {
	fmt.Println(len(x))
	for k, v := range x {
		fmt.Println(k, v)
	}
	fmt.Printf("x类型是%T，值是%v \n", x, x)

}

/*
匿名函数
*/
func demo02() {
	//defer 延迟执行函数
	defer func() { fmt.Print("这个是匿名函数，老子要最后执行的") }()

	// go 并发执行匿名函数
	go func() {
		fmt.Println("go并发执行的。 ")
	}()

	// 匿名函数
	var f int64 = func(x, y int64) int64 {
		return x + y
	}(2, 3)

	f1 := func(a, b int) int {
		fmt.Println("劳资普通·无名，下面劳资睡一会")
		time.Sleep(1)
		return a + b
	}(2, 3)

	fmt.Println(f, f1)

}

/*
闭包函数
*/

func main() {
	// 不定函数 function001(1,2,2,3)
	// 匿名函数
	demo02()
}
