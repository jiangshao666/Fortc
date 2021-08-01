package fiface


type IConnMgr interface {
	Add(connId uint32, conn IConnection)
	Get(connId uint32) (IConnection, error)
	Remove(conn IConnection)
	GetLen() uint32
	ClearConn()
	ClearOneConn(connId uint32)
}