package main

import "fmt"

func main() {
	var (
		a int
		b bool
	)
	a = 10
	b = false
	fmt.Printf("a = %d,b = %t \n", a, b)

	const (
		i         = "hello"
		j float64 = 12.4
		z         = 10
	)

	fmt.Printf("i = %s,j =%F,z = %d \n", i, j, z)
}
