package main

import "fmt"

func fibonacii(c, quit chan int) {
	x, y := 1, 1
	ws := 0
	for {
		select {
		case c <- x: // 如果 c 被写入数据
			x = y
			y = x + y
		case <-quit: //如果 quit 读到数据
			fmt.Printf("quit 经历了%d个default \n", ws)
			return
		default:
			ws += 1
		}

	}
}
func main() {
	c := make(chan int)
	quit := make(chan int)
	// 在channel作用下 主go程main和从匿名go是同步的
	// 所以相当于fibonacii在匿名go程的for下被执行 然后进select case
	go func() {
		for i := 0; i < 6; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	/*
		输出:
			1
			1
			2
			4
			8
			16
			quit 经历了1580271个default
	*/

	fibonacii(c, quit)
}
