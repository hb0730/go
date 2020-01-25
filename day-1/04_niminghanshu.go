package main

import "fmt"

/**
匿名函数
*/
func test() (a, b int) {
	fmt.Println("test function")
	return 10, 20
}

/**
匿名函数
*/
func main() {
	// 多重赋值
	a, b := 10, 20
	fmt.Printf("a = %d,b=%d \n", a, b)
	//交换
	a, b = b, a
	fmt.Printf("a = %d,b=%d \n", a, b)

	//_匿名函数
	a, _ = test()

	fmt.Printf("a = %d", a)
}
