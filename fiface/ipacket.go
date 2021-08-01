package fiface

import "net"

type IPacket interface {
	Unpack(conn *net.TCPConn) (IMessage, error)
	Pack(msg IMessage) ([]byte, error)
}