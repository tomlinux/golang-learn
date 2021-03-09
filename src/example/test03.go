package main

import "fmt"

func arrayModify(array [5]int) {
	newArray := array
	newArray[0] = 88
}
func sliceModify(slice []int) {
	newSlice := slice
	newSlice[0] = 88
}
func main() {
	array := [5]int{1, 2, 3, 4, 5}
	slice := []int{1, 2, 3, 4, 5}
	arrayModify(array)
	sliceModify(slice)
	fmt.Println(array)
	fmt.Println(slice)
}



