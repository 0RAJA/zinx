package main

import (
	"fmt"
	"github.com/0RAJA/zinx/ziface"
	"github.com/0RAJA/zinx/znet"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type PingRouter struct {
	znet.BaseRouter
}

func (p *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call PingRouter Handle")
	//先读取客户端的数据，再回写ping...ping...ping
	fmt.Println("recv from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))

	err := request.GetIConnect().SendMsgWithBuff(znet.NewMessage(0, []byte("ping...ping...ping")))
	if err != nil {
		log.Println(err)
	}
}

type HelloRouter struct {
	znet.BaseRouter
}

func (h *HelloRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call HelloZinxRouter Handle")
	//先读取客户端的数据，再回写ping...ping...ping
	fmt.Println("recv from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))

	err := request.GetIConnect().SendMsgWithBuff(znet.NewMessage(1, []byte("Hello Zinx Router V0.10")))
	if err != nil {
		log.Println(err)
	}
}

func OnStart(connection ziface.IConnection) {
	fmt.Println("DoConnectionLost is Called ... ")

	//=============设置两个链接属性，在连接创建之后===========
	fmt.Println("Set conn Name, Home done!")
	connection.SetProperty("Name", "raja")
	connection.SetProperty("Home", "test")
	//===================================================

	err := connection.SendMsg(znet.NewMessage(2, []byte("DoConnection BEGIN...")))
	if err != nil {
		log.Println(err)
	}
}

func OnStop(connection ziface.IConnection) {
	//============在连接销毁之前，查询conn的Name，Home属性=====
	if name, ok := connection.GetProperty("Name"); ok {
		fmt.Println("Conn Property Name = ", name)
	}

	if home, ok := connection.GetProperty("Home"); ok {
		fmt.Println("Conn Property Home = ", home)
	}
	//===================================================

	fmt.Println("DoConnectionLost is Called ... ")
}

func main() {
	s := znet.NewServer()
	s.SetOnConnStart(OnStart)
	s.SetOnConnStop(OnStop)
	s.AddRouter(0, new(PingRouter))
	s.AddRouter(1, new(HelloRouter))
	s.Start()
	quit(s)
}

func quit(s ziface.IServer) {
	quitChan := make(chan os.Signal, 1)
	signal.Notify(quitChan, syscall.SIGINT, syscall.SIGTERM)
	<-quitChan
	s.Stop()
}
