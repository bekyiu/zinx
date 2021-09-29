package znet

import (
	// 路径 不是包名
	"fmt"
	"net"
	. "zinx/util"
	"zinx/ziface"
)

type Server struct {
	Name   string
	IPVer  string
	IPAddr string
	Port   int
	MsgHandler ziface.IMsgHandler
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
		var cid uint32 = 0
		for {
			conn, err := listener.AcceptTCP()
			fmt.Println("accept a connection: ", conn.RemoteAddr())
			if err != nil {
				panic(err)
			}
			// 构造连接 处理业务
			dealConn := NewConnection(conn, cid, s.MsgHandler)
			cid++
			dealConn.Start()
		}
	}()
}

func (s *Server) Stop() {

}

func (s *Server) Serve() {
	s.Start()

	// 阻塞
	select {}
}

func (s *Server) AddRouter(msgId uint32, router ziface.IRouter) {
	s.MsgHandler.AddRouter(msgId, router)
}

// 包名.xxx
func NewServer() (s ziface.IServer) {
	s = &Server{
		Name:   GlobalConfig.Name,
		IPVer:  "tcp4",
		IPAddr: GlobalConfig.Host,
		Port:   GlobalConfig.Port,
		MsgHandler: NewMsgHandler(),
	}
	fmt.Println(GlobalConfig)
	return
}
