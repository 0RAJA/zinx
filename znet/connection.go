package znet

import (
	"errors"
	"fmt"
	"github.com/0RAJA/zinx/global"
	"github.com/0RAJA/zinx/ziface"
	"io"
	"log"
	"net"
)

type Connection struct {
	//当前连接的socket TCP套接字
	Conn *net.TCPConn
	//当前连接的ID 也可以称作为SessionID，ID全局唯一
	ConnID uint32
	//当前连接的关闭状态
	isClosed bool
	//告知该链接已经退出/停止的channel
	ExitBuffChan chan bool
	//消息处理
	msgHandle ziface.IMsgHandler
	//消息传顺chan 用于读写协程间通信
	msgChan chan []byte
}

func NewConnection(conn *net.TCPConn, connID uint32, msgHandle ziface.IMsgHandler) ziface.IConnection {
	return &Connection{Conn: conn, ConnID: connID, msgHandle: msgHandle, ExitBuffChan: make(chan bool, 1), msgChan: make(chan []byte)}
}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

// Start 启动连接，让当前连接开始工作
func (c *Connection) Start() {
	go c.StartWrite()
	go c.StartReader()
	<-c.ExitBuffChan //得到退出通知,不再阻塞
}

// Stop 停止连接，结束当前连接状态M
func (c *Connection) Stop() {
	//当前连接已经关闭
	if c.isClosed {
		return
	}
	c.isClosed = true
	//TODO Connection Stop() 如果用户注册了该链接的关闭回调业务，那么在此刻应该显示调用
	//关闭Socket连接
	c.Conn.Close()
	//通知从缓冲队列读数据的业务,此连接已经关闭
	//c.ExitBuffChan <- true
	//关闭这个链接的所有chan
	close(c.ExitBuffChan)
}

func (c *Connection) StartWrite() {
	fmt.Println("[Write goroutines is running]")
	defer fmt.Println("[Write goroutines is done]")
	for {
		select {
		case data := <-c.msgChan:
			if _, err := c.Conn.Write(data); err != nil {
				log.Println("write err:", err)
				return
			}
		case <-c.ExitBuffChan:
			return
		}
	}
}

func (c *Connection) SendMsg(msg ziface.IMessage) error {
	if c.isClosed {
		return errors.New("connection closed")
	}
	dp := NewDataPack()
	data, err := dp.Pack(msg)
	if err != nil {
		return err
	}
	c.msgChan <- data
	return nil
}

func (c *Connection) Handle(req ziface.IRequest) {
	c.msgHandle.HandleRequest(req)
}

// StartReader 处理conn读数据的G
func (c *Connection) StartReader() {
	fmt.Println("Reader G is running")
	defer fmt.Println(c.RemoteAddr().String(), "conn reader exit")
	defer c.Stop()
	dp := NewDataPack()
	for {
		headData := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(c.Conn, headData); err != nil {
			log.Println("read head err:", err)
			break
		}
		msg, err := dp.UnPack(headData)
		if err != nil {
			log.Println("unPack err:", err)
			break
		}
		if global.ServerSetting.MaxPacketSize > 0 && msg.GetDataLen() > global.ServerSetting.MaxPacketSize {
			log.Println("the message is too long")
			break
		}
		//判断是否有数据
		if msg.GetDataLen() > 0 {
			data := make([]byte, msg.GetDataLen())
			msg.SetData(data)
			//根据dataLen从io中读取字节流
			if _, err := io.ReadFull(c.Conn, data); err != nil {
				log.Println("upPack data err:", err)
				return
			}
			// 处理消息
			// 如果开启工作池和消息队列模式,则将其送至队列中,否则正常启动一个新协程来处理消息
			if global.ServerSetting.WorkerPoolSize > 0 {
				c.msgHandle.SendMsgToTaskQueue(NewRequest(c, msg))
			} else {
				go c.Handle(NewRequest(c, msg))
			}
		}
	}
}
