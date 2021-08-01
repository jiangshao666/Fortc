package fnet

import(
	"bytes"
	"io"
	"fmt"
	"errors"
	"net"
	"encoding/binary"
	"github.com/jiangshao666/Fortc/fiface"
	"github.com/jiangshao666/Fortc/utils"
)

var defaultMsgHeadLen=8

type Packet struct {}

func NewPacket() fiface.IPacket{
	return &Packet{}
}

// 反序列化
func (p *Packet) Unpack(conn *net.TCPConn) (fiface.IMessage, error) {
	headBytes := make([]byte, defaultMsgHeadLen)
	if _, err :=io.ReadFull(conn, headBytes); err !=nil {
		fmt.Println("read msg head error ", err)
		return nil, err
	}

	headReader := bytes.NewReader(headBytes)
	// 解析包体长度
	msgLen := uint32(0)
	if err := binary.Read(headReader, binary.LittleEndian, &msgLen); err != nil {
		fmt.Println("read msgLen error ", err)
		return nil, err
	}
	// 判断包体长度
	if utils.GlobalConfig.MaxPacketSize > 0 && msgLen >= utils.GlobalConfig.MaxPacketSize {
		return nil, errors.New(fmt.Sprintf("recv too large packet, size=%d", msgLen))
	}

	// 解析msgId
	msgId := uint32(0)
	if err := binary.Read(headReader, binary.LittleEndian, &msgId); err != nil {
		fmt.Println("read msgId error ", err)
		return nil, err
	}

	dataBuff := make([]byte, msgLen)
	if _, err :=io.ReadFull(conn, dataBuff); err !=nil {
		fmt.Println("read msg data error ", err)
		return nil, err
	}
	return NewMessage(msgId, dataBuff), nil
}

// 序列化
func (p *Packet)Pack(msg fiface.IMessage) ([]byte, error) {
	byteBuff := bytes.NewBuffer([]byte{})
	// 写长度
	if err := binary.Write(byteBuff, binary.LittleEndian, msg.GetMsgLen()); err !=nil {
		return nil, err
	}
	// 写msgId
	if err := binary.Write(byteBuff,binary.LittleEndian, msg.GetMsgId()); err !=nil {
		return nil, err
	}

	// 写msgData
	if err := binary.Write(byteBuff,binary.LittleEndian, msg.GetMsgData()); err !=nil {
		return nil, err
	}

	return byteBuff.Bytes(), nil
}