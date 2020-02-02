package main

import "fmt"

/**
数组做函数参数传递
实则为copy,不会改变原来
*/
func main() {
	array := [5]int{5, 4, 3, 4, 1}
	ints := modify(array)
	fmt.Println("main = ", array)
	fmt.Println("modify = ", ints)
}
func modify(array [5]int) [5]int {
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}
