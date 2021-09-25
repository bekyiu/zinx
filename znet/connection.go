package znet

import (
	"fmt"
	"net"
	"zinx/ziface"
)

type Connection struct {
	// 对应的socket
	Conn *net.TCPConn
	// id
	ConnId uint32
	// 是否关闭
	isClose bool
	// 对应的处理方法
	handleAPI ziface.HandleFunc
	// 告知当前连接需要停止
	ExitChan chan bool
}

func (c *Connection) startReader() {
	fmt.Println("reader start")
	defer fmt.Println("reader end")
	defer c.Stop()

	for {
		buf := make([]byte, 512)
		count, err := c.Conn.Read(buf)
		if err != nil {
			panic(err)
		}

		// 调用业务方法
		if err := c.handleAPI(c.Conn, buf, count); err != nil {
			panic(err)
		}
	}

}

func (c *Connection) Start() {
	fmt.Printf("conn[%d] start\n", c.ConnId)
	// 读业务
	go c.startReader()
}

func (c *Connection) Stop() {
	fmt.Printf("conn[%d] stop\n", c.ConnId)
	if c.isClose == false {
		c.isClose = true
		c.Conn.Close()
		close(c.ExitChan)
	}
}

func (c *Connection) GetTCPConn() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnId() uint32 {
	return c.ConnId
}

func (c *Connection) GetRemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (Connection) Send(data []byte) error {
	panic("implement me")
}


func NewConnection(conn *net.TCPConn, connId uint32, handleAPI ziface.HandleFunc) *Connection {
	return &Connection{
		Conn:      conn,
		ConnId:    connId,
		isClose:   false,
		handleAPI: handleAPI,
		ExitChan:  make(chan bool, 1),
	}
}
