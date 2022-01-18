package znet

import "github.com/0RAJA/zinx/ziface"

type Request struct {
	conn ziface.IConnection //建立的连接
	msg  ziface.IMessage    //客户端请求的数据
}

func NewRequest(conn ziface.IConnection, msg ziface.IMessage) *Request {
	return &Request{conn: conn, msg: msg}
}

func (r *Request) GetIConnect() ziface.IConnection {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.msg.GetData()
}

func (r *Request) GetMsgID() uint32 {
	return r.msg.GetMsgID()
}
