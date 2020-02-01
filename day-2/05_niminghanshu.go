package main

import "fmt"

/**
匿名函数
*/
func main() {
	a := 10
	str := "mike"
	//
	f1 := func() {
		fmt.Println("a =", a)
		fmt.Println("str = ", str)
	}
	f1()
	// 匿名函数起别名
	type funcType func()
	var f2 funcType
	f2 = f1
	f2()

	//声明并同时调用
	func() {
		fmt.Printf("a = %d , str = %s \n", a, str)
	}()

	func(i int, j string) {
		fmt.Printf("i = %d, j = %s \n", i, j)
	}(a, str)

	x, y := func(i int, j string) (int, string) {
		fmt.Printf("i = %d, j = %s \n", i, j)
		return i, j
	}(a, str)
	fmt.Printf("x = %d,y = %s \n", x, y)
}
