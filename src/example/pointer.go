package main

import (
	"fmt"
)

/*
指针概述
指针就是地址，指针变量的值就是内存中的一块地址编号
所有普通类型都有对应的指针类型：int,*float64,*bool,*book,[10]int,*map[string]interface{}
对指针变量的打印推荐使用%p占位符，代表以地址的形式打印
对普通变量取地址： agePtr := &age
取地址中的值：age := *agePtr
*/

func demoa1() {
	//声明普通变量
	var age int = 20

	//对age取地址，赋值给整型指针agePointer
	var agePointer *int = &age

	//类型：整型指针，值：age的内存地址
	//type=*int,value=0xc042062080
	fmt.Printf("type=%T,value=%p\n", agePointer, agePointer)

	//*agePointer取地址中的值，20
	fmt.Println("agePointer所指向的地址中的值", *agePointer)
}

func demoa2() {
	var age int = 20
	//agePtr中存放的是age的内存地址
	var agePtr *int = &age
	//apPtr是一个指向指针（agePtr）的指针
	var apPtr **int = &agePtr
	//type=**int,value=0xc042082018
	fmt.Printf("type=%T,value=%p\n", apPtr, apPtr)
	//*apPtr得到agePtr（存放age的地址），*（*apPtr）得到age的值20
	fmt.Println(**apPtr)
}


func main()  {
	demoa2()
	var a1 [10]int = [10]int{1,2,3}
	fmt.Print(a1)
	age := 10
	var a2ptr *int = &age
	fmt.Print(a2ptr,age)

}