package main

import (
	"fmt"
)

func main(){
	fmt.Println("hello，world")
	v :=1
	fmt.Printf("%d类型%T,值是%v \n",v,v,v)
	// errors.New("sffff")
	const (
		x  uint16 = 120
		y
		s  = "abc"
		z
	)
	fmt.Printf("类型%T 值是%v \n",v,v)
	fmt.Printf("类型%T 值是%v \n",z,z)
}











