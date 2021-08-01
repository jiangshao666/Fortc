package fiface

type IRequest interface {
	GetConnection() IConnection
	GetMsgId() uint32
	GetMsgData() []byte	
}