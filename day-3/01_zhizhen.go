package main

import "fmt"

/**
指针
*/
func main() {
	var a int = 10
	//每个变量有2个含义: 变量的内存,变量的地址
	fmt.Printf("a = %d \n", a)
	fmt.Printf("&a = %v \n", &a)

	//保存变量的内存地址,需要指针类型，*int 保存int地址,**int 保存，*int地址
	var p *int
	p = &a
	fmt.Printf("p = %v,&a = %v \n", p, &a)
	*p = 666
	fmt.Printf("*p = %v , a = %v \n", *p, a)
}
