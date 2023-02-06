package main

import "fmt"

//主go程
func main() {
	//channel 定义
	c := make(chan int)
	//从go程
	go func() {
		defer fmt.Println("goroutine exit().")
		fmt.Println("go routine runtime..")
		c <- 666
	}()
	fmt.Println(<-c)
	/*
		channel
			会发送阻塞信号使得从go程和主go程同步，
			若主go程不取得c,那么从go程就会一直在阻塞状态
			所以如果注释了这行，这里的go程匿名函数就会一直阻塞不被执行
	*/
	fmt.Println("main go routine exit().")
}
