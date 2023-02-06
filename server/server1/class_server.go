package server1

import (
	"fmt"
	"net"
	"strconv"
)

type Server struct {
	IP   string
	PORT int
}

// 类创建器
func NewServer(ip string, port int) *Server {
	return &Server{
		IP:   ip,
		PORT: port,
	}
}

// 处理其他业务-来自Server
func (s *Server) Handler(conn net.Conn) {
	//客户端打印连接进来的日志
	conn.Write([]byte("欢迎来到EASTFLOW.ICU"))
	//服务端打印连接进来的日志
	//log.Fatal(conn.RemoteAddr().String())
	fmt.Println(conn.RemoteAddr().String(), "进入")
}

// 服务启动器-来自Server
func (s *Server) Start() {
	//需要net库 创建一个TCP监听 异常处理这个先不管不然代码比较难看
	L, _ := net.Listen("tcp", s.IP+":"+strconv.Itoa(s.PORT))
	//设置回收
	defer L.Close()
	for {
		//默认同意所有请求接入 返回远程实体
		conn, _ := L.Accept()
		//可能存在不止一个客户端 引入协程解决占用
		// go func() {
		// 	//客户端打印连接进来的日志
		// 	conn.Write([]byte("欢迎来到EASTFLOW.ICU"))
		// 	//服务端打印连接进来的日志
		// 	//log.Fatal(conn.RemoteAddr().String())
		// 	fmt.Println(conn.RemoteAddr().String(), "进入")
		// }(	)
		go s.Handler(conn)
	}
}
