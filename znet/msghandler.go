package znet

import (
	"fmt"
	"github.com/0RAJA/zinx/global"
	"github.com/0RAJA/zinx/ziface"
	"log"
	"math/rand"
)

type MsgHandle struct {
	APIs           map[uint32]ziface.IRouter
	WorkerPoolSize uint32
	TaskQueue      []chan ziface.IRequest
}

func (m *MsgHandle) StartOneWorker(queue chan ziface.IRequest) {
	for {
		select {
		case req := <-queue:
			m.HandleRequest(req)
		}
	}
}

func (m *MsgHandle) StartWorkerPool() {
	for i := 0; i < int(m.WorkerPoolSize); i++ {
		queue := make(chan ziface.IRequest, global.ServerSetting.MaxTaskQueueSize)
		m.TaskQueue[i] = queue
		go m.StartOneWorker(queue)
	}
}

func (m *MsgHandle) SendMsgToTaskQueue(request ziface.IRequest) {
	//随机分配处理队列
	workerID := rand.Intn(int(m.WorkerPoolSize))
	fmt.Println("Add ConnID=", request.GetIConnect().GetConnID(), " request msgID=", request.GetMsgID(), " to WorkerID=", workerID)
	m.TaskQueue[workerID] <- request
}

func NewMsgHandler() *MsgHandle {
	return &MsgHandle{APIs: make(map[uint32]ziface.IRouter), WorkerPoolSize: global.ServerSetting.WorkerPoolSize, TaskQueue: make([]chan ziface.IRequest, global.ServerSetting.WorkerPoolSize)}
}

func (m *MsgHandle) HandleRequest(request ziface.IRequest) {
	if handle, ok := m.APIs[request.GetMsgID()]; ok {
		handle.PreHandle(request)
		handle.Handle(request)
		handle.PostHandle(request)
	} else {
		log.Println("API msgID = ", request.GetMsgID(), " is not FOUND!")
	}
}

func (m *MsgHandle) AddRouter(msgID uint32, router ziface.IRouter) {
	if _, ok := m.APIs[msgID]; ok {
		panic(fmt.Sprint("API msgID = ", msgID, " is not FOUND!"))
	}
	m.APIs[msgID] = router
}
