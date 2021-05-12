package main

import (
	"fmt"
	"strconv"
)

//
//// string到int64
//int0_64, err := strconv.ParseInt(string0, 10, 64)
//
//// int到string
//int0 := 123
//string1 := strconv.Itoa(int0)
//
//// int64到string
//string1_64 := strconv.FormatInt(int64,10)

//var  string0 string = "hello world"
//
//
//var int1, _ = strconv.Atoi(string0)
//
//var{
//	string0 string
//	int1  strconv.Atoi(string0)
//}

var err error

var string2 string = "hello world"

func main() {

	// string到int

	//
	//fmt.Println(int1,string0)

	// string到int
	int, err := strconv.Atoi(string2)
	if err != nil {
		fmt.Println(err.Error())
	}
	// string到int64
	int64, _ := strconv.ParseInt(string2, 10, 64)

	// int到string
	string3 := strconv.Itoa(int)

	// int64到string
	string4 := strconv.FormatInt(int64, 10)

	fmt.Println(int, int64, string3, string4)

}
