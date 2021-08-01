package fiface

type IMessage interface {
	GetMsgId() uint32
	GetMsgData() []byte
	GetMsgLen() uint32
}