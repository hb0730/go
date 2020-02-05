package main

import (
	"fmt"
	"time"
)

/**
timer 定时器
*/
func main() {
	//test1()
	test2()
}

/**
定义
*/
func test1() {
	//三秒后写入输数据,并且只响应一次
	t := time.NewTimer(time.Second)
	fmt.Println("当前时间", time.Now())
	time := <-t.C
	//当前时间
	fmt.Println("time =", time)
}

/**
实现延迟
*/
func test2() {
	fmt.Println("延迟两秒")
	t := time.NewTimer(2 * time.Second)
	<-t.C
	ddd()
}
func ddd() {
	fmt.Println("ddd method")
}
