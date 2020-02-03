package main

import "fmt"

/**
方法
*/
func main() {
	//qubie();
	jiegoutileixing()
}

/**
普通函数与方法的区别
*/
func qubie() {
	var result int
	result = add01(1, 1)
	fmt.Println("add01 - ", result)
	var a myInt = 1
	result = int(a.add02(1))
	fmt.Print("add02= ", result)
}

//面向过程
func add01(a, b int) int {
	return a + b
}

type myInt int

// 面向对象
func (tmp myInt) add02(t myInt) myInt {
	return tmp + t
}

/**
结构体类型
*/
func jiegoutileixing() {
	s1 := Student{id: 1, addr: "上海"}
	s1.printlnInfo()
	s2 := myStudent{id: 2, addr: "上海"}
	s2.printlnInfoType()

	var s4 Student
	(&s4).setInfo()
	fmt.Printf("s3 = %+v", s4)
}

type Student struct {
	id   int
	addr string
}

func (tmp Student) printlnInfo() {
	fmt.Printf("tmp = %+v \n", tmp)
}

type myStudent Student

func (tmp myStudent) printlnInfoType() {
	fmt.Printf("tmp = %+v \n", tmp)
}

/*赋值*/
func (tmp *Student) setInfo() {
	tmp.addr = "上海"
	tmp.id = 3
}
