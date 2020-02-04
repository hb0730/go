package main

import (
	"encoding/json"
	"fmt"
)

/**
json
*/
func main() {
	//test1();
	//test2();
	//test3();
	//test4();
	test5()
}

/**
json
*/
func test1() {
	s1 := IT{Company: "gs", Subjects: []string{"go", "java", "js", "c++"}, IsOk: true, Price: 66.66}
	b, err := json.Marshal(s1)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))

	b, err = json.MarshalIndent(s1, "", " ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))
}

/**
json 小写 缺失等
*/
func test2() {
	s1 := IT2{Company: "gs", Subjects: []string{"go", "java", "js", "c++"}, IsOk: true, Price: 66.66}
	b, err := json.Marshal(s1)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))
}

/**
map to json
*/
func test3() {
	m1 := make(map[string]interface{}, 4)
	m1["company"] = "gs"
	m1["subjects"] = []string{"go", "java", "js", "c++"}
	m1["isOk"] = true
	m1["price"] = 66.66
	j, err := json.Marshal(m1)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(j))
}

/**
json to 结构体
*/
func test4() {
	var s string = `{"Company":"gs","IsOk":true,"Price":66.66,"Subjects":["go","java","js","c++"]}`
	var tmp IT
	err := json.Unmarshal([]byte(s), &tmp)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(tmp)
}

/**
json to map
*/
func test5() {
	var s string = `{"company":"gs","isOk":true,"price":66.66,"subjects":["go","java","js","c++"]}`
	var m map[string]interface{}
	err := json.Unmarshal([]byte(s), &m)
	if err != nil {
		fmt.Println("error:", err)
	}
	//var c interface{};
	//c = m["company"]
	for key, value := range m {
		// 当前value type 为interface
		fmt.Printf("map[%s]=%v %T\n", key, value, value)
		//断言类型
		switch data := value.(type) {
		default:
			m[key] = data
		}
	}
}

//json 成员首字母大写
type IT struct {
	Company  string
	Subjects []string
	IsOk     bool
	Price    float64
}

//json 成员首字母大写
type IT2 struct {
	Company  string   `json:"company"`     //二次编码
	Subjects []string `json:"-"`           //此字段不会输出
	IsOk     bool     `json:"isOk,string"` //string类型
	Price    float64  `json:"price"`
}
