package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (this *User) Call() {
	fmt.Println("user is called ..")
	fmt.Printf("%v\n", this)
}

func main() {
	user := User{1, "Done", 18}
	DoFiledAndMethod(user)
}

func DoFiledAndMethod(input interface{}) {
	//获取input 的type
	inputType := reflect.TypeOf(input)
	//获取input的value
	inputValue := reflect.ValueOf(input)
	//fmt.Println("inputValue：", inputType, "\ninputValue:", inputValue)
	//获取type 获取里面的字段
	//1 获取interface的reflect.Type 通过Type得到NumField
	//2 得到每个field 数据类型
	//3 通过filed 有一个Interface方法得到对应的value
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i).Interface()
		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}

	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}
}
