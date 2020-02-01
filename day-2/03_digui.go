package main

import "fmt"

/**
递归函数
*/
func main() {
	result := test(100)
	fmt.Println("result =", result)
}

func test(a int) (result int) {
	if a == 1 {
		result = 1
		return
	}
	return a + test(a-1)
}
