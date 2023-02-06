package lib1

import "fmt"

func Lib1Test() {
	fmt.Println("hi, i am lib1")
}

func init() {
	fmt.Println(("lib1.init()...."))
}
