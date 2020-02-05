package main

import (
	"fmt"
	"time"
)

/**
select 使用
*/
func main() {
	//test1()
	test2()
}
func test1() {
	// 保存数字
	ch := make(chan int)
	//是否结束
	quit := make(chan bool)
	go producer(ch, quit)
	consumer(ch, quit)
}
func producer(out chan<- int, quit <-chan bool) {
	x, y := 1, 1
	for {
		select {
		case out <- x:
			x, y = y, y+x
		case b := <-quit:
			if b == true {
				fmt.Println("已结束")
				return
			}
		}
	}
}
func consumer(int <-chan int, quit chan<- bool) {
	for i := 0; i < 8; i++ {
		mun := <-int
		fmt.Println("mun = ", mun)
	}
	quit <- true
}

/**
超时
*/
func test2() {
	ch := make(chan int)
	quit := make(chan bool)
	go func() {
		for {
			select {
			case mun := <-ch:
				fmt.Println("mun = ", mun)
			case <-time.After(3 * time.Second):
				fmt.Println("超时")
				quit <- true
			}
		}
	}()
	for i := 0; i < 5; i++ {
		ch <- i
	}
	<-quit
	defer fmt.Println("程序退出")
}
