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
	fmt.Println("[Call PingRouter Handle]")
	fmt.Println(string(request.GetData()))
	err := request.GetIConnect().SendMsgWithBuff(znet.NewMessage(0, request.GetData()))
	if err != nil {
		log.Println("ping...ping...ping err", err)
	}
}

type HelloRouter struct {
	znet.BaseRouter
}

func (h *HelloRouter) Handle(request ziface.IRequest) {
	fmt.Println("[Call HelloRouter Handle]")
	fmt.Println(string(request.GetData()))
	err := request.GetIConnect().SendMsg(znet.NewMessage(0, request.GetData()))
	if err != nil {
		log.Println("ping...ping...ping err", err)
	}
}

func OnStart(connection ziface.IConnection) {
	if err := connection.SendMsgWithBuff(znet.NewMessage(0, []byte(fmt.Sprint("[Connection]", connection.GetConnID(), " is start")))); err != nil {
		log.Println(err)
	}
}

func OnStop(connection ziface.IConnection) {
	if err := connection.SendMsgWithBuff(znet.NewMessage(0, []byte(fmt.Sprint("[Connection]", connection.GetConnID(), " is stop")))); err != nil {
		log.Println(err)
	}
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
