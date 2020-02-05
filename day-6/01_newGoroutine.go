package main

import (
	"fmt"
	"time"
)

func main() {
	//新建一个协程
	go test1()

	for {
		fmt.Println("this is main goroutine")
		time.Sleep(time.Second)
	}
}
func test1() {
	for {
		fmt.Println("this is test1 goroutine")
		time.Sleep(time.Second)
	}
}
