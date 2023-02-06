package main

import (
	"fmt"
	"strings"
	"time"
)

func IschatGPT(msg string) string {
	chat := strings.Split(msg, "\n")
	res := ""
	if chat[0] == "#gpt" {
		for x := range chat {
			if x != 0 {
				res += chat[x]
			}
		}
	}
	return res
}

func timef() {
	//CreateTime 1675317745
	//Current 1675321578

	fmt.Println(time.Now().Unix()) //- 1675317745)
	fmt.Println(time.Now().Unix() - 1675317745)
}

func main() {
	// switch {
	// case 0:
	// 	fmt.Println('0')
	// 	break
	// case 1:
	// 	fmt.Println('2')
	// 	break
	// }
	//fmt.Println(IschatGPT("#gpt\n同学们早安，离寒假结束，新学期伊始还剩下：%d 秒.\n\n%s\n\n这是一条实验性的定时早报,手动查询可发送群指令 #Doomsday\n >退订回#T 祝您依然学习愉快."))
	timef()
}
