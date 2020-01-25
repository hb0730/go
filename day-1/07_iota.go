package main

import "fmt"

func main() {
	// iota 每一行++
	const (
		a = iota
		b = iota
		c = iota
	)
	fmt.Printf("a = %d,b =%d,c =%d \n", a, b, c)
	const (
		d    = iota
		e, f = iota, iota
		g    = iota
	)
	fmt.Printf("d = %d,e = %d, f = %d, g = %d \n", a, e, f, g)
	// iota 遇见 const 则为0
	const name = iota
	fmt.Printf("name = %d", name)
}
