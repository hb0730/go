package main

import "fmt"

/**
切片创建方式
*/
func main() {
	//自动推到
	s := []int{1234}
	fmt.Println("s = ", s)
	//借助make函数 make[切片类型,长度,容量]；
	s1 := make([]int, 10)
	fmt.Println("s1=", s1)

	//未指定容量默认相等
	s2 := make([]int, 5)
	fmt.Println("s2 = ", s2)

}
