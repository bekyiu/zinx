package znet

import (
	"fmt"
	"net"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	server := NewServer("zinx v1.0")
	server.Serve()
}

func TestClient(t *testing.T) {
	fmt.Println("client start")

	conn, _ := net.Dial("tcp", "127.0.0.1:8999")

	for {
		_, _ = conn.Write([]byte("hello nanase"))
		buf := make([]byte, 512)
		count, _ := conn.Read(buf)
		fmt.Println("server back: ", string(buf[:count]))
		time.Sleep(time.Second * 5)
	}
}
