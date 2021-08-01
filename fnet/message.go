package fnet


type Message struct {
	msgId uint32
	msgData []byte
	msgLen uint32
}

func NewMessage(msgId uint32, data []byte) *Message{
	return &Message{
		msgId: msgId,
		msgData: data,
		msgLen: uint32(len(data)),
	}
}

func (m *Message) GetMsgId() uint32 {
	return m.msgId
}

func (m *Message) GetMsgData() []byte {
	return m.msgData
}

func (m* Message) GetMsgLen() uint32 {
	return m.msgLen
}

