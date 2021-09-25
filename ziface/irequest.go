package ziface

// 封装客户端的请求和连接
type IRequest interface {
	// 获取当前连接
	GetConn() IConnection
	// 获取请求数据
	GetData() []byte
}