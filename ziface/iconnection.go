package ziface

import "net"

type IConnection interface {
	// 启动连接
	Start()
	// 停止连接
	Stop()
	// 获取当前连接绑定的socket
	GetTCPConn() *net.TCPConn
	// 获取当前连接模块的id
	GetConnId() uint32
	// 获取客户端状态
	GetRemoteAddr() net.Addr
	// 发数据
	Send(msgId uint32, data []byte) error
	//
	SetProperty(k string, v interface{})
	GetProperty(k string) (interface{}, error)
	RemoveProperty(k string)
}

// 处理业务的方法
type HandleFunc func(*net.TCPConn, []byte, int) error
