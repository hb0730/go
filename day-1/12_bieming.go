package main

import "fmt"

/**
别名
*/
func main() {
	type bigint int
	var a bigint
	a = 10
	fmt.Printf("a type is %T,a = %d \n", a, a)

	type (
		minint int
		b      bool
	)
	var s1 minint
	var s2 b
	s1 = 1
	s2 = false
	fmt.Printf("s1 type is %T,s2 type is %T \n", s1, s2)
}
