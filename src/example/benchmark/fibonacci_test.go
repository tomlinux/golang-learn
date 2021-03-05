package benchmark
//导入测试工具包
import "testing"

//测试用例1：多次测试函数GetFibonacci1，获得平均执行时间

func BenchmarkGetFibonacci1(b *testing.B) {
	b.Log("BenchmarkGetFibonacci1")

	//汇报内存开销
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		GetFibonacci1(10)
	}
}

//测试用例2：多次测试函数GetFibonacci2，获得平均执行时间
func BenchmarkGetFibonacci2(b *testing.B) {
	b.Log("BenchmarkGetFibonacci2")

	//汇报内存开销
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		GetFibonacci2(10)
	}
}
