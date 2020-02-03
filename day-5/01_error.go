package main

import (
	"errors"
	"fmt"
)

/**
error异常
*/
func main() {
	//test1();
	//test2();
	test3()
}

/**
error定义
*/
func test1() {
	error1 := fmt.Errorf("%s", "this is normal error1")
	fmt.Println(error1)

	error2 := errors.New("this is normal error2")
	fmt.Println(error2)
}

/**
panic定义
*/
func test2() {
	panic("error test3")
}

/**
recover 定义
*/
func test3() {
	defer func() {
		i := recover()
		if i != nil {
			fmt.Println(i)
		}
	}()
	panic("error test3")

	fmt.Println("ssss")
}

type errorMessage struct {
	error
	message string
}

func (message *errorMessage) Error() string {
	return message.message
}
