package ziface

import "net"

/*属于链接的接口文件*/

type IConnection interface {
	// Start 启动连接，让当前连接开始工作
	Start()
	// Stop 停止连接，结束当前连接状态M
	Stop()
	// GetTCPConnection 从当前连接获取原始的socket TCPConn
	GetTCPConnection() *net.TCPConn
	// GetConnID 获取当前连接ID
	GetConnID() uint32
	// RemoteAddr 获取远程客户端地址信息
	RemoteAddr() net.Addr
	// SendMsg 加密并发送信息
	SendMsg(message IMessage) error
	// Handle 处理数据
	Handle(req IRequest)
}

/*
该接口的一些基础方法，代码注释已经介绍的很清楚，
这里先简单说明一个HandFunc这个函数类型，这个是所有conn链接在处理业务的函数接口，
第一参数是socket原生链接，第二个参数是客户端请求的数据，第三个参数是客户端请求的数据长度。
这样，如果我们想要指定一个conn的处理业务，只要定义一个HandFunc类型的函数，然后和该链接绑定就可以了。
*/
