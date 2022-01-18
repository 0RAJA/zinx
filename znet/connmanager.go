package znet

import (
	"errors"
	"fmt"
	"github.com/0RAJA/zinx/ziface"
	"log"
	"sync"
)

type ConnManager struct {
	l        sync.RWMutex
	connMaps map[uint32]ziface.IConnection
}

func NewConnManager() *ConnManager {
	return &ConnManager{connMaps: make(map[uint32]ziface.IConnection)}
}

func (c *ConnManager) Add(connection ziface.IConnection) {
	if _, ok := c.connMaps[connection.GetConnID()]; ok {
		log.Println("connID is existed,connID:", connection.GetConnID())
	}
	defer c.l.Unlock()
	c.l.Lock()
	c.connMaps[connection.GetConnID()] = connection
}

func (c *ConnManager) Remove(conn ziface.IConnection) {
	defer c.l.Unlock()
	c.l.Lock()
	delete(c.connMaps, conn.GetConnID())
}

func (c *ConnManager) Get(connID uint32) (ziface.IConnection, error) {
	defer c.l.RUnlock()
	c.l.RLock()
	if conn, ok := c.connMaps[connID]; ok {
		return conn, nil
	}
	return nil, errors.New(fmt.Sprint("connID:", connID, " is not FOUND!"))
}

func (c *ConnManager) Len() int {
	return len(c.connMaps)
}

func (c *ConnManager) Clear() {
	defer c.l.Unlock()
	c.l.Lock()
	for _, conn := range c.connMaps {
		conn.Stop()
	}
	c.connMaps = make(map[uint32]ziface.IConnection)
}
