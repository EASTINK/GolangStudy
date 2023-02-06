package lib2

import "fmt"

func Lib2Test() {
	fmt.Println("hi, i am lib2")
}

func init() {
	fmt.Println(("lib2.init()...."))
}
