package main

import (
	"fmt"
	"net"
	"time"

	"github.com/jiangshao666/Fortc/fnet"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("client start err, exit!", err)
		return
	}

	for {
		//发封包message消息
		dp := fnet.NewPacket()
		msgPacket := fnet.NewMessage(1, []byte("UserLogin Demo Test MsgID=1, {Name: Jiangshao}"))
		msg, _ := dp.Pack(msgPacket)
		n, err := conn.Write(msg)
		if err != nil {
			fmt.Println("write error err ", err)
			return
		}
		fmt.Println("send content len", n)

	
		// 准备接收服务器响应
		connTcp := conn.(*net.TCPConn)
		respMsg,err := dp.Unpack(connTcp)

		fmt.Println("==> Test Router:[Ping] Recv Msg: ID=", respMsg.GetMsgId(), ", len=", respMsg.GetMsgLen(), ", data=", string(respMsg.GetMsgData()))
		

		time.Sleep(1 * time.Second)
	}
}