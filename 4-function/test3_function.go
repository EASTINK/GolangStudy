package main

import "fmt"

// 基本函数
func foo1(a string, b int) int {
	fmt.Println("---- foo1 ----")
	fmt.Println("a = ", a, "b = ", b)
	c := 100
	return c

}

// 多返回值[匿名]-基本函数
func foo2(a string, b int) (int, int) {
	fmt.Println("---- foo2 ----")
	fmt.Println("a = ", a, "b = ", b)
	return 666, 777
}

// 多返回值[形参名]-基本函数
func foo3(a string, b int) (r1 int, r2 int) {
	fmt.Println("---- foo3 ----")
	fmt.Println("a = ", a, "b = ", b)
	//此时r1 , r2 也是foo3的形参， 拥有默认值int 0， 作用域是foo3(){}
	fmt.Println("r1 = ", r1, "r2 = ", r2)
	//返回形参
	r1 = 1000
	r2 = 2000
	return
}

// 如果返回形参类型一致，可以省略声明如：
func foo4(a string, b int) (r1, r2 int) {
	fmt.Println("---- foo4 ----")
	fmt.Println("a = ", a, "b = ", b)
	return 3000, 4000
}

func main() {
	//foo1
	c := foo1("foo1", 444)
	fmt.Println("c = ", c)
	//foo2
	ret1, ret2 := foo2("foo2", 222)
	fmt.Println("ret1=", ret1, "ret2=", ret2)
	//foo3
	ret3, ret4 := foo3("foo3", 333)
	fmt.Println("ret3 = ", ret3, "ret4 = ", ret4)
	//foo4
	ret3, ret4 = foo4("foo4", 333)
	fmt.Println("ret3 = ", ret3, "ret4 = ", ret4)
}
