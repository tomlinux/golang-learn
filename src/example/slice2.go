package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

//判断切片相等
func AreEquivalentSlice(temp, answers []interface{}) bool {
	if len(temp) != len(answers) {
		return false
	} else {
		for i, v := range temp {
			//先判断类型是否相同
			if reflect.TypeOf(v) != reflect.TypeOf(answers[i]) {
				return false
			}

			if v != answers[i] {
				return false
			}
		}
	}
	return true
}

/*切片洗牌*/
/*打乱任意类型切片*/
func ShuffleSlice(slice []interface{}) []interface{} {
	var newSlice []interface{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for {
		if len(newSlice) == len(slice) {
			break
		}

		index := r.Intn(len(slice))
		value := slice[index]
		if value != nil {
			newSlice = append(newSlice, value)
			slice[index] = nil
		}
	}
	return newSlice
}
//func ShuffleIntSlice(slice []int)  []interface{} {
//
//	fmt.Println("hello world")
//}



/*获得指定范围的乱序切片*/
//func GetShuffledIntSlice(start, end int) []int {
//	var slice []interface{}
//	for i := start; i < end; i++ {
//		slice = append(slice, i)
//	}
//	//得到slice乱序切片
//	slice = ShuffleSlice(slice)
//	return slice
//}


func main() {
	arr1 := make([]interface{}, 3, 5)
	//var arr2 = []int{1, 2}
	arr3 := make([]interface{}, 3, 6)
	fmt.Println(arr1, arr3)

	//var arr4 = []int{1,3}
	//var arr5 = make([]interface{},3,6)

	fmt.Println(AreEquivalentSlice(arr1, arr3))

	//fmt.Println(AreEquivalentSlice(arr4,arr5))
	//测试
	var dataSlice []int = []int{1, 2, 3, 3}
	var interfaceSlice []interface{} = make([]interface{}, len(dataSlice))
	for i, d := range dataSlice {
		interfaceSlice[i] = d
	}
	fmt.Printf("dataSlice类型：%T,%#v \n",dataSlice,dataSlice)
	fmt.Printf("interfaceSlice类型：%T,%#v \n",interfaceSlice,interfaceSlice)
	fmt.Println(reflect.TypeOf(dataSlice),reflect.TypeOf(interfaceSlice))

}
