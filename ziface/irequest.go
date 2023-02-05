package ziface

// IRequest  封装请求数据
type IRequest interface {
	GetConnection() IConnection // 获取当前连接
	GetData() []byte            // 获取请求数据
	GetMsgID() uint32           //获取请求的消息ID
}
