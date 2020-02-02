package main

import "fmt"

/**
切片追加
*/
func main() {
	array := []int{}
	array = append(array, 1)
	array = append(array, 2)
	array = append(array, 3)
	fmt.Println("array = ", array)
}
