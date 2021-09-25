package znet

import (
	// 路径 不是包名
	"fmt"
	"net"
	"zinx/ziface"
)

type Server struct {
	Name   string
	IPVer  string
	IPAddr string
	Port   int
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
			dealConn := NewConnection(conn, cid, func(conn *net.TCPConn, data []byte, len int) error {
				conn.Write(data[:len])
				return nil
			})
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

// 包名.xxx
func NewServer(name string) (s ziface.IServer) {
	s = &Server{
		Name:   name,
		IPVer:  "tcp4",
		IPAddr: "0.0.0.0",
		Port:   8999,
	}
	return
}
