package ziface

/*
IRequest 消息请求抽象类
我们现在需要把客户端请求的连接信息 和 请求的数据，放在一个叫Request的请求类里，
这样的好处是我们可以从Request里得到全部客户端的请求信息，也为我们之后拓展框架有一定的作用，
一旦客户端有额外的含义的数据信息，都可以放在这个Request里。
可以理解为每次客户端的全部请求数据，Zinx都会把它们一起放到一个Request结构体里。
*/

type IRequest interface {
	GetIConnect() IConnection //获取连接信息
	GetData() []byte
	GetMsgID() uint32
}
