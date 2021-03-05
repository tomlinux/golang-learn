package benchmark

import "fmt"

//压力测试概述
//压力测试用来检测函数(方法）的性能，和编写单元功能测试的方法类似，但需要注意以下几点：
//
//文件名命名规则：xxx_test.go
//函数名命名规则：func BenchXxx(b *testing.B),其中XXX可以是任意字母数字的组合，但是首字母不能是小写字母
//函数内必须使用b.N进行轮询测试
//函数内可以选择使用b.ReportAllocs()汇报内存开销
//在GoLandIDE中你可以在待测包上右键，Run->gobentch xxx，以执行整包的压力测试，默认从上向下依次执行所有
//终端执行当前包下的所有压力测试：
//go test -bench=.
//终端执行多次求平均值
//go test -bench=. -count=3

//å获取斐波那契数列第n项的递归实现
//1,1,2,3,5,8,13,21,34,55
func GetFibonacci1(n int) int {
	if n == 0 || n == 1 {
		return 1
	} else {
		return GetFibonacci1(n-1) + GetFibonacci1(n-2)
	}
}


//获取斐波那契数列第n项的非递归实现
//1,1,2,3,5,8,13,21,34,55
func GetFibonacci2(n int) int {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}

func  main()  {
	fmt.Println("hello")
}

