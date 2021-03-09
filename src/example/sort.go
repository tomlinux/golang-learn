package main

import (
	"fmt"
	"time"
)

/*
选择排序
选择排序的思路是，每锁定一把“交椅”，就把剩下元素中最适合的扶上这把”交椅”
*/

func SortSliceSelected(slice []int) []int {
	for i := 0; i < len(slice)-1; i++ {
		for j := i; j < len(slice); j++ {
			if slice[j] < slice[i] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
	return slice

}

/*
冒泡排序
      0  1  2 3  4 5  6  7
列表： 1，3，5，2，3，4，5，6  slice   len :=8-1 =7
i= 0  1  3  5 2  3 4 5 6
i= 1  1  3  5 2  3 4 5 6
i= 2

第二趟
*/

func BubbleSliceSort(slice []int) []int {
	/*冒泡排序,正序排列*/
	for i := len(slice) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if slice[j+1] < slice[j] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}


	return slice
}
/*
插入排序
 */




func main() {
	var s1 = []int{1, 3, 2, 7, 8, 11, 9}
	beforeTime := time.Now()
	time.Sleep(time.Second * 1)
	fmt.Println(SortSliceSelected(s1))
	elapsed := time.Since(beforeTime)
	fmt.Println(elapsed)

	// 冒泡排序
	var s2 = []int{1,3,1,1,2,3,4,10,100,23}
	fmt.Println(BubbleSliceSort(s2))






}
