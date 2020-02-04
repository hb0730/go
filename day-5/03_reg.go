package main

import (
	"fmt"
	"regexp"
)

/**
正则表达式
*/
func main() {
	//test1()
	test2()
}

/**
正则表达式
*/
func test1() {
	var reg string = "abb a1b a2b aqb acb gggg"
	compile := regexp.MustCompile("a.b")
	allString := compile.FindAllString(reg, -1)
	fmt.Println(allString)
}

/**
提取所有小数
*/
func test2() {
	var reg string = "3.14 asc 3.12 3.41 7. 9.gg 8.98"
	compile := regexp.MustCompile(`\d+\.\d+`)
	allString := compile.FindAllString(reg, -1)
	fmt.Println(allString)
}
