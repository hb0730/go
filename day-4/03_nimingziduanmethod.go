package main

import "fmt"

/**
匿名字段方法
*/
func main() {
	var s1 Student
	s1.id = 1
	s1.name = "张三"
	s1.printStudentInfo()
	s1.printPersonInfo()
	fmt.Println()
	s1.printInfo()
	s1.Person.printInfo()
}

type Person struct {
	name string
	sex  byte
	age  int
}

/**
方法
*/
func (tmp *Person) printPersonInfo() {
	fmt.Printf("name = %s ,sex = %c ,age = %d", tmp.name, tmp.sex, tmp.age)
}

func (tmp *Person) printInfo() {
	fmt.Println("person info 方法")
}

type Student struct {
	Person
	id   int
	addr string
}

func (tmp *Student) printStudentInfo() {
	tmp.printPersonInfo()
	fmt.Printf(" , id = %d,addr = %s \n", tmp.id, tmp.addr)
}

func (tmp *Student) printInfo() {
	fmt.Println("student info 方法重命名")
}
