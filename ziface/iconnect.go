package ziface

import "net"

// IConnection 连接模块接口
type IConnection interface {
	// Start 启动连接
	Start()
	// Stop 停止连接
	Stop()
	// GetTcpConnection 获取当前模块句柄
	GetTcpConnection() *net.TCPConn
	// GetConnId 获取当前连接id
	GetConnId() uint32
	// RemoteAddr 获取客户端连接状态，ip 端口
	RemoteAddr() net.Addr
	// SendMsg 发送
	SendMsg(id uint32, data []byte) error
}

// HandleFunc 定义处理业务的函数
// *net.TCPConn 连接句柄 []byte 处理数据内容 int 处理数据长度
type HandleFunc func(*net.TCPConn, []byte, int) error
