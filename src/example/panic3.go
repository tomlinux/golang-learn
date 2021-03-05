package main

import (
	"fmt"
	"time"
)

/*
自定义异常概述

*系统提供的异常类error无所不包但控制的粒度很粗糙
*如果想要进行很精细化的报错和处理错误，就需要用到自定义异常
*我们只要实现error接口中的Error() string方法，就可以自定义异常了
*自定义的异常与系统异常无二，同样可以panic，同样可以辩证互斥地返回【结果错误对】

自定义异常
下面的例子定义了一个【非法参数异常】
在这个自定义的结构体中，封装了异常提示信息、发生时间、造成异常的用户三个属性
这样的异常在报出和处理时，无疑具有更直观更丰富的参考价值

*/

//自定义【非法参数异常】，封装异常提示信息、发生时间、造成异常的用户三个属性
type InvalidArgError struct {
	info string
	when time.Time
	user string
}

//实现系统的异常接口
func (iae *InvalidArgError)Error() string{
	return fmt.Sprintln(iae.info,iae.when,iae.user)
}

//模仿SDK，提供一个创建异常对象的工厂方法，自动记录发生异常的时间和用户
func NewInvalidArgError(info string) *InvalidArgError {
	iaePtr := new(InvalidArgError)
	iaePtr.info = info
	iaePtr.when = time.Now()
	iaePtr.user = "bill"
	return iaePtr
}


//以错误的形式返回异常信息
//如果参数为负数时，返回一个自定义的异常
func getCircleAreaIII(radius float32)(ret float32,err interface{}) {
	if radius < 0{
		//err = errors.New("傻鸟，半径不能为负")
		err = NewInvalidArgError("傻鸟，半径不能为负")
		return
	}
	ret = 3.14 * radius * radius
	return
}

//返回自定义异常
func demo41() {
	ret, err := getCircleAreaIII(-5)
	if err != nil {
		//fmt.Println(err)
		fmt.Println(err)
	} else {
		fmt.Println("ret=", ret)
	}
}

func main() {
	demo41()
}

