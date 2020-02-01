package main

import "fmt"

/**
new 分配内存
*/
func main() {
	var p *int
	p = new(int)
	*p = 666
	fmt.Println("*p = ", *p)
}
