package main

import "fmt"

func main() {
	a := "abcdefghijklmn"
	b := a[0 : len(a)/4]
	fmt.Println(b)
}
