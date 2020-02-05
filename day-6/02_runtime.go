package main

import (
	"fmt"
	"runtime"
)

/**
runtime 包
*/
func main() {
	//test1();
	//test2();
	test3()
}

/**
runtime.Gosched 让出时间片
*/
func test1() {
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("go")
		}
	}()
	for i := 0; i < 2; i++ {
		//让出时间片,让其他先执行
		runtime.Gosched()
		fmt.Println("hello")
	}
}

/**
runtime.Goexit 退出当前协程
*/
func test2() {
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("go")
			if i == 4 {
				runtime.Goexit()
			}
		}
	}()
	for {
		fmt.Println("hello")
	}
}

/**
runtime.GOMAXPROTS 设置运行的cpu核数
*/
func test3() {
	in := runtime.GOMAXPROCS(1)
	fmt.Println("之前运算核数", in)
	for {
		go fmt.Print(1)
		fmt.Print(0)
	}
}
