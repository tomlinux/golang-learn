package main

import (
	"fmt"
	"time"
)
//goto SOMEWHERE就是：直接去到标记为SOMEWHERE地方，Go 语言的 goto 语句可以无条件地转移到过程中指定的行。
//goto语句通常与条件语句配合使用。可用来实现条件转移， 构成循环，跳出循环体等功能。但是，在结构化程序设计中一般不主张使用goto语句， 以免造成程序流程的混乱，使理解和调试程序都产生困难。
func main() {
	var i int
	for {
		if i > 10 {
			//去到MAKEOVER标记的地方
			goto MAKEOVER
		}
		fmt.Println(i)
		time.Sleep(500 * time.Millisecond)
		i++
	}

MAKEOVER:
	fmt.Print("游戏结束！请等下班🛬！")

}
