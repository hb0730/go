package main

import "fmt"

/**
if 控制
*/
func main() {

	fmt.Printf("请输入")
	var s1 int
	fmt.Scanf("%d", &s1)
	if s1 == 1 {
		fmt.Println("s1 = 11111111")
	} else if s1 == 2 {
		fmt.Println("s2 = 222")
	} else {
		fmt.Println("other")
	}

	//初始化
	if a := 10; a == 10 {
		fmt.Println("a==10")
	} else if a == 20 {
		fmt.Println("a==20")
	} else {
		fmt.Println("a")
	}
}
