package main

import "fmt"

/**
二维数组
*/
func main() {
	//test();
	//test1();
	test2()
}

func test() {
	var a [5][5]int
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			a[i][j] = i + j
		}
	}
	fmt.Println("a = ", a)
}

/**
初始化
*/
func test1() {
	b := [5][5]int{
		{1, 2, 3, 4, 5}, {}, {}, {}, {}}

	fmt.Println("b = ", b)
}

/**
指定初始化
*/
func test2() {
	c := [5][5]int{2: {3: 10}}

	fmt.Println("c = ", c)
}
