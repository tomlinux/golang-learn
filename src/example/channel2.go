package main

import (
	"fmt"
)

/*
测试channel用于通知中断退出的问题
通过channel实现退出的通知
定义一个用于退出的channel比如quit，不断执行任务的线程通过select监听quit的读取，


当读取到quit中的消息时，退出当前的任务线程，这里是主线程通知任务线程退出。

*/
func main() {
	g := make(chan int)
	quit := make(chan bool)

	go func() {
		for {
			select {
			case v := <-g:
				fmt.Println(v)
			case <-quit:
				fmt.Println("B退出")
				return
			}
		}
		fmt.Println("是最后退出吗 ?")
	}()
	//
	for i := 0; i < 3; i++ {
		g <- i
	}
	quit <- true
	fmt.Println("testAB退出")
	//time.Sleep(time.Second*1)
}




