package main

import "fmt"

/**
数组
*/
func main() {
	//test();
	test1()
}

/**
数组定义
*/
func test() {
	var id [50]int
	for i := 0; i < len(id); i++ {
		id[i] = i + 1
		fmt.Printf("id[%d]=%d \n", i, id[i])
	}
}

/**
数组初始化
*/
func test1() {
	var a [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("a = ", a)

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(" b = ", b)

	//部分初始化
	c := [5]int{1, 2, 3}
	fmt.Println(" c = ", c)

	//指定初始化 2: 下标为2的初始化
	d := [5]int{2: 10}
	fmt.Println(" d = ", d)
}
