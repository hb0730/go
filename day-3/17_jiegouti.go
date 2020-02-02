package main

import "fmt"

/**
结构体
*/
func main() {
	//initDate();
	//chengyuan();
	//equals();
	paramsTest()
}

/**
初始化
*/
func initDate() {
	//全部初始化
	var s1 Student = Student{1, "张三", 'm', 18, "中国上海"}
	fmt.Println("s1 = ", s1)
	//部分初始化,未初始化默认为0
	var s2 Student = Student{name: "李四"}
	fmt.Println("s2 = ", s2)

	//指针初始化
	var p1 *Student = &Student{name: "王五"}
	fmt.Println("*p1 = ", *p1)
}

/**
结构体成员使用
*/
func chengyuan() {
	//定义成员变量
	var s1 Student
	s1.name = "铁憨憨"
	fmt.Println("s1 = ", s1)
	//指针
	//1.
	var s2 Student
	var p1 *Student
	p1 = &s2
	p1.name = "立完案"
	fmt.Println("p1 = ", *p1)

	//2
	p2 := new(Student)
	p2.name = "张哈哈"
	fmt.Println("p2 = ", p2)
}

/**
结构体比较
*/
func equals() {
	//
	var s1 Student = Student{name: "张三"}
	var s2 Student = Student{name: "张三"}
	var s3 Student = Student{name: "李四"}
	var s4 Student = Student{age: 18}
	if s1 == s2 {
		fmt.Println("s1 ==s2")
	} else {
		fmt.Println("s2 != s2")
	}
	if s1 == s3 {
		fmt.Println("s1 ==s3")
	} else {
		fmt.Println("s1!=s3")
	}
	if s3 == s4 {
		fmt.Println("s3 ==s4")
	} else {
		fmt.Println("s3!=s4")
	}
}

/**
结构体参数传递
*/
func paramsTest() {
	var s1 Student = Student{1, "张三", 'm', 18, "中国上海"}
	fmt.Println("paramsTest", s1)
	_params(s1)
	fmt.Println("_params", s1)
	params(&s1)
	fmt.Println("params *", s1)
}
func _params(s1 Student) {
	s1.name = "ssss"
}
func params(s1 *Student) {
	s1.name = "ddddd"
}

/**
定义结构体
*/
type Student struct {
	/**
	学号
	*/
	id int
	/**
	姓名
	*/
	name string
	/**
	性别
	*/
	sex byte
	/**
	年龄
	*/
	age int
	/**
	地址
	*/
	addr string
}
