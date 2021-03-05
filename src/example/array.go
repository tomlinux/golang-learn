package main

import "fmt"

/* 数组概述
*  数组是长度固定，类型固定的数据容器
* 根据下标访问和修改元素内容
* 下标从0开始，最后一个元素的下标是长度减一
* 可以使用len(arr)获得数组的长度
 */
func main() {
	var a1 []int
	//开辟10长度的整型数组
	//[]  是切片
	fmt.Printf("%T %v \n",a1,a1)

	var a11 [10]int
	//[0 0 0 0 0 0 0 0 0 0]
	fmt.Println(a11)


	//开辟10长度的整型数组
	var a2 [10]int = [10]int{}
	//[0 0 0 0 0 0 0 0 0 0]
	fmt.Println(a2)

	//创建10长度的整型数组，并给前6项赋值
	var a3 = [10]int{0, 1, 2, 3, 4, 5}
	//[0 1 2 3 4 5 0 0 0 0]
	fmt.Println(a3)

	//根据实际元素的个数创建数组，其长度依然是固定不变的
	var a4 = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//[0 1 2 3 4 5 6 7 8 9]
	fmt.Println(a4)

	//同a4，只是声明方式不一样
	a5 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//[0 1 2 3 4 5 6 7 8 9]
	fmt.Println(a5)

	//通过range遍历数组元素
	//index为遍历中的元素下标，value为值
	for index, value := range a5 {
		fmt.Println(index, value)
	}

}
