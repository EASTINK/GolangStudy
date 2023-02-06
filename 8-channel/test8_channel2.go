package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//设置缓冲大小
	size := 3
	sendsize := 4
	//定义chanel
	c := make(chan int, size) //当无法存入时阻塞
	fmt.Println("chanel c: 默认缓冲空间=", cap(c))
	//从go程发送元素
	go func() {
		defer fmt.Println("------------从go程 发送完毕------------")
		for i := 0; i < sendsize; i++ {
			randint := rand.Int()
			// //发送随机数
			c <- randint
			//发送延迟
			time.Sleep(1 * time.Second)
			fmt.Printf("从go程 发送1个元素：%d  %d / %d \n", randint, i+1, sendsize)
		}
	}()

	//主go程接受从go程发送过来的c
	for i := 0; i < sendsize; i++ {
		//接收延迟
		time.Sleep(6 * time.Second)
		value := <-c
		fmt.Printf("| 主go程 收到元素：%d  进度: %d / %d \n", value, i+1, sendsize)
		fmt.Printf("| 从go程 剩余空间: %d \n", size-len(c))
	}

	defer fmt.Println("------------主go程 接收完毕------------")
}
