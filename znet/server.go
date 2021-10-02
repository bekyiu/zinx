package znet

import (
	// 路径 不是包名
	"fmt"
	"net"
	. "zinx/util"
	"zinx/ziface"
)

type Server struct {
	Name           string
	IPVer          string
	IPAddr         string
	Port           int
	MsgHandler     ziface.IMsgHandler
	ConnPool       ziface.IConnPool
	AfterConnStart func(connection ziface.IConnection)
	BeforeConnStop func(connection ziface.IConnection)
}

func (s *Server) Start() {
	fmt.Printf("start server listener at %s:%d\n", s.IPAddr, s.Port)

	go func() {
		// 构造tcp addr
		addr, err := net.ResolveTCPAddr(s.IPVer, fmt.Sprintf("%s:%d", s.IPAddr, s.Port))
		if err != nil {
			panic(err)
		}

		// 监听服务器地址
		listener, err := net.ListenTCP(s.IPVer, addr)
		if err != nil {
			panic(err)
		}

		fmt.Println("start success")
		for {
			conn, err := listener.AcceptTCP()
			fmt.Println("accept a connection: ", conn.RemoteAddr())
			if err != nil {
				panic(err)
			}
			// 构造连接 处理业务
			dealConn, ok := s.ConnPool.Allocate(conn, s)
			if ok != nil {
				continue
			}
			dealConn.Start()
		}
	}()
}

func (s *Server) Stop() {
	s.ConnPool.Clear()
}

func (s *Server) Serve() {
	s.MsgHandler.StartWorkPool()
	s.Start()
	// 阻塞
	select {}
}

func (s *Server) AddRouter(msgId uint32, router ziface.IRouter) {
	s.MsgHandler.AddRouter(msgId, router)
}

func (s *Server) GetMsgHandler() ziface.IMsgHandler {
	return s.MsgHandler
}

func (s *Server) GetConnPool() ziface.IConnPool {
	return s.ConnPool
}

func (s *Server) SetAfterConnStart(hook func(conn ziface.IConnection)) {
	s.AfterConnStart = hook
}

func (s *Server) SetBeforeConnStop(hook func(conn ziface.IConnection)) {
	s.BeforeConnStop = hook
}

func (s *Server) CallAfterConnStart(conn ziface.IConnection) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	if s.AfterConnStart != nil {
		s.AfterConnStart(conn)
	}
}
func (s *Server) CallBeforeConnStop(conn ziface.IConnection) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	if s.BeforeConnStop != nil {
		s.BeforeConnStop(conn)
	}
}

// 包名.xxx
func NewServer() (s ziface.IServer) {
	s = &Server{
		Name:       GlobalConfig.Name,
		IPVer:      "tcp4",
		IPAddr:     GlobalConfig.Host,
		Port:       GlobalConfig.Port,
		MsgHandler: NewMsgHandler(),
		ConnPool:   NewConnPool(),
	}
	fmt.Println(GlobalConfig)
	return
}
