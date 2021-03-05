package main

import (
	"FirstProject/tools"
	"fmt"
	"math/rand"
	"net"
	"time"
)

const (
	//iota 常量iota 是从0 开始自动增长
	Sunday = iota //0
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

var (
	a int = 123
	b int = 456
)

const (
	A = iota * 2
	B
	C
)

func sum(a, b int) int {
	return a + b
}

//1. golang 函数的使用方法
func main01() {
	//fmt.Print("hello")
	fmt.Print(tools.Sum(1, 2))
}

//2. 数组的声明和数组的变量方法
func main02() {
	a1, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(a1)
		// 数组遍历
		for index, num := range a1 {
			fmt.Println(index, "=>", num)
		}
	}

}

//3. iota 和变量的使用方法
func main03() {
	fmt.Println(Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday)
	fmt.Println(a, b)
	fmt.Println("========")
	fmt.Println(A, B, C)
}

//4. 标准输入的用法 。eg： 博彩游戏
/*规则如下
1.准备好money 和人 （团队)  。依次开始下注金额
2.
 */
func main() {
	var money int
	var team string
	fmt.Println("请输入你下注的金额和团队！格式是下注金额+团队")
	fmt.Scanf("%d+%s", &money, &team)
	fmt.Printf("您下注了%s%d万元，人生巅峰即将开始...\n",team,money)
	time.Sleep(1 * time.Second)
	//time.Now().Unix()获取当前时间距离1970年零时逝去的秒数
	//rand.NewSource(time.Now().Unix())每秒更新一个随机数的种子,一旦种子变化随机数也随之变化
	fmt.Println(time.Now().Unix())
	myrand := rand.New(rand.NewSource(time.Now().Unix()))
	//获得[0-100)的随机数
	luckyNumber := myrand.Intn(100)
	fmt.Println("luckyNumber=", luckyNumber)

	if luckyNumber > 10 {
		fmt.Println("靠海别野欢迎你")
	} else {
		fmt.Println("天台欢迎你")
	}



}
