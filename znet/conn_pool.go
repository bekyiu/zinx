package znet

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"sync"
	"zinx/ziface"
)

type ConnPool struct {
	// key: connId
	connMap map[uint32]ziface.IConnection
	// map不是线程安全的
	lock sync.RWMutex
	// 连接池一共能容纳多少个拦截
	size int
}

func NewConnPool() *ConnPool {
	pool := ConnPool{
		connMap: make(map[uint32]ziface.IConnection),
		size:    100,
	}
	return &pool
}

func (p *ConnPool) Remove(id uint32) {
	p.lock.Lock()
	defer p.lock.Unlock()
	p.connMap[id].GetTCPConn().Close()
	fmt.Println("remove and close conn[", id, "]")
	delete(p.connMap, id)
}

func (p *ConnPool) Get(id uint32) (ziface.IConnection, error) {
	p.lock.RLock()
	defer p.lock.RUnlock()

	if v, ok := p.connMap[id]; ok {
		return v, nil
	} else {
		return nil, errors.New("connId: " + strconv.Itoa(int(id)) + " not exist")
	}
}

func (p *ConnPool) Num() int {
	p.lock.RLock()
	defer p.lock.RUnlock()
	l := len(p.connMap)
	return l
}

func (p *ConnPool) Clear() {
	p.lock.Lock()
	defer p.lock.Unlock()

	for id, conn := range p.connMap {
		conn.Stop()
		delete(p.connMap, id)
	}

	fmt.Println("clear all connections")
}

var connId uint32 = 0

func (p *ConnPool) Allocate(conn *net.TCPConn, server ziface.IServer) (ziface.IConnection, error) {
	p.lock.Lock()
	defer p.lock.Unlock()

	if len(p.connMap) < p.size {
		conn := NewConnection(conn, connId, server)
		p.connMap[connId] = conn
		fmt.Println("allocate a new connection ", "[", connId, "]")
		connId++
		return conn, nil
	}
	return nil, errors.New("to many connections")
}
