package server2

import "net"

/*
创建用户类:
Name名字、
Addr地址、
Conn 连接实体、
Context channel上下文
*/
type User struct {
	Name    string
	Addr    string
	Conn    net.Conn
	Context chan string
}

/*
创建器-来自User类

	NewUser(conn) *User
*/
func (s *User) NewUser(conn net.Conn) *User {
	userAddr := conn.RemoteAddr().String() //获得连接地址

	user := &User{
		Name:    userAddr,
		Addr:    userAddr,
		Conn:    conn,
		Context: make(chan string), //初始化上下文
	}

	go user.ListenMsg() //监听外部向内部发送的信息
	return user
}

/*
监听消息 -来自User类
*/
func (s *User) ListenMsg() {
	for {
		msg := <-s.Context                         // 从外部接收数据赋值给msg
		s.Conn.Write([]byte("收到消息：" + msg + "\n")) //在客户端显示外部数据
	}
}
