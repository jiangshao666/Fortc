package fnet

import (
	"sync"
	"fmt"
	"errors"
	"github.com/jiangshao666/Fortc/fiface"
)

type ConnMgr struct {
	connections map[uint32] fiface.IConnection
	lockConn sync.RWMutex
}


func NewConnMgr() *ConnMgr {
	return &ConnMgr{
		connections: make(map[uint32] fiface.IConnection),
	}
}

func (cm *ConnMgr) Add(connId uint32, conn fiface.IConnection) {
	cm.lockConn.Lock()
	defer cm.lockConn.Unlock()
	cm.connections[connId] = conn
	fmt.Println("Add connection connId", connId, " successful, num", cm.GetLen())
}

func (cm *ConnMgr)Get(connId uint32) (fiface.IConnection, error) {
	cm.lockConn.RLock()
	defer cm.lockConn.RUnlock()
	if conn, ok := cm.connections[connId]; ok {
		return conn, nil
	}
	return nil, errors.New(fmt.Sprintf("conn of connId %d not exist", connId))
}

func (cm *ConnMgr)Remove(conn fiface.IConnection) {
	cm.lockConn.Lock()
	defer cm.lockConn.Unlock()

	delete(cm.connections, conn.GetConnId())
	fmt.Println("Remove connection connId ", conn.GetConnId(), " successful, num", cm.GetLen())
}

func (cm *ConnMgr)GetLen() uint32 {
	return uint32(len(cm.connections))
}

func (cm *ConnMgr)ClearConn() {
	cm.lockConn.Lock()
	defer cm.lockConn.Unlock()

	for _, conn := range cm.connections {
		conn.Stop()
		delete(cm.connections, conn.GetConnId())
	}
	fmt.Println("Clear All Conns successful, num", cm.GetLen())
}

func (cm *ConnMgr)ClearOneConn(connId uint32) {
	cm.lockConn.Lock()
	defer cm.lockConn.Unlock()

	if conn, ok := cm.connections[connId]; ok {
		conn.Stop()
		delete(cm.connections, conn.GetConnId())
		fmt.Println("ClearOneConn connId ", connId, " successful, num", cm.GetLen())
		return
	}
	fmt.Println("ClearOneConn connId ", connId, " falied")
	return
}