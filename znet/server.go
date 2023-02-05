package znet

import (
	"fmt"
	"myTcpBase/ziface"
	"net"
)

type Server struct {
	// Name 服务名
	Name string
	// IPVersion ip 版本
	IPVersion string
	// IP ip
	IP string
	// Port 端口
	Port int
	//当前Server的消息管理模块，用来绑定MsgID和对应的处理方法
	msgHandler ziface.IMsgHandle
}

func NewServer(name string) ziface.IServer {
	return &Server{
		Name:       name,
		IPVersion:  "tcp",
		IP:         "0.0.0.0",
		Port:       8888,
		msgHandler: NewMsgHandle(),
	}
}

func (s *Server) Start() {
	fmt.Printf("[Start] server name: %s,server ip:%s,server port:%d", s.Name, s.IP, s.Port)
	go func() {
		// 获取服务器 tcp addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Printf("resoulve addr err:%v", err)
			return
		}
		// 监听服务地址
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			panic(err)
		}
		//已经监听成功
		fmt.Println("start server ", s.Name, " success, now listenning...")
		// 监听客户端连接，处理具体业务
		var cid uint32
		for {
			// 监听tcp连接
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err ", err)
				continue
			}
			fmt.Println("Get conn remote addr = ", conn.RemoteAddr().String())
			// 处理业务
			cid++
			connection := NewConnection(conn, cid, s.msgHandler)
			go connection.Start()
		}
	}()
}

func (s *Server) Stop() {

}

func (s *Server) Server() {
	// 启动服务
	s.Start()
	// 阻塞main
	select {}
}

func (s *Server) AddRouter(msgId uint32, router ziface.IRouter) ziface.IServer {
	s.msgHandler.AddRouter(msgId, router)
	return s
}
