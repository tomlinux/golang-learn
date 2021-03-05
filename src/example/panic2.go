package main

import (
	"errors"
	"fmt"
)

/*
1. 以返回错误替代恐慌
概述

通过恐慌报错的方式虽然直白有效，但动不动就崩溃显得有些暴力
Go语言还给我们提供了一种相对温和但同样有效的异常解决方案，那就是同时返回结果和错误（error实例）
如果结果正确时错误为空，如果错误不为空时结果为空（或没有业务意义的默认值）
这种方式显得温和而辩证，兼容性好，也不会造成程序崩溃
至于究竟是严厉好还是温和兼容好，开发者们可以见仁见智
下面实例中的圆面积计算函数中：

如果调用者传入了一个负数半径，程序也不会panic，但是会返回一个提示错误的error实例，此时结果是毫无意义的默认值0
如果调用者传入了合法的半径，则返回正确结果和一个为空的error实例

 */

func main() {
	//获得调用的结果
	ret, err := getCircleAreaII(5)

	//如果返回的错误不为空，则说明程序出现了异常
	if err!=nil{
		fmt.Println(err)
	}else {
		fmt.Println("ret=",ret)
	}
}

/*
如果发生异常时，以返回错误的方式代替恐慌
正确结果和错误，必有一个为空
这样相对温和，也不会造成程序崩溃
*/
func getCircleAreaII(radius float32)(ret float32,err error) {
	if radius < 0{
		err = errors.New("傻鸟，半径不能为负")
		return
	}
	ret = 3.14 * radius * radius
	return
}