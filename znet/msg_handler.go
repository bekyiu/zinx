package znet

import (
	"strconv"
	"zinx/ziface"
)

type MsgHandler struct {
	RouterMap map[uint32]ziface.IRouter
}

func NewMsgHandler() *MsgHandler {
	return &MsgHandler{RouterMap: make(map[uint32]ziface.IRouter)}
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
