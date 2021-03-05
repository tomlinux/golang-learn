package main

import "fmt"

func getResult(a float32, b float32, operator string)  (ret float32) {
	/*
	   单分支
	   if operator=="+"{
	       ret = a+b
	   }
	*/

	/*
	   //双分支
	   if operator=="+"{
	       ret = a+b
	   }else {
	       fmt.Printf("不支持的操作符：%s\n",operator)
	   }
	*/

	//多分支
	if operator == "+" {
		ret = a + b
	} else if operator == "-" {
		ret = a - b
	} else if operator == "*" {
		ret = a * b
	} else if operator == "/" {
		ret = a / b
	} else if operator == "%" {
		ret = float32(int(a) % int(b))
	} else {
		fmt.Printf("不支持的操作符：%s\n", operator)
	}

	return
}
func getResult2(a float32, b float32, operator string) (ret float32) {
	//判断operator的取值
	switch operator {

	//情形1,2,3...
	case "+":
		ret = a + b
	case "-":
		ret = a - b
	case "*":
		ret = a * b
	case "/":
		ret = a / b
	case "%":
		ret = float32(int(a) % int(b))

	//如果不符合上述任何一种情形
	default:
		fmt.Printf("不支持的操作符:%s\n", operator)
	}
	return
}


func  main()  {
	fmt.Println(getResult2(1.2,3,"+"))
}
