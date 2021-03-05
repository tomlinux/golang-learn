package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"strconv"
)

// 1. 字符串转换Ascll编码
func main001() {
	key := "1234567890"
	// byte 字符串转换ASCLL编码
	var data = []byte(key)
	fmt.Print(data, "\n")
	fmt.Println("key = ", key)
	i := 0
	for ; i < len(key); i++ {
		fmt.Printf("key[%d]=%x  ", i, key[i])
	}
	fmt.Println()

	fmt.Println("key = ", key)

	for i = 0; i < len(key); i++ {
		fmt.Printf("data[%d]=%x  ", i, data[i])
	}
	fmt.Println()

	fmt.Println("end")

}

//总结：
//字符串key直接取数组下标是可以正常执行的，因此go中不需要像java一样使用.getByte()取字节数据
//当然，go中也可以使用
//var byte_data = []byte(string_data)
//将string数据转换成byte类型

//2.字符串转换成数组，将两个字符拼接成一个字节
func main002() {
	key := "12345678"
	//arr := bytes.NewBufferString(key)
	arr, err := hex.DecodeString(key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("arr", arr)
	i := 0
	for i = 0; i < len(arr); i++ {
		fmt.Printf("arr[%d]=%x  ", i, arr[i])
	}

}

// 3. 数组转换成字符串，字节的十六进制字节变成字符
func main003() {
	i := 0
	//  数组和切片的区别是数组是有容量cap 切片没有
	//数组,
	arr := [...]byte{0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38}
	//切片
	slice := []byte{0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38}

	//数组用字符形式打印，即可变成字符串输出
	for ; i < len(arr); i++ {
		fmt.Printf("%c", arr[i])
	}

	fmt.Println()

	fmt.Println("arr:", arr)
	fmt.Println("slice:", slice)

	//arr_Str := hex.EncodeToString(arr)
	//fmt.Println("arr_str = ",arr_Str)
	var SliceStr string = hex.EncodeToString(slice)
	fmt.Println("slice_str = ", SliceStr)
	fmt.Printf("切片的容量是：%d,数组的容量是%d ", cap(arr), cap(slice))

	array3 := [...]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Print(array3, len(array3), cap(array3))

	var slice3 = make([]int, 3, 4) //  make 一个参数是类型。二个是参数是长度，三个参数是容量
	fmt.Print(slice3, len(slice3), cap(slice3))

}

// 4. 数组转换成字符串，两个字节拼接成一个字符
//前提：字节数组为字符0~F所对应的ASCLC码表示
func main004() {
	slice := []byte{0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38}
	fmt.Println(string(slice))
}

// 5.string和int类型的转化
func main005() {
	//string转成int：
	string := "1234567891011121314"
	int, _ := strconv.Atoi(string)
	//string转成int64：
	int64, _ := strconv.ParseInt(string, 10, 64)
	//int转成string：
	string2 := strconv.Itoa(int)
	//int64转成string：
	string3 := strconv.FormatInt(int64, 10)
	fmt.Print(int,"#",int64,"#",string2,"#",string3)
}
