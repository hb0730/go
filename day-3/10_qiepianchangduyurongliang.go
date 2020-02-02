package main

import "fmt"

/**
切片长度与容量
*/
func main() {
	array := []int{1, 3, 4, 5, 6, 7}
	// array[初始值,终止值,容量]
	s := array[0:4:5]
	fmt.Println("s = ", s,
		"len = ", len(s), //长度 终止值-初始值
		"cap = ", cap(s)) //容量  容量大小-初始值
}
