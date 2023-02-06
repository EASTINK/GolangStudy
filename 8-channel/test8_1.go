package main

import (
	"fmt"
	"math/rand"
	"time"
)

//从,goroutine
func newTask(x int) {
	rand.Seed(time.Now().Unix())
	fmt.Printf("我是子: %d \n", x)
}

//主，goroutine
func main() {
	x := 0
	for {
		if x == 500 {
			break
		}
		x++
		fmt.Printf("我是主 x = %d\n", x)
		go newTask(x)
		//time.Sleep(time.Millisecond.Round(100))
	}
}
