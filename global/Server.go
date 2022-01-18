package global

import (
	"encoding/json"
	"flag"
	"github.com/0RAJA/zinx/ziface"
	"log"
	"os"
)

const (
	DefaultMaxConn        = 12000
	DefaultMaxPacketSize  = 4096
	DefaultConfPath       = "conf/zinx.json"
	DefaultTaskChanSize   = 1024
	DefaultWorkerPoolSize = 10
)

var (
	ServerSetting *Server
)

type Server struct {
	/*Server*/
	TCPServer ziface.IServer
	Name      string
	IP        string
	Port      int
	/*Zinx*/
	Version          string
	MaxConn          int    /*最大连接数*/
	MaxPacketSize    uint32 /*最大包长*/
	WorkerPoolSize   uint32 /*工作池容量*/
	MaxTaskQueueSize uint32 /*消息队列容量*/
	/*conf*/
	ConfPath string
}

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func init() {
	ServerSetting = &Server{
		Name:             "Zinx",
		IP:               "127.0.0.1",
		Port:             8080,
		Version:          "v0.0.1",
		MaxConn:          DefaultMaxConn,
		MaxPacketSize:    DefaultMaxPacketSize,
		MaxTaskQueueSize: DefaultTaskChanSize,
		WorkerPoolSize:   DefaultWorkerPoolSize,
	}
	flag.StringVar(&ServerSetting.ConfPath, "path", DefaultConfPath, "配置文件相对路径")
	flag.Parse()
	ServerSetting.Reload(ServerSetting.ConfPath)
}

func (s *Server) Reload(path string) {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}
	if err := json.Unmarshal(data, ServerSetting); err != nil {
		log.Println(err)
	}
}
