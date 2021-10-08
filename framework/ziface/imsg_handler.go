package ziface

type IMsgHandler interface {
	// 执行对应的handler
	DoHandler(req IRequest)
	// 添加msgId 和 router的映射
	AddRouter(msgId uint32, router IRouter)
	//
	StartWorkPool()
	// 添加一个任务到随机一个任务队列
	AddTask(req IRequest)
}
