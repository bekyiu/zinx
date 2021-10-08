package ziface

// 路由
type IRouter interface {
	PreHandle(req IRequest)
	Handle(req IRequest)
	PostHandle(req IRequest)
}