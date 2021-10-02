package ziface

type IServer interface {
	// 启动服务
	Start()
	// 停止服务
	Stop()
	// 启动业务
	Serve()
	// 给当前服务注册一个路由
	AddRouter(msgId uint32, router IRouter)
	//
	GetMsgHandler() IMsgHandler
	//
	GetConnPool() IConnPool
}
