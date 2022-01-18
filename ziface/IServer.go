package ziface

//服务器接口

type IServer interface {
	Start()                                                   //启动服务器
	Stop()                                                    //关闭服务器
	Server()                                                  //开启业务服务方法
	AddRouter(msgID uint32, router IRouter)                   //路由功能：给当前服务注册一个路由业务方法，供客户端链接处理使用
	GetConnMgr() IConnManager                                 //获取连接管理器
	SetOnConnStart(hookFuncs ...func(connection IConnection)) //设置开始的钩子
	SetOnConnStop(hookFuncs ...func(connection IConnection))  //设置结束的钩子
	CallOnConnStart(connection IConnection)                   //调用连接OnConnStart Hook函数
	CallOnConnStop(connection IConnection)                    //调用连接OnConnStop Hook函数
}
