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
	// hook
	// 创建连接后
	SetAfterConnStart(hook func(connection IConnection))
	// 停止连接之前
	SetBeforeConnStop(hook func(connection IConnection))
	CallAfterConnStart(conn IConnection)
	CallBeforeConnStop(conn IConnection)
}
