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

		for {
			conn, err := listener.AcceptTCP()
			fmt.Println("accept a connection: ", conn.RemoteAddr())
			if err != nil {
				panic(err)
			}
			go func() {
				for {
					buf := make([]byte, 512)
					count, _ := conn.Read(buf)
					if err != nil {
						panic(err)
					}
					fmt.Println("accept value: ", string(buf))
					_, _ = conn.Write(buf[:count])
					if err != nil {
						panic(err)
					}
				}
			}()
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
