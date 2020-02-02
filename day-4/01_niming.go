package main

import "fmt"

func main() {
	initData()
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
