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
将待排数组，视为已排序部分（最开始是第一个元素）+未排序部分，然后从未排序部分中逐个抽取元素，插入已排序部分中的正确位置

插入排序：
将待排序列视为排好和未排两部分组成的，排序的过程相当于不断从未排序列中抽取元素，插入到排好序的序列中
未排序：3,1,4,5,9,2,6,8,7,0，视为：			（已排序部分）3 + （未排序部分）1,4,5,9,2,6,8,7,0
第一步：从未排部分抽取1插入已拍序列中，得到：	（已排序部分）1,3 + （未排序部分）4,5,9,2,6,8,7,0
第二步：从未排部分抽取4插入已拍序列中，得到：	（已排序部分）1,3,4 + （未排序部分）5,9,2,6,8,7,0
第三步：从未排部分抽取5插入已拍序列中，得到：	（已排序部分）1,3,4,5 + （未排序部分）9,2,6,8,7,0
以此类推
*/
func InsertSliceInsert(s []int) (slice []int) {
	/*将第0位元素视为已排好的序列，从第1位元素一直遍历到最后*/
	for i := 1; i < len(s); i++ {

		//临时保存当前待插入的元素
		temp := s[i]

		/*将temp与已排序好的序列中的元素逐一比较，并插入在正确的位置上*/
		for j := 0; j < i; j++ {

			/*将temp插入在正确的位置上（即j位置）*/
			if s[i] < s[j] {

				//将【s[j+1]...s[i-1]】向后拖动一位
				for k := i - 1; k >= j; k-- {
					s[k+1] = s[k]
				}

				//将5放置在正确的位置上
				s[j] = temp

				//跳出temp的插入逻辑，开始锁定下一个元素
				break
			}
		}
	}

	return

}

/*
如果说SortInsert函数使用了三层循环，不是很优雅的话，我们稍加改进后如下
*/
func SortInsert2(s []int) {
	for i := 1; i < len(s); i++ {
		temp := s[i]

		for j := i - 1; j >= 0; j-- {
			if s[j] > temp {
				//比temp大的都向后顺移
				s[j+1] = s[j]
			} else {
				//找到一个不大于temp的数，在它后面放置temp
				s[j+1] = temp

				//当前temp已经找到归宿，开始下一个待插入元素的定位
				break
			}

			//如果所有元素都比temp大，则把temp放在0号位置上
			if j == 0 {
				s[j] = temp
			}
		}
	}
}

/*
快速排序
对一个切片进行切片不会复制底层的数组，这意味着切片的子切片只是与原切片共享底层数组的一个视图。
因此，如果您想要将它与初始的切片分开需要使用copy()。
对于 append 函数，如果它没有足够的容量来保存新值，底层数组将会重新分配内存和大小。
这意味着 append 的结果能不能指向原始数组取决于它的初始容量。这会导致难以发现的不确定 bugs。
*/



func main() {
	var s1 = []int{1, 3, 2, 7, 8, 11, 9}
	beforeTime := time.Now()
	time.Sleep(time.Second * 1)
	fmt.Println(SortSliceSelected(s1))
	elapsed := time.Since(beforeTime)
	fmt.Println(elapsed)

	// 冒泡排序
	var s2 = []int{1, 3, 1, 1, 2, 3, 4, 10, 100, 23}
	fmt.Println(BubbleSliceSort(s2))

	// 插入排序
	s3 := []int{3, 1, 9, 5, 4, 2, 6, 8, 7, 0}
	SortInsert2(s3)
	fmt.Println(s3)

}
