package main

import "fmt"

/**
方法值与方法表达式
*/
func main() {
	test1()
}

/**
方法值
*/
func test1() {
	var s1 Person
	s1.age = 18
	//保存方法，方法值
	psetValue := s1.setValue
	psetValue()
	psetPointer := s1.setPointer
	psetPointer()
}

/**
方法表达式
*/
func test2() {
	var s1 Person
	s1.age = 18
	//方法表达式
	psetValue := Person.setValue
	psetValue(s1)

	psetPointer := (*Person).setPointer
	psetPointer(&s1)
}

type Person struct {
	name string
	sex  byte
	age  int
}

func (tmp Person) setValue() {
	tmp.name = "gg"
	fmt.Printf("setValue %p %v \n", &tmp, tmp)
}
func (tmp *Person) setPointer() {
	tmp.name = "jj"
	fmt.Printf("setPointer %p %v \n", tmp, tmp)
}
