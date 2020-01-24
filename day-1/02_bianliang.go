package main

import "fmt"

func main() {
	var a int
	fmt.Println("a = ", a)
	a = 10
	fmt.Println("a =", a)

	// 多变量声明 var a,b,int

	//初始化赋值

	var b int = 10
	fmt.Println("b=", b)

	//类型

	c := 10
	fmt.Printf("c type is %T", c)
}
