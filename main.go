package main

import (
	"fmt"
	"github.com/0RAJA/zinx/ziface"
	"github.com/0RAJA/zinx/znet"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

type PingRouter struct {
	znet.BaseRouter
}

func (p *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("[Call PingRouter Handle]")
	fmt.Println(string(request.GetData()))
	err := request.GetIConnect().SendMsg(znet.NewMessage(0, []byte("ping "+strconv.Itoa(int(request.GetMsgID())))))
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
	err := request.GetIConnect().SendMsg(znet.NewMessage(0, []byte("hello "+strconv.Itoa(int(request.GetMsgID())))))
	if err != nil {
		log.Println("ping...ping...ping err", err)
	}
}

func quit(s ziface.IServer) {
	quitChan := make(chan os.Signal, 1)
	signal.Notify(quitChan, syscall.SIGINT, syscall.SIGTERM)
	<-quitChan
	s.Stop()
}

func main() {
	s := znet.NewServer()
	s.AddRouter(0, new(PingRouter))
	s.AddRouter(1, new(HelloRouter))
	s.Start()
	quit(s)
}
