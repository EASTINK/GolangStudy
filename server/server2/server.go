package server2

import (
	"net"
	"strconv"
	"sync"
)

/*
server类
IP  启动地址
PORT 目的端口
Onlie Msg  在线用户表
Msgcast 广播器
*/
type Server struct {
	IP      string
	PORT    int
	Onlie   map[string]*User //string->user
	mapLock sync.RWMutex     //读写锁 线程安全
	Msgcast chan string
}

// 创建器-来自Server
func NewServer(ip string, port int) *Server {
	return &Server{
		IP:      ip,
		PORT:    port,
		Onlie:   make(map[string]*User),
		Msgcast: make(chan string),
	}
}

// 业务-来自Server
func (s *Server) Handler(conn net.Conn) {
	//当用户上线 把它加入在线表
	user := &User{}
	user = user.NewUser(conn) //创建用户
	s.mapLock.Lock()          //当前线程加锁
	s.Onlie[user.Name] = user //加入表
	s.mapLock.Unlock()        //当前线程解锁
	//向所有用户广播当前用户上线
	s.Msgcast <- "[" + user.Addr + "]" + user.Name + ":" + "我上线了" //向广播器发送该用户上线的消息 -> MsgCast_listener -> User[]
	//selct阻塞
	select {} //保持阻塞状态等待唤醒
}

// 广播器-监听发送进来的消息
func (s *Server) Msg_ListenMeg() {
	for {
		msg := <-s.Msgcast // 接收传入广播器的信息
		s.mapLock.Lock()   //当前线程加锁
		for _, user := range s.Onlie {
			user.Context <- msg //向用户转发传入的消息
		}
		s.mapLock.Unlock() //当前线程解锁
	}
}

// 启动器-来自Server
func (s *Server) Start() {
	//需要net库 创建一个TCP监听 异常处理先不管 nil的代码比较难看...
	L, _ := net.Listen("tcp", s.IP+":"+strconv.Itoa(s.PORT))
	//设置回收
	defer L.Close()
	//加载广播监听器
	go s.Msg_ListenMeg()
	for {
		//默认同意所有请求接入 返回远程实体
		conn, _ := L.Accept()
		//加载该用户的业务协程
		go s.Handler(conn)
	}
}
