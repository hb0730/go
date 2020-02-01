package main

import "fmt"

func main() {
	test()
	test2("hello")
	s := test3("有返回值")
	fmt.Println(s)
}

/**
无参无返回值
*/
func test() {
	fmt.Println("无参无返回值")
}

/**
有参无返
*/
func test2(b string) {
	fmt.Printf("有参无返回值%s \n", b)
}

/**
有参有返
*/
func test3(b string) string {
	fmt.Printf("有参有返回值%s \n", b)
	return b
}

/**
多值返回
*/
func test4() (result1 string, result2 string) {
	result1, result2 = "result1", "result2"
	return
}
