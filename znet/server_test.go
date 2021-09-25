package znet

import (
	"fmt"
	"net"
	"testing"
	"time"
	"zinx/ziface"
)

type PingRouter struct {
	BaseRouter
}

func (r *PingRouter) PreHandle(req ziface.IRequest) {
	fmt.Println("pre")
}
func (r *PingRouter) Handle(req ziface.IRequest) {
	fmt.Println("in")
	conn := req.GetConn().GetTCPConn()
	conn.Write(req.GetData())
}
func (r *PingRouter) PostHandle(req ziface.IRequest) {
	fmt.Println("post")
}

func TestServer(t *testing.T) {
	server := NewServer()
	server.AddRouter(&PingRouter{})
	server.Serve()
}

func TestClient(t *testing.T) {
	fmt.Println("client start")

	conn, _ := net.Dial("tcp", "127.0.0.1:9999")

	for {
		_, _ = conn.Write([]byte("hello nanase"))
		buf := make([]byte, 512)
		count, _ := conn.Read(buf)
		fmt.Println("server back: ", string(buf[:count]))
		time.Sleep(time.Second * 5)
	}
}
