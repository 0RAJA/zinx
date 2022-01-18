package main

import (
	"fmt"
	"github.com/0RAJA/zinx/znet"
	"io"
	"log"
	"net"
	"strconv"
	"time"
)

func main() {
	fmt.Println("Client Test Start")

	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal("conn err:", err)
	}
	defer conn.Close()
	dp := znet.NewDataPack()
	go func() {
		dp := znet.NewDataPack()
		for {
			headData := make([]byte, dp.GetHeadLen())
			if _, err := io.ReadFull(conn, headData); err != nil {
				log.Println(err)
				break
			}
			msg, err := dp.UnPack(headData)
			if err != nil {
				log.Println(err)
				break
			}
			if msg.GetDataLen() > 0 {
				data := make([]byte, msg.GetDataLen())
				if _, err := io.ReadFull(conn, data); err != nil {
					log.Println(err)
					break
				}
				msg.SetData(data)
				fmt.Println("[Client Resv] ", string(data))
			}
		}
	}()
	for i := 0; i < 5; i++ {
		message := znet.NewMessage(uint32(i%2), []byte("hello world"+strconv.Itoa(i)))
		data, err := dp.Pack(message)
		if err != nil {
			log.Println("client pack err:", err)
			continue
		}
		_, err = conn.Write(data)
		if err != nil {
			log.Println(err)
			continue
		}
	}
	time.Sleep(3 * time.Second)
}
