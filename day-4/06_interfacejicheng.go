package main

import "fmt"

/**
接口继承
*/
func main() {
	//test1();
	//test2()
	test3()
}

/**
超类与子集
*/
func test1() {
	student := Student{name: "张三"}
	sudent(student, "学生狗")
}

/**
类型转换
*/
func test2() {
	var ipro personer
	var i Humaner
	student := Student{name: "张三"}
	ipro = student
	//ipro =i //error
	i = ipro //超类转换子类
	i.sayHi()
}

/**
类型判断
*/
func test3() {
	//interface{}为空接口
	i := make([]interface{}, 3)
	i[0] = "mike"
	i[1] = 1
	i[2] = 'd'
	for index, data := range i {
		fmt.Printf("类型 %T", data)
		if value, ok := data.(int); ok == true {
			fmt.Printf("x[%d],类型为int,值为%d \n", index, value)
		} else if value, ok := data.(string); ok == true {
			fmt.Printf("x[%d],类型为string,值为%s \n", index, value)
		} else if value, ok := data.(byte); ok == true {
			fmt.Printf("x[%d],类型为byte,值为%c \n", index, value)
		} else {
			fmt.Println("未匹配")
		}
		fmt.Println("data = ", data)
		switch value := data.(type) {
		case int:
			fmt.Printf("x[%d],类型为int,值为%d \n", index, value)
			break
		case string:
			fmt.Printf("x[%d],类型为string,值为%s \n", index, value)
			break
		case byte:
			fmt.Printf("x[%d],类型为byte,值为%c \n", index, value)
			break
		default:
			fmt.Println("未匹配")
		}
	}
}

type Student struct {
	id   int
	name string
}

func (tmp Student) sayHi() {
	fmt.Printf("Student %s sayHi \n", tmp.name)
}
func (tmp Student) sing(lrc string) {
	fmt.Printf("student %s, sing %s", tmp.name, lrc)

}

type Humaner interface {
	sayHi()
}

type personer interface {
	Humaner
	sing(lrc string)
}

func sudent(s personer, lrc string) {
	s.sayHi()
	s.sing(lrc)
}
