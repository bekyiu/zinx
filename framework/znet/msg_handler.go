package znet

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
	"zinx/framework/util"
	"zinx/framework/ziface"
)

type MsgHandler struct {
	// msgId对应的router
	RouterMap map[uint32]ziface.IRouter
	// 存放请求的队列
	TaskQueue []chan ziface.IRequest
	// 处理任务的goroutine数量
	WorkerPoolSize uint32
}

func NewMsgHandler() *MsgHandler {
	return &MsgHandler{
		RouterMap:      make(map[uint32]ziface.IRouter),
		WorkerPoolSize: util.GlobalConfig.WorkerPoolSize,
		TaskQueue:      make([]chan ziface.IRequest, util.GlobalConfig.WorkerPoolSize),
	}
}

// 启动一个工作池, 工作池是单例的
func (h *MsgHandler) StartWorkPool() {
	for i := 0; i < int(h.WorkerPoolSize); i++ {
		h.TaskQueue[i] = make(chan ziface.IRequest, util.GlobalConfig.TaskQueueSize)
		go h.startWorker(i, h.TaskQueue[i])
	}
}

func (h *MsgHandler) startWorker(workerId int, taskQueue chan ziface.IRequest) {
	fmt.Printf("worker[%d] start\n", workerId)
	for {
		select {
		case req := <-taskQueue:
			fmt.Println("current workerId: " + strconv.Itoa(workerId))
			h.DoHandler(req)
		}
	}
}

func (h *MsgHandler) AddTask(req ziface.IRequest) {
	rand.Seed(time.Now().Unix())
	index := rand.Int() % len(h.TaskQueue)
	h.TaskQueue[index] <- req
}

func (h *MsgHandler) DoHandler(req ziface.IRequest) {
	id := req.GetMsg().GetId()
	router, ok := h.RouterMap[id]
	if !ok {
		panic("no router map to " + strconv.Itoa(int(id)))
	}
	router.PreHandle(req)
	router.Handle(req)
	router.PostHandle(req)

}

func (h *MsgHandler) AddRouter(msgId uint32, router ziface.IRouter) {
	if _, ok := h.RouterMap[msgId]; ok {
		panic("router map to " + strconv.Itoa(int(msgId)) + " is already exist")
	}
	h.RouterMap[msgId] = router
}
