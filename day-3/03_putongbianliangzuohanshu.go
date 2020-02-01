package main

import "fmt"

/**
普通变量做函数
*/
func main() {
	a, b := 10, 20

	//通过函数交换a,b值
	// 地址传递
	swap(&a, &b)
	fmt.Printf("a = %d, b = %d", a, b)
}
func swap(p1, p2 *int) {
	*p1, *p2 = *p2, *p1
	fmt.Printf("*p1 = %v ,*p2 = %v \n", *p1, *p2)
}
