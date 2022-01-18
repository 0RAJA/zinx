package ziface

/*消息管理模块*/

type IMsgHandler interface {
	HandleRequest(request IRequest)
	AddRouter(msgID uint32, router IRouter)
	StartWorkerPool()
	SendMsgToTaskQueue(request IRequest)
}
