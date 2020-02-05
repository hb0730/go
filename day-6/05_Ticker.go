package main

import (
	"fmt"
	"time"
)

/**
周期性 定时任务
*/
func main() {
	test1()
}

/**
定义
*/
func test1() {
	t := time.NewTicker(2 * time.Second)
	i := 0
	for {
		<-t.C
		i++
		fmt.Println("i = ", i)
		if i == 5 {
			t.Stop()
			fmt.Println("停止ticker并退出")
			break
		}
	}
}
