package znet

import (
	"zinx/ziface"
)

type Request struct{
	Conn ziface.IConnection
	Msg ziface.IMessage
}


func (r *Request) GetConn() ziface.IConnection {
	return r.Conn
}

func (r *Request) GetMsg() ziface.IMessage {
	return r.Msg
}
