package main

import "fmt"

/**
全局变量
*/
var name string
var age int

func main() {
	name = "make"
	age = 19
	test()
}

func test() {
	fmt.Printf("name = %s, age = %d", name, age)
}
