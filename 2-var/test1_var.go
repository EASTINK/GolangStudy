package main

import "fmt"

/*
四种变量声明方式
*/

// 全局测试：
var gA int
var gB int = 100
var gC = 200

//:= 只能够在函数体内声明
//gD:=400;

// 局部测试
func main() {
	gA = 0
	//方法一 声明一个变量 默认的值是0
	var a int
	fmt.Println("a= ", a)
	fmt.Printf("type of a = %T\n", a)
	//方法二 声明一个变量同时初始化一个值
	var b int = 100
	fmt.Println("b = ", b)
	fmt.Printf("type of b = %T\n", b)
	//方法三 初始化时数据类型，通过赋值自动匹配变量的数据类型
	var c = "abcd"
	fmt.Println("c = ", c)
	fmt.Printf("type of c = %T\n", c)
	//方法四 (常用方法) 省去var 直接匹配
	e := 100
	fmt.Println("e =", e)
	fmt.Printf("type of e = %T \n", e)

	f := "abcd"
	fmt.Println("f =", f)
	fmt.Printf("type of f = %T \n", f)

	g := 3.14
	fmt.Println("g =", g)
	fmt.Printf("type of g = %T \n", g)

	//测试全局
	fmt.Println("gA =", gA, "gB =", gB, "gC =", gC) //,"gD = ",gD)
	//声明多个变量
	var xx, yy int = 100, 200
	fmt.Println("xx = ", xx, "yy = ", yy)
	var kk, ll = 100, "Aceld"
	fmt.Println("kk = ", kk, "ll = ", ll)
	//多行声明
	var (
		vv int  = 100
		jj bool = true
	)
	fmt.Println("vv = ", vv, "jj = ", jj)
}
