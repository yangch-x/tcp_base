package ziface

/*
	tlv 协议   Tag(Type)—Length—Value
	目前只有 id | len | data
	封包数据和拆包数据
	直接面向TCP连接中的数据流,为传输数据添加头部信息，用于处理TCP粘包问题。
*/
// IDataPack 数据包接口
type IDataPack interface {
	GetHeadLen() uint32                //获取包头长度方法
	Pack(msg IMessage) ([]byte, error) //封包方法
	Unpack([]byte) (IMessage, error)   //拆包方法
}

const (
	//DefaultPackDataPack 标准封包和拆包方式
	DefaultPackDataPack string = "default_pack"
)
