package main

import "fmt"
/*
切片概述
*切片可以理解为长度可以动态变化的数组
*切片和数组的相同点1：通过下标来访问元素
*切片和数组的相同点2：通过下标或range方式遍历元素
*不同点是切片的长度不是固定的，你可以动态地向切片中追加新的元素
*切片中可以容纳的元素个数成为容量，容量>=长度
*你可以通过len(slice)和cap(slice)分别获取切片的长度和容量
*通过make(type,len,cap)的方式你可以创建出自定义初始长度和容量的切片
*在追加元素的过程中，如果容量不够用时，就存在动态扩容的问题
*动态扩容采用的是倍增策略，即：新容量=2*旧容量
*扩容后的切片会得到一片新的连续内存地址，所有元素的地址都会随之发生改变
 */

//向切片中追加新的元素
//
//容量每突破一次就翻倍
//每翻倍一次，切片就获得一片新的堆地址，每一个元素的地址也都随之发生变化
//切片名称的栈地址是不会变的

func demob2() {
	s := make([]int, 0)
	s = append(s, 1)
fmt.Println(s[0])
	fmt.Println("++++++++")
	//value=[1],len=1,cap=1,saddr=0xc04205c3e0,elemaddr=0xc042062080
	fmt.Printf("value=%v,len=%d,cap=%d,saddr=%p,elemaddr=%p\n", s, len(s), cap(s),&s,&s[0])
	s = append(s, 2)
	//value=[1 2],len=2,cap=2,saddr=0xc04205c3e0,elemaddr=0xc0420620d0
	fmt.Printf("value=%v,len=%d,cap=%d,saddr=%p,elemaddr=%p\n", s, len(s), cap(s),&s,&s[0])
	s = append(s, 3)
	//value=[1 2 3],len=3,cap=4,saddr=0xc04205c3e0,elemaddr=0xc0420600e0
	fmt.Printf("value=%v,len=%d,cap=%d,saddr=%p,elemaddr=%p\n", s, len(s), cap(s),&s,&s[0])
	s = append(s, 4)
	//value=[1 2 3 4],len=4,cap=4,saddr=0xc04205c3e0,elemaddr=0xc0420600e0
	fmt.Printf("value=%v,len=%d,cap=%d,saddr=%p,elemaddr=%p\n", s, len(s), cap(s),&s,&s[0])
	s = append(s, 5)
	//value=[1 2 3 4 5],len=5,cap=8,saddr=0xc04205c3e0,elemaddr=0xc042088100
	fmt.Printf("value=%v,len=%d,cap=%d,saddr=%p,elemaddr=%p\n", s, len(s), cap(s),&s,&s[0])
	s = append(s, 6, 7, 8)
	//value=[1 2 3 4 5 6 7 8],len=8,cap=8,saddr=0xc04205c3e0,elemaddr=0xc042088100
	fmt.Printf("value=%v,len=%d,cap=%d,saddr=%p,elemaddr=%p\n", s, len(s), cap(s),&s,&s[0])
	s = append(s, 9)
	//value=[1 2 3 4 5 6 7 8 9],len=9,cap=16,saddr=0xc04205c3e0,elemaddr=0xc042092080
	fmt.Printf("value=%v,len=%d,cap=%d,saddr=%p,elemaddr=%p\n", s, len(s), cap(s),&s,&s[0])
}


func main()  {
	//有类无类冒等三种方式创建出类型、值、长度、容量都相同的切片
	var s2 = []int{0,1,2}
	var s1 []int = []int{0,1,2}
	s3 := []int{0,1,2}

	//这里的初始长度和容量是相等的
	//type=[]int,value=[0 1 2],len=3,cap=3
	fmt.Printf("type=%T,value=%v,len=%d,cap=%d\n",s1,s1,len(s1),cap(s1))
	fmt.Printf("type=%T,value=%v,len=%d,cap=%d\n",s2,s2,len(s2),cap(s2))
	fmt.Printf("type=%T,value=%v,len=%d,cap=%d\n",s3,s3,len(s3),cap(s3))

	//创建长度和容量都为3的切片
	s4 := make([]int, 3)
	fmt.Printf("type=%T,value=%v,len=%d,cap=%d\n", s4, s4, len(s4), cap(s4))

	//创建3长度、4容量的切片
	s5 := make([]int, 3, 4)
	fmt.Printf("type=%T,value=%v,len=%d,cap=%d\n", s5, s5, len(s5), cap(s5))


	demob2()


}
