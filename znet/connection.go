package znet

import (
	"errors"
	"fmt"
	"io"
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
	// 告知当前连接需要停止
	ExitChan chan bool
	// 对应的处理方法
	MsgHandler ziface.IMsgHandler
	// 用于读写分离
	MsgChan chan []byte
	// 当前连接属于哪个server
	Server ziface.IServer
}

func (c *Connection) startReader() {
	fmt.Println("reader start")
	defer fmt.Println("reader end")
	defer c.Stop()

	for {
		dp := NewDataPack()
		header := make([]byte, dp.GetHeaderLen())
		// 先读header
		if _, err := io.ReadFull(c.Conn, header); err != nil {
			fmt.Println(err)
			break
		}

		msg, err := dp.Unpack(header)
		if err != nil {
			panic(err)
		}
		// 根据读到的len读data

		if msg.GetDataLen() <= 0 {
			continue
		}

		data := make([]byte, msg.GetDataLen())
		if _, err := io.ReadFull(c.Conn, data); err != nil {
			panic(err)
		}

		msg.SetData(data)

		req := Request{
			Conn: c,
			Msg:  msg,
		}

		// 存入任务队列
		c.MsgHandler.AddTask(&req)
	}

}

func (c *Connection) startWriter() {
	fmt.Println("writer start")
	defer fmt.Println("writer end")

	for {
		select {
		case data := <-c.MsgChan:
			if _, err := c.Conn.Write(data); err != nil {
				panic(err)
			}
		case <-c.ExitChan:
			return
		}
	}

}

func (c *Connection) Start() {
	fmt.Printf("conn[%d] start\n", c.ConnId)
	// 读业务
	go c.startReader()
	go c.startWriter()
}

func (c *Connection) Stop() {
	if c.isClose == false {
		c.isClose = true
		c.Server.GetConnPool().Remove(c.ConnId)
		// close后可读
		close(c.ExitChan)
		close(c.MsgChan)
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

func (c *Connection) Send(msgId uint32, data []byte) error {
	if c.isClose {
		return errors.New("connection is already closed")
	}
	msg := NewMessage(msgId, data)
	dp := NewDataPack()
	bytes, err := dp.Pack(msg)
	if err != nil {
		panic(err)
	}
	c.MsgChan <- bytes
	return nil
}

func NewConnection(conn *net.TCPConn, connId uint32, server ziface.IServer) *Connection {
	return &Connection{
		Conn:       conn,
		ConnId:     connId,
		isClose:    false,
		MsgHandler: server.GetMsgHandler(),
		ExitChan:   make(chan bool, 1),
		MsgChan:    make(chan []byte),
		Server:     server,
	}
}
