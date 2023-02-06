package main

import "fmt"

func main() {

	c := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			//close(c)
			c <- i
		}

		close(c) //关闭channel
	}()

	/*

		for {
			if data, ok := <-c; ok {
				fmt.Println(data)
			} else {
				break
			}
		}
	*/
	//可以使用range迭代数据
	for data := range c {
		fmt.Println(data)
	}

	fmt.Println("Finished..")
}
