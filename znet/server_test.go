package znet

import (
	"fmt"
	"io"
	"net"
	"testing"
	"time"
	"zinx/ziface"
)

type PingRouter struct {
	BaseRouter
}


func (r *PingRouter) Handle(req ziface.IRequest) {
	conn := req.GetConn()
	fmt.Printf("received: %s\n", string(req.GetMsg().GetData()))
	conn.Send(1, []byte("i love nanase"))
}


func TestServer(t *testing.T) {
	server := NewServer()
	server.AddRouter(0, &PingRouter{})
	server.Serve()
}

func TestClient(t *testing.T) {
	fmt.Println("client start")

	conn, _ := net.Dial("tcp", "127.0.0.1:9999")

	for {
		// 写
		dp := NewDataPack()
		msg := NewMessage(0, []byte("hello nanase"))
		data, _ := dp.Pack(msg)
		_, _ = conn.Write(data)

		// 读
		buf := make([]byte, dp.GetHeaderLen())
		io.ReadFull(conn, buf)
		header, _ := dp.Unpack(buf)
		buf = make([]byte, header.GetDataLen())
		io.ReadFull(conn, buf)
		fmt.Printf("received: %s\n", string(buf))

		time.Sleep(time.Second * 5)
	}
}

func TestPanic(t *testing.T) {
	ch := make(chan int)
	go func() {
		ch <- 1
	}()
	go func() {
		ch <- 2
	}()
	go func() {
		ch <- 3
	}()
	time.Sleep(time.Second)
	fmt.Println(<- ch)
	fmt.Println(<- ch)
	fmt.Println(<- ch)
}
