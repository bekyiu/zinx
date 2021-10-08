package znet

import (
	"zinx/framework/ziface"
)

type BaseRouter struct{}

func (r *BaseRouter) PreHandle(req ziface.IRequest) {
}

func (r *BaseRouter) Handle(req ziface.IRequest) {
}

func (r *BaseRouter) PostHandle(req ziface.IRequest) {
}
