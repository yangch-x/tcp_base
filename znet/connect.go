package znet

import (
	"context"
	"fmt"
	"io"
	"myTcpBase/utils"
	"myTcpBase/ziface"
	"myTcpBase/zpack"
	"net"
)

type Connect struct {
	Conn       *net.TCPConn       // 连接句柄
	ConnId     uint32             // 连接id
	isClosed   bool               // 连接状态 true表示已经关闭
	ExitChan   chan bool          // 退出信号
	MsgHandler ziface.IMsgHandle  //当前Server的消息管理模块，用来绑定MsgID和对应的处理方法
	ctx        context.Context    //告知该链接已经退出/停止的channel
	cancel     context.CancelFunc //告知该链接已经退出/停止的cancel函数
	msgChan    chan []byte        // 消息读写chan
}

func NewConnection(conn *net.TCPConn, connId uint32, handle ziface.IMsgHandle) ziface.IConnection {
	return &Connect{
		Conn:       conn,
		ConnId:     connId,
		isClosed:   false,
		ExitChan:   make(chan bool, 1),
		msgChan:    make(chan []byte, utils.GlobalObject.MaxMsgChanLen),
		MsgHandler: handle,
	}
}

// StartRead 读函数
func (c *Connect) StartRead() {
	fmt.Println("[Reader Goroutine is running]")
	defer c.Stop()
	defer fmt.Println(c.RemoteAddr().String(), "[conn Reader exit!]")
	for {
		select {
		case <-c.ctx.Done():
			return
		default:
			p := zpack.NewDataPack()

			// 消息包头及消息体长度
			headData := make([]byte, p.GetHeadLen())
			_, err := io.ReadFull(c.Conn, headData)
			if err != nil {
				fmt.Println("read msg head error ", err)
				return
			}
			fmt.Printf("read headData %+v\n", headData)

			// 拆包，拆解数据包中len及data长度
			msg, err := p.Unpack(headData)
			if err != nil {
				fmt.Println("unpack error ", err)
				return
			}
			// 读取数据包data
			var data []byte
			if msg.GetDataLen() > 0 {
				data = make([]byte, msg.GetDataLen())
				_, err = io.ReadFull(c.Conn, data)
				if err != nil {
					fmt.Println("read msg head error ", err)
					return
				}
			}
			msg.SetData(data)
			req := Request{
				coon: c,
				msg:  msg,
			}
			go c.MsgHandler.DoMsgHandler(&req)
		}
	}
}

// StartWrite 写函数
func (c *Connect) StartWrite() {
	defer c.Stop()
	fmt.Println("[Writer Goroutine is running]")
	defer fmt.Println(c.RemoteAddr().String(), "[conn Writer exit!]")
	for {
		select {
		case data, ok := <-c.msgChan:
			if ok {
				//有数据要写给客户端
				if _, err := c.Conn.Write(data); err != nil {
					fmt.Println("Send Buff Data error:, ", err, " Conn Writer exit")
					return
				}
			} else {
				fmt.Println("msgBuffChan is Closed")
				break
			}
		case <-c.ctx.Done():
			return
		}
	}
}

func (c *Connect) Start() {
	c.ctx, c.cancel = context.WithCancel(context.Background())
	//按照用户传递进来的创建连接时需要处理的业务，执行钩子方法
	//	c.s.CallOnConnStart(c)
	//1 开启用户从客户端读取数据流程的Goroutine
	go c.StartRead()
	//2 开启用于写回客户端数据流程的Goroutine
	go c.StartWrite()

	select {
	case <-c.ctx.Done():
		//c.finalizer()
		return
	}
}

func (c *Connect) Stop() {
	if c.isClosed {
		return
	}
	c.isClosed = true
	err := c.Conn.Close()
	if err != nil {
		return
	}
	close(c.ExitChan)
}

func (c *Connect) GetTcpConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connect) GetConnId() uint32 {
	return c.ConnId
}

func (c *Connect) RemoteAddr() net.Addr {
	return c.RemoteAddr()
}

func (c *Connect) SendMsg(id uint32, data []byte) error {
	if c.isClosed {
		fmt.Println("connect id closed ")
		return nil
	}
	p := zpack.NewDataPack()
	msg := zpack.NewMsgPackage(id, data)
	resMsg, err := p.Pack(msg)
	if err != nil {
		fmt.Printf("msg pacl err:%v", err)
		return err
	}
	if _, err = c.Conn.Write(resMsg); err != nil {
		fmt.Printf("write info to client err%s", err)
		return err
	}
	return nil
}
