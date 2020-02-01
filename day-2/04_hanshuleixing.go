package main

import "fmt"

/**
函数类型
*/
func main() {
	var test funcTest
	test = Add
	i := test(1, 1)
	fmt.Println("i =", i)
}

func Add(a, b int) int {
	return a + b
}

type funcTest func(int, int) int
