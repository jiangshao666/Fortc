package fiface

import (
	"net"
)

type IConnection interface {

	Start()
	Stop()

	GetConnId() uint32
	GetConn() *net.TCPConn
	GetRemoteAddr() net.Addr

	SetProperty(key string, value interface{})
	GetProperty(key string) (interface{}, error)
	RemoveProperty(key string)

	SendMsg(msg IMessage) error
	SendBuffMsg(msg IMessage) error
}