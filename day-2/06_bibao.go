package main

import "fmt"

func main() {
	f := test()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
func test() func() int {
	var a int
	return func() int {
		a++
		i := a * a
		return i
	}
}
