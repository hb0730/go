package main

import (
	"fmt"
)

func main() {
	// 布尔类型
	boolTest()
	intTest()
	floatTest()
	stringTest()
	byteTest()
}
func byteTest() {
	fmt.Println(" byte type")
	var a byte
	fmt.Println(" a =", a)
	a = 'A'
	fmt.Println(" a =", a)

}
func stringTest() {
	fmt.Println(" string type")
	var a string
	fmt.Println(" a = ", a)
	a = "hello"
	fmt.Println(" a = ", a)

}
func floatTest() {
	fmt.Println("float type")
	var a float64
	fmt.Println("a =", a)
	a = 12.3
	fmt.Println(" a = ", a)
}
func intTest() {
	fmt.Println(" int type")
	var a int
	fmt.Println(" a = ", a)
}

func boolTest() {
	var a bool
	fmt.Println(" a = ", a)
	a = true
	fmt.Println(" a= ", a)

	b := false
	fmt.Println("b = ", b)
	var c = true
	fmt.Println(" c = ", c)

}
