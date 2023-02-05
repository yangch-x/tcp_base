package znet

import "myTcpBase/ziface"

// BaseRouter 路由基础对象
type BaseRouter struct{}

func (b *BaseRouter) PreHandle(req ziface.IRequest) {}

func (b *BaseRouter) Handle(req ziface.IRequest) {}

func (b *BaseRouter) PostHandle(req ziface.IRequest) {}
