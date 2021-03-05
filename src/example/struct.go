package main

import (
	"fmt"
	"runtime"
)

type book struct {
	name  string
	price float64
}

/*
创建对象指针并给属性赋值
1.内置函数new(Type) 可以创建任意类型的对象，并获得其指针
2.通过结构体指针访问结构体的成员，与通过值来访问，写法上是一模一样的
*/

func printCallerName() string {
	pc, _, _, _ := runtime.Caller(2)
	return runtime.FuncForPC(pc).Name()
}

func demo01() {
	bookPtr := new(book)
	fmt.Printf("bookPtr指针类型是%T，value=%#v ，address=%p \n", bookPtr, *bookPtr, bookPtr)
	bookPtr.name = "测测"
	bookPtr.price = 10.26
	fmt.Println(*bookPtr)
}

/*
工具函数：通过值传递查看结构体属性
*/
func showBookInfo(b book) {
	fmt.Printf("b的类型是%T,value:%#v,address:%p \n ", b, b, &b)
	fmt.Println(b.name, b.price)
}

/*
工具函数：通过引用传递（指针传递）查看结构体属性
*/
func showBookInfo2(b *book) {
	fmt.Println(b.name, b.price)
}

/*
分别通过指针和值来访问结构体的属性
*值传递是拷贝式的，传递的是一个副本
*值传递中，无论外界函数如何修改被传参数，都不会影响到本体
*引用传递，传递的是对象的地址
*引用传递对被传参数的修改，将直接改变本体！
*/
func demo03() {
	b1 := book{"三国", 45.67}
	fmt.Printf("b1的类型是%T,value:%#v,address:%p \n", b1, b1, &b1)

	showBookInfo(b1)
	showBookInfo2(&b1)

	bp := new(book)
	bp.name = "三国2"
	bp.price = 45.68

	showBookInfo(*bp)
	showBookInfo2(bp)

}

func main() {
	//b := book{"天下无贼", 10.24}
	//b.name = "天下无贼2"
	//fmt.Print(b.name, "\t", b.price)
	//b1 := book{}
	//b1.name = "天下名师"
	//b1.price = 10.25
	//fmt.Printf("type=%T，value=%#v\n", b1, b1)
	demo03()

}
