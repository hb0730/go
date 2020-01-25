package main

import "fmt"

/**
复数
*/
func main() {
	var s1 complex128
	s1 = 3.2 + 3.14i
	fmt.Println("s1 = ", s1)
	var s2 = 3.2 + 3.14i
	fmt.Printf("s2 type is %T\n", s2)

	fmt.Printf("real(t2) = %f,imag(t2) = %f", real(s2), imag(s2))
}
