package ziface

type IRouter interface {
	PreHandle(req IRequest)  // 处理业务之前的方法
	Handle(req IRequest)     // 处理业务方法
	PostHandle(req IRequest) // 处理业务之后的方法
}
