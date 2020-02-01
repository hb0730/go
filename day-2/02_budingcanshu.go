package main

import "fmt"

/**
不定参数
*/
func main() {

	test(1, 2, 3, 4)

}
func test(args ...int) {
	fmt.Println("元素大小:", len(args))
	for e := range args {
		fmt.Println("参数 arg=", e)
	}
	fmt.Println("参数传递")
	test2(args...)
	fmt.Println("从下标0-2,传递")
	test2(args[:2]...)
	fmt.Println("从小标2(不包含)之后的传递")
	test2(args[2:]...)
}
func test2(args ...int) {
	for i, e := range args {
		fmt.Printf("arg[%d] values %d \n", i, e)
	}
}
