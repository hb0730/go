package main

import "fmt"

/**
接口
*/
func main() {
	//test1();
	test2()
}

type Humaner interface {
	SayHi()
}

type Student struct {
	id   int
	addr string
}

/**
定义
*/
func test1() {
	var i Humaner
	student := Student{id: 1, addr: "上海"}
	i = student
	i.SayHi()
	var s myStr = "h"
	i = s
	i.SayHi()

}

func (tmp Student) SayHi() {
	fmt.Printf("Student %d %s \n", tmp.id, tmp.addr)
}

type myStr string

func (tmp myStr) SayHi() {
	fmt.Printf("myStr %s \n", tmp)
}

/**
多态
*/
func test2() {
	student := Student{id: 1, addr: "上海"}
	WhoSayHi(student)
	var s myStr = "h"
	WhoSayHi(s)
}

func WhoSayHi(i Humaner) {
	i.SayHi()
}
