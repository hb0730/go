package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**
字符串处理
*/
func main() {
	//test1();
	test2()
}

/**
字符串操作
*/
func test1() {
	//是否包含每个字符串
	contains := strings.Contains("seafood", "food")
	fmt.Println("是否包含某个字符串", contains)
	contains = strings.Contains("seafood", "xxxx")
	fmt.Println("是否包含某个字符串", contains)
	//拼接
	str := []string{"foo", "bar", "baz"}
	s := strings.Join(str, ",")
	fmt.Println("拼接", s)
	s = strings.Join(str, ":")
	fmt.Println("拼接", s)
	//查找小标index
	index := strings.Index("seafood", "food")
	fmt.Println("下标", index)
	index = strings.Index("seafood", "x")
	fmt.Println("下标", index)
	//repeat 返回count个s串联的字符串。
	repeat := strings.Repeat("seafood", 2)
	fmt.Println("串联", repeat)
	repeat = strings.Repeat("seafood", 0)
	fmt.Println("串联", repeat)
	//replace 返回将s中前n个不重叠old子串都替换为new的新字符串，如果n<0会替换所有old子串。
	replace := strings.Replace("sea sea sea ", "a", "afood", 2)
	fmt.Println("替换", replace)
	replace = strings.Replace("sea sea sea ", "a", "afood", 0)
	fmt.Println("替换", replace)
	replace = strings.Replace("sea sea sea ", "a", "afood", -1)
	fmt.Println("替换", replace)
	//split 切割
	split := strings.Split("seafood,seafood,seafood", ",")
	fmt.Println("切割", split)
	//trim 去除前后两端特定的字符串
	trim := strings.Trim("aseafooda", "a")
	fmt.Println("去除", trim)
	trim = strings.TrimSpace(" seafood ")
	fmt.Println("去除空格", trim)
	//fields 返回将字符串按照空白（unicode.IsSpace确定，可以是一到多个连续的空白字符）分割的多个字符串。如果字符串全部是空白或者是空字符串的话，会返回空切片。
	fields := strings.Fields(" seafood seafood seafood seafood ")
	fmt.Println("连续分割", fields)
}

/**
字符转换
*/
func test2() {
	//append 追加
	bytes := make([]byte, 5)
	quote := strconv.AppendQuote(bytes, "ssss")
	fmt.Println("追加", quote)

	//format 类型转换，其他类型转换成字符串
	formatInt := strconv.FormatInt(1234, 10)
	fmt.Println("其他类型转换", formatInt)
	//pare 将字符串转换成其他类型
	ok, _ := strconv.ParseBool("1")
	fmt.Println("字符串转换", ok)
}
