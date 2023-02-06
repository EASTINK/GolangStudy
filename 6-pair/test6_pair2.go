package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)

	if err != nil {
		fmt.Println("open file error", err)
		return
	}
	// r pair <type : vlaue:>
	var r io.Reader
	// r pair <type : os.FIle ,vlaue: /dev/tty>
	r = tty
	var w io.Writer
	// pair <type : os.FIle .value:/dev/tty>
	w = r.(io.Writer)

	w.Write([]byte("HELLO THIS is A TEST\n"))

}
