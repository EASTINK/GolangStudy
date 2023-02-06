package main

import (
	"fmt"
)

// 用const 来定义枚举类型
const (
	BEIJING  = 0
	SHANGHAI = 1
	SHENZHEN = 2
)

// 可以在const() 添加一个关键字iota 让每行iota会累加1 第一行的默认值为0
const (
	BEIJINGA  = iota
	SHANGHAIA //-> 1
	SHENZHENA //-> 2
)

// 多应用1
const (
	BEIJINGB  = 10 * iota //-> 0 * 10
	SHANGHAIB             //-> 1 * 10
	SHENZHENB             //-> 2 * 10
)

// 多应用2
const (
	a, b = iota + 1, iota + 2 //iota=0,公式：iota + 1, iota + 2, a = 1, b = 2
	c, d                      //iota=1,公式：iota + 1, iota + 2, c = 2, d = 3
	e, f                      //iota=2,公式：iota + 1, iota + 2, e = 3, f = 4
	g, h = iota * 2, iota * 3 //iota=3,公式：iota * 2, iota * 3, g = 6, h = 9
	i, k                      //iota=4,公式：iota * 2, iota * 3, i = 8, k = 12
)

func main() {
	//常量（只读）
	const length int = 10
	// length = 20  -> cannot assign to length (constant 10 of type int)
	//fmt.Println("length  = ", length)
	fmt.Printf("T =  %T\n", length)
	//测试iota常量
	fmt.Println("BEIJING = ", BEIJINGA)   //-> 0
	fmt.Println("SHANGHAI = ", SHANGHAIA) //-> 1
	fmt.Println("SHENZHEN = ", SHENZHENA) //-> 2
	//测试iota常量2
	fmt.Println("BEIJING = ", BEIJINGB)   //-> 0
	fmt.Println("SHANGHAI = ", SHANGHAIB) //-> 10
	fmt.Println("SHENZHEN = ", SHENZHENB) //-> 20
	//测试iota常量3
	fmt.Println(
		"a = ", a, " b = ", b,
		"\nc = ", c, " d = ", d,
		"\ne = ", e, " f = ", f,
		"\ng = ", g, " h = ", h,
		"\ni = ", i, " k = ", k)

	//iota 只能在const中使用 无法赋值给变量
	//var a int = iota
}
