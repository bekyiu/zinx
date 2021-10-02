package ziface

import "net"

// 连接池 管理连接
type IConnPool interface {
	Remove(id uint32)
	Get(id uint32) (IConnection, error)
	// 获取当前连接数
	Num() int
	// 清楚并终止所有连接
	Clear()
	// 从连接池分配一个连接
	Allocate(conn *net.TCPConn, server IServer) (IConnection, error)
}
