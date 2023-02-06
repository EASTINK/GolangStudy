package main

//引入sever包
import "server/server2"

func main() {
	server := server2.NewServer("127.0.0.1", 8890)
	server.Start()
}
