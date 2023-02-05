package ziface

// IServer server接口定义
type IServer interface {
	Start() // 开启服务
	Stop()  // 停止服务
	Server()
	AddRouter(msgId uint32, router IRouter) IServer // 给当前服务提供添加router方法
}
