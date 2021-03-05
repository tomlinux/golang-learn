package main

import "fmt"

/*

类型断言概述

• 还记得大明湖畔的多态吗？——一个父类接口可以有很多不同的子类形态，骑兵步兵都是战士
• 当我从一个战士的集合中随意抽取一员时，它事实上是骑兵还是步兵呢
• 我们需要使用类型断言来加以判断——类型断言就是类型判断（翻译Go的砖家可能语死早）
• 类型断言的具体方式有二
• 类型断言方法一
• 在Go语言中，我们可以使用type switch语句查询接口变量的真实数据类型，语法如下
 */

//1.类型断言方法一

func demo110() {
	//switch value := x.(type) {
	//case x:
	//default:
	//}
	//[]interface{}任意类型的slice
	buf := make([]interface{}, 3)
	buf[0] = 5
	buf[1] = 1.3
	buf[2] = "无兄弟不编程"
	fmt.Printf("buf类型是%T，value是%v \n", buf, buf)
	for i := 0; i < len(buf); i++ {
		switch value := buf[i].(type) {
		case int:
			fmt.Printf("%d\n", value)
		case float32:
		case float64:
			fmt.Printf("%f\n", value)
		case string:
			fmt.Printf("%s\n", value)
		default:
			fmt.Println("类型未知!")
		}

	}
}

/*
类型断言方法二
x.(T)
同样，x必须是接口类型。
str := value.(string)

面的转换有一个问题，如果该值不包含一个字符串，则程序会产生一个运行时错误。为了避免这个问题，
可以使用“comma, ok”的习惯用法来安全地测试值是否为一个字符串：

*/
type values interface {
	//新陈代谢：吃进来+排出去,shit就是翔啊~
	Eat(food string) (shit string)
	//GAME OVER
	Die()
}

func demo120()  {

	//str, ok := values.(string)
	//if ok {
	//	fmt.Printf("string value is: %q\n", str)
	//} else {
	//	fmt.Printf("value is not a string\n")
	//}
//如果类型断言失败，则str将依然存在，并且类型为字符串，不过其为零值，即一个空字符串。

}


func main()  {
	demo120()
	
}