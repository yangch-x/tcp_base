package znet

import "myTcpBase/ziface"

type Request struct {
	coon ziface.IConnection // 已经建立好的连接
	msg  ziface.IMessage    //客户端请求的数据
}

func (r *Request) GetConnection() ziface.IConnection {
	return r.coon
}

func (r *Request) GetData() []byte {
	return r.msg.GetData()
}

// GetMsgID 获取请求的消息的ID
func (r *Request) GetMsgID() uint32 {
	return r.msg.GetMsgID()
}
