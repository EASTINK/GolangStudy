package main

import "fmt"

/*
	定义一个简单的类实现
*/
type ClassA struct {
	name string
	id   string
}

func (this ClassA) GetName() string {
	return this.name
}

func (this ClassA) GetID() string {
	return this.id
}

func (this *ClassA) SetID(id string) {
	this.id = id
}

func (this *ClassA) SetName(name string) {
	this.name = name
}

/*
	定义一个可用的接口
*/

type AIF interface {
	GetName() string
	GetID() string
	// SetName(name string)
	// SetID(id string)
}

func main() {
	//内部结构： pair<type:ClassA,value:{ss,1}
	var b AIF = &ClassA{"ss", "1"}
	fmt.Println(b.GetName())
}
