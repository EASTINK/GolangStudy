package main

import "fmt"

func main() {

	c := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
			//close(c)
		}
		close(c)

		/*
			tips
				close(channel chan <- type)
				1、channel 不像文件一样需要经常关闭 只有当确实没有任何发送数据以后 / 你想显式的结束range循环之类 才去关闭channel
				2、关闭channel后，无法再向channel发送数据 （引发panic错误后导致接受立即返回零值
				3、关闭channel后 可以继续从channel 接收数据

			输出:
					0
					1
					2
					3
					4
					Finished..
		*/
	}()

	for {

		if data, ok := <-c; ok {
			fmt.Println(data)
		} else {
			break
		}
	}

	fmt.Println("Finished..")
}
