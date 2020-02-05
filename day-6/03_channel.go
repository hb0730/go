package main

import (
	"fmt"
	"time"
)

/**
协程之间同步
*/
func main() {
	//test1()
	//test2()
	//test3()
	//test4()
	test5()
}

//定义全局chan
var ch = make(chan int)

func printer(str string) {
	for _, data := range str {
		fmt.Printf("%c", data)
		time.Sleep(time.Second)
	}
	fmt.Printf("\n")
}
func p1() {
	printer("hello")
	ch <- 666
}
func p2() {
	//读取chan中数据，如果chan中无数据就会堵塞
	<-ch
	printer("word")
}
func test1() {
	go p1()
	go p2()
	for {

	}
}

/**
实现主协程结束
*/
func test2() {
	go func() {
		defer fmt.Println("子协程执行完毕")
		for i := 0; i < 5; i++ {
			fmt.Println("i = ", i)
		}
		ch <- 10
	}()

	int := <-ch
	fmt.Println("int ", int)
	defer fmt.Println("主协程结束")
}

/**
chanel 缓存
*/
func test3() {
	ch := make(chan int, 10)
	fmt.Printf("len(ch) = %d, cap(ch) = %d", len(ch), cap(ch))

	go func() {
		for i := 0; i < 14; i++ {
			ch <- i
			fmt.Println(" i =", i)
		}
	}()

	time.Sleep(time.Second)
	for i := 0; i < 14; i++ {
		in := <-ch
		fmt.Println("in = ", in)
	}
}

/**
关闭channel
*/
func test4() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
			fmt.Println(" i =", i)
		}
		close(ch)
	}()
	for {
		//如果是ok,管道未关闭
		//if in,ok := <-ch;ok==true{
		//	fmt.Println("in = ", in)
		//}else {
		//	break;
		//}
		for data := range ch {
			fmt.Println("in = ", data)
		}
	}
}

/**
单向channel
*/
func test5() {
	ch := make(chan int)
	go producer(ch)
	consumer(ch)
}

/**
生成者
*/
func producer(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i * i
	}
	close(ch)
}

/**
消费
*/
func consumer(in <-chan int) {
	for i := range in {
		fmt.Println("in = ", i)
	}
}
