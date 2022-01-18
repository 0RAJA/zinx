package ziface

//服务器接口

type IServer interface {
	Start()                                 //启动服务器
	Stop()                                  //关闭服务器
	Server()                                //开启业务服务方法
	AddRouter(msgID uint32, router IRouter) //路由功能：给当前服务注册一个路由业务方法，供客户端链接处理使用
}
