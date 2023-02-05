package zpack

// Message 消息
type Message struct {
	DataLen uint32 //消息的长度 4byte
	ID      uint32 //消息的ID  4byte
	Data    []byte //消息的内容
}

// NewMsgPackage 创建一个Message消息包
func NewMsgPackage(ID uint32, data []byte) *Message {
	return &Message{
		DataLen: uint32(len(data)),
		ID:      ID,
		Data:    data,
	}
}

func (msg *Message) Init(ID uint32, data []byte) {
	msg.ID = ID
	msg.Data = data
	msg.DataLen = uint32(len(data))
}

// GetDataLen 获取消息数据段长度
func (msg *Message) GetDataLen() uint32 {
	return msg.DataLen
}

// GetMsgID 获取消息ID
func (msg *Message) GetMsgID() uint32 {
	return msg.ID
}

// GetData 获取消息内容
func (msg *Message) GetData() []byte {
	return msg.Data
}

// SetDataLen 设置消息数据段长度
func (msg *Message) SetDataLen(len uint32) {
	msg.DataLen = len
}

// SetMsgID 设计消息ID
func (msg *Message) SetMsgID(msgID uint32) {
	msg.ID = msgID
}

// SetData 设计消息内容
func (msg *Message) SetData(data []byte) {
	msg.Data = data
}
