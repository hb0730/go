package main

import "fmt"

/**
输入
*/
func main() {
	var a int
	fmt.Println("请输入:")
	//fmt.Scanf("%d", &a)
	//fmt.Scan(&a);
	fmt.Scanln(&a)
	fmt.Println("a =", a)
}
