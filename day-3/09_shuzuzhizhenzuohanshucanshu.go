package main

import "fmt"

/**
数组指针做函数参数
*/
func main() {
	array := [5]int{1, 2, 3, 4, 5}
	modify(&array)
	fmt.Println("main=", array)
}

/**
地址传递
*/
func modify(p *[5]int) {
	(*p)[0] = 8
	fmt.Println("modify=", *p)
}
