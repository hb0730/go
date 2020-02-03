package main

import (
	"fmt"
)

/**
匿名字段
*/
func main() {
	//initData()
	//tongming();
	//feijiegoutinimingziduan();
	jiegoutizhizhen()
}

/**
初始化
*/
func initData() {
	var student Student
	student.Person.name = "张三"
	student.id = 1
	fmt.Println("student = ", student)
	var s1 Student = Student{Person: Person{name: "李四"}, id: 2}
	fmt.Printf("s1 = %+v \n", s1)
}

/**
同名字段
*/
func tongming() {
	//初始化
	var teacher Teacher
	teacher.name = "ggg"
	teacher.Person.name = "sss"
	teacher.id = 1
	teacher.age = 'g'
	fmt.Printf(" t1 = %+v \n", teacher)
}

/**
非结构体匿名字段
*/
func feijiegoutinimingziduan() {
	var s1 Student1
	s1.int = 1
	s1.myStr = "上海"
	s1.Person.name = "hh"
	fmt.Printf("s1 = %+v \n", s1)
}

/**
结构体指针
*/
func jiegoutizhizhen() {
	s1 := Student2{Person: &Person{name: "张三"}, id: 1}
	fmt.Printf("s1 = %+v \n", s1)

	var s2 Student2
	s2.Person = &Person{name: "李四"}
	fmt.Printf("s2 = %+v \n", s2)

	var s3 Student2
	s3.Person = new(Person)
	s3.name = "王五"
	fmt.Printf("s3 = %+v", s3)
}
func test1() {

}

//指针
type Student2 struct {
	*Person
	id   int
	addr string
}

//起别名
type myStr string

/**
非结构体匿名字段
*/
type Student1 struct {
	Person
	int
	myStr
}

/**
重名字段
*/
type Teacher struct {
	Person
	name string
	id   int
	addr string
}
type Person struct {
	name string
	sex  byte
	age  int
}
type Student struct {
	Person
	id   int
	addr string
}
