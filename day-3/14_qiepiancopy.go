package main

import "fmt"

/**
切片copy
*/
func main() {
	array := []int{1, 2, 3, 4}
	s := make([]int, 4)
	copy(s, array)
	fmt.Println("array = ", array)
	fmt.Println("s = ", s)
}
