package main

import "fmt"

/**
循环
*/
func main() {
	for i := 1; i < 100; i++ {
		fmt.Println("i = ", i)
	}
	//迭代
	str := "abc"
	for i, data := range str {
		fmt.Printf("str[%d]=%c\n", i, data)
	}
}
