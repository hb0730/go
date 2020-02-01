package main

import "fmt"

func main() {
	a := 10
	str := "mike"
	defer func(s int, s1 string) {
		fmt.Printf("内部 a = %d , str = %s \n", s, s1)
	}(a, str)
	a = 20
	str = "hello"
	fmt.Printf("外部 c = %d , d = %s \n", a, str)
}
