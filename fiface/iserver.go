package fiface

type IServer interface {
	Start()
	Stop()
	Serve()
	AddRouter(msgId uint32, router IRouter )
	GetConnMgr() IConnMgr
	GetMsgHandler() IMsgHandler
	Packet() IPacket

	SetOnConnStart( func(conn IConnection) )
	SetOnConnStop( func(conn IConnection))
	CallOnConnStart(conn IConnection)
	CallOnConnStop(conn IConnection)



}