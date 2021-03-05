package main

import (
	"flag"
	"fmt"
	"os"
)

/*
命令行参数概述

* 没有GUI界面的应用程序怎么接收用户的输入呢
* 除了运行时接收标准用户输入以外，主要就是通过命令行参数了
* Go语言提供了读取一揽子所有命令行字符串、读取关键字参数两种方式
* 其中按名称读取关键字参数是主要的用法，需要大家牢牢掌握


*/
func demo51() {
	//os.Args中封装了用户启动当前程序所使用的所有字符串
	fmt.Println(os.Args)

	//遍历和处理用户在命令行中输入的每一个字符串
	for index, value := range os.Args {
		fmt.Println(index, value)
	}
}

/*
方式2：使用系统提供的参数地址
*/
func demo52() {
	//05xxx.exe -name bill -age 60 -money 1234567890.12 -isstupid true

	//nme=参数名字，value=参数的默认值，usage=参数的说明信息，返回值=存储了【用户输入的参数】的内存地址
	namePtr := flag.String("name", "无名氏", "用户名")
	agePtr := flag.Int("age", 0, "年龄")
	moneyPtr := flag.Float64("money", 0, "财富")
	isstupidPtr := flag.Bool("isstupid", true, "是否愚蠢")

	//解析参数
	flag.Parse()

	//从上边返回的地址中读取【用户输入的参数】
	fmt.Println(*namePtr, *agePtr, *moneyPtr, *isstupidPtr)
}

/*
方式3：使用自定义的参数地址
*/
func demo53() {
	//05xxx.exe -name bill -age 60 -money 1234567890.12 -isstupid true
	//开辟4块内存，用于存储cmd-line参数
	var name string
	var age int
	var money float64
	var isstupid bool
	fmt.Println(&name, &age, &money, &isstupid)

	//分别从命令行读取name，age,money,isstupid参数，存入上边开辟的内存地址中
	flag.StringVar(&name, "name", "无名氏", "用户名")
	flag.IntVar(&age, "age", 0, "年龄")
	flag.Float64Var(&money, "money", 0, "财富")
	flag.BoolVar(&isstupid, "isstupid", true, "是否愚蠢")

	//解析命令行参数
	flag.Parse()

	//得到结果
	fmt.Println(name, age, money, isstupid)
}

/*
Example:
人格测试小游戏
根据用户输入的姓名，年龄，财富，判断用户的人格
 */
func demo54() {
	name, age, money, isstupid := getUserInfoFromCmdline()

	//平均水平：以【40岁有50万标准】=1.0
	const k = 1 / (5.0e5 / 40)
	//成功程度，与money成正比，与age成反比,算出成功指数
	successIndex := k * money / float64(age)

	//富有且认为自己愚蠢的，加分，贫穷且认为聪明的减分
	if money > 1.0e6 && isstupid {
		successIndex *= 1.1
	}
	if money < 1.0e4 && !isstupid {
		successIndex *= 0.9
	}

	//得到评语
	var remark string
	if successIndex > 1 {
		remark = "棒棒的"
	} else {
		remark = "呵呵"
	}

	//姓名用于输出报告
	fmt.Printf("尊敬的%s阁下：你的成功指数是%.2f,%s\n", name, successIndex, remark)
}

//从命令行读取用户信息
func getUserInfoFromCmdline() (name string, age int, money float64, isstupid bool) {
	//05xxx.e -name bill -age 60 -money 1234567890.12 -isstupid true

	//nme=参数名字，value=参数的默认值，usage=参数的说明信息，返回值=存储了【用户输入的参数】的内存地址
	namePtr := flag.String("name", "无名氏", "用户名")
	agePtr := flag.Int("age", 0, "年龄")
	moneyPtr := flag.Float64("money", 0, "财富")
	isstupidPtr := flag.Bool("isstupid", true, "是否愚蠢")

	//解析参数
	flag.Parse()

	//从上边返回的地址中读取【用户输入的参数】
	fmt.Println(*namePtr, *agePtr, *moneyPtr, *isstupidPtr)
	return *namePtr, *agePtr, *moneyPtr, *isstupidPtr
}



func main()  {

	demo54()
	
}