package main

import (
	"fmt"
	"myTcpBase/ziface"
	"myTcpBase/znet"
)

func main() {
	// 模拟服务器启动
	s := znet.NewServer("server1")
	s.AddRouter(1, &PingRouter{}).AddRouter(2, &AddRouters{}).AddRouter(3, &UpdateRouters{})
	s.Server()
}

// PingRouter 自定义router
type PingRouter struct {
	znet.BaseRouter
}

func (b *PingRouter) PreHandle(request ziface.IRequest) {
	fmt.Println("Call PingRouter Handle")
	fmt.Println("recv from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))
	err := request.GetConnection().SendMsg(request.GetMsgID(), []byte(fmt.Sprintf("【PingRouter】收到消息%s", request.GetData())))
	if err != nil {
		fmt.Println(err)
	}
}

// AddRouters 自定义router
type AddRouters struct {
	znet.BaseRouter
}

func (b *AddRouters) PreHandle(request ziface.IRequest) {
	fmt.Println("Call PingRouter Handle")
	fmt.Println("recv from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))
	err := request.GetConnection().SendMsg(request.GetMsgID(), []byte(fmt.Sprintf("【AddRouters】收到消息%s", request.GetData())))
	if err != nil {
		fmt.Println(err)
	}
}

// UpdateRouters 自定义router
type UpdateRouters struct {
	znet.BaseRouter
}

func (b *UpdateRouters) PreHandle(request ziface.IRequest) {
	fmt.Println("Call PingRouter Handle")
	fmt.Println("recv from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))
	err := request.GetConnection().SendMsg(request.GetMsgID(), []byte(fmt.Sprintf("【UpdateRouters】收到消息%s", request.GetData())))
	if err != nil {
		fmt.Println(err)
	}
}

// DeleteRouters 自定义router
type DeleteRouters struct {
	znet.BaseRouter
}

func (b *DeleteRouters) PreHandle(request ziface.IRequest) {
	fmt.Println("Call PingRouter Handle")
	fmt.Println("recv from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))
	err := request.GetConnection().SendMsg(request.GetMsgID(), []byte(fmt.Sprintf("【DeleteRouters】收到消息%s", request.GetData())))
	if err != nil {
		fmt.Println(err)
	}
}
