package ziface

type IServer interface {
	// 启动服务
	Start()
	Stop()
	// 启动业务
	Serve()
}