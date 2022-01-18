package ziface

import "net"

/*属于链接的接口文件*/

type IConnection interface {
	Start()                                 // Start 启动连接，让当前连接开始工作
	Stop()                                  // Stop 停止连接，结束当前连接状态M
	GetTCPConnection() *net.TCPConn         // GetTCPConnection 从当前连接获取原始的socket TCPConn
	GetConnID() uint32                      // GetConnID 获取当前连接ID
	RemoteAddr() net.Addr                   // RemoteAddr 获取远程客户端地址信息
	SendMsg(message IMessage) error         // SendMsg 加密并发送信息(无缓冲)
	SendMsgWithBuff(message IMessage) error //直接发送消息(有缓冲)
}

/*
该接口的一些基础方法，代码注释已经介绍的很清楚，
这里先简单说明一个HandFunc这个函数类型，这个是所有conn链接在处理业务的函数接口，
第一参数是socket原生链接，第二个参数是客户端请求的数据，第三个参数是客户端请求的数据长度。
这样，如果我们想要指定一个conn的处理业务，只要定义一个HandFunc类型的函数，然后和该链接绑定就可以了。
*/
