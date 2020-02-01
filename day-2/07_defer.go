package main

import (
	"fmt"
)

/**
defer
函数关闭前操作
*/
func main() {
	fmt.Println("aaaaaaaaaaaaa")
	defer fmt.Println("bbbbbbbbbbbbbbb")
	fmt.Println("cccccccccccccccc")
	defer test2(0)
	test()
}

func test() {
	defer fmt.Println("ddddddd")
	defer fmt.Println("eeeeeeeeeee")
	defer fmt.Println("fffffffffff")

}

func test2(i int) {
	result := 1 / i
	fmt.Println("ggggggggggggggggg", result)
}
