package znet

import (
	"fmt"
	"github.com/0RAJA/zinx/global"
	"github.com/0RAJA/zinx/ziface"
	"log"
	"net"
)

type Server struct {
	Name        string
	IPVersion   string //tcp4 or other
	IP          string
	Port        int
	MsgHandle   ziface.IMsgHandler //当前Server由用户绑定的回调router,也就是Server注册的链接对应的处理业务
	ConnMgr     ziface.IConnManager
	OnConnStart []func(connection ziface.IConnection)
	OnConnStop  []func(connection ziface.IConnection)
	over        chan struct{}
}

func (s *Server) CallOnConnStart(connection ziface.IConnection) {
	for _, f := range s.OnConnStart {
		f(connection)
	}
}

func (s *Server) CallOnConnStop(connection ziface.IConnection) {
	for _, f := range s.OnConnStop {
		f(connection)
	}
}

func (s *Server) SetOnConnStart(hookFuncs ...func(connection ziface.IConnection)) {
	for _, f := range hookFuncs {
		if f != nil {
			s.OnConnStart = append(s.OnConnStart, f)
		}
	}
}

func (s *Server) SetOnConnStop(hookFuncs ...func(connection ziface.IConnection)) {
	for _, f := range hookFuncs {
		if f != nil {
			s.OnConnStop = append(s.OnConnStop, f)
		}
	}
}

func NewServer() *Server {
	return &Server{
		Name:      global.ServerSetting.Name,
		IPVersion: "tcp4",
		IP:        global.ServerSetting.IP,
		Port:      global.ServerSetting.Port,
		MsgHandle: NewMsgHandler(),
		ConnMgr:   NewConnManager(),
		over:      make(chan struct{}),
	}
}

/*实现 ziface.IServer 里的全部接口方法*/

func (s *Server) AddRouter(msgID uint32, router ziface.IRouter) {
	s.MsgHandle.AddRouter(msgID, router)
}

func (s *Server) GetConnMgr() ziface.IConnManager {
	return s.ConnMgr
}

func (s *Server) Start() {
	fmt.Printf("Server listener at IP:%s,Port:%d,is starting\n", s.IP, s.Port)
	go func() {
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			log.Println("resolve tcp addr err: ", err)
			return
		}
		lister, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			log.Println("listen:", s.IP, "err:", err)
			return
		}
		//监听成功
		fmt.Println(s.Name, " listen success")
		//TODO server.go 应该有一个自动生成ID的方法
		var cid uint32 = 0
		//启动工作池和消息队列
		s.MsgHandle.StartWorkerPool()
		//启动连接服务
		for {
			//3.1 阻塞等待客户端建立连接请求
			conn, err := lister.AcceptTCP()
			if err != nil {
				log.Println("Accept err ", err)
				continue
			}
			//设置服务器最大连接控制,如果超过最大连接，那么则关闭此新的连接
			if s.ConnMgr.Len() >= global.ServerSetting.MaxConn {
				conn.Close()
				continue
			}
			//3.3 处理该新连接请求的 业务 方法， 此时应该有 handler 和 conn是绑定的
			dealConn := NewConnection(s, conn, cid, s.MsgHandle)
			cid++
			//开启处理业务
			go dealConn.Start()
		}
	}()
}

func (s *Server) Stop() {
	fmt.Println("[STOP] Zinx server , name:", s.Name)
	//将所有连接清除然后再退出
	s.ConnMgr.Clear()
	close(s.over)
}

func (s *Server) Server() {
	s.Start()
	//TODO Server.Serve() 是否在启动服务的时候 还要处理其他的事情呢 可以在这里添加
	<-s.over
}
