package znet_test

import (
	"fmt"
	"github.com/0RAJA/zinx/ziface"
	"github.com/0RAJA/zinx/znet"
	"log"
	"net"
	"testing"
	"time"
)

//模拟客户端
func testClient(t *testing.T) {
	time.Sleep(3 * time.Second)
	fmt.Println("Client Test Start")

	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		t.Fatal("conn err:", err)
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		message := []byte("hello world")
		_, err := conn.Write(message)
		if err != nil {
			t.Fatal("write error", err)
		}
		buf := make([]byte, znet.MaxBuff)
		cnt, err := conn.Read(buf)
		if err != nil {
			t.Fatal("read err:", err)
		}
		fmt.Printf("[Client Recv:] %v\n", string(buf[:cnt]))
	}
}

type PingRouter struct{}

func (p *PingRouter) PreHandle(request ziface.IRequest) {
	fmt.Println("[Call Router PreHandle]")
	_, err := request.GetIConnect().GetTCPConnection().Write([]byte("before ping ..."))
	if err != nil {
		log.Println("before ping err", err)
	}
}

func (p *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("[Call Router PreHandle]")
	_, err := request.GetIConnect().GetTCPConnection().Write([]byte("ping...ping...ping ..."))
	if err != nil {
		log.Println("ping...ping...ping err", err)
	}
}

func (p *PingRouter) PostHandle(request ziface.IRequest) {
	fmt.Println("[Call Router PreHandle]")
	_, err := request.GetIConnect().GetTCPConnection().Write([]byte("after ping ..."))
	if err != nil {
		log.Println("after ping err", err)
	}
}

func TestServer(t *testing.T) {
	s := znet.NewServer()

	s.AddRouter(new(PingRouter))

	s.Start()
	testClient(t)
	s.Stop()
}

func TestMain(m *testing.M) {
	m.Run()
}
