package main

import (
	"fmt"
	"regexp"
)

/**
正则表达式
*/
func main() {
	test1()
}

func test1() {
	var reg string = "abb a1b a2b aqb acb"
	compile := regexp.MustCompile("a.b")
	allString := compile.FindAllString(reg, -1)
	fmt.Println(allString)
}
