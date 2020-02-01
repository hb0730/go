package test

import "fmt"

func Add(a, b int) int {
	return a + b
}
func Minus(a, b int) int {
	return a - b
}

/**
初始化函数,自动调用
*/
func init() {
	fmt.Println("初始化test")
}
