package frouter

import(
	"fmt"
	"github.com/jiangshao666/Fortc/fnet"
	"github.com/jiangshao666/Fortc/fiface"
)

type PingRouter struct {
	fnet.Router
}

func (pr *PingRouter)Handle(request fiface.IRequest) {
	fmt.Println("Call PingRouter Handle")
	//先读取客户端的数据，再回写ping...ping...ping
	fmt.Println("recv from client : msgId=", request.GetMsgId(), ", data=", string(request.GetMsgData()))

	msg := fnet.NewMessage(0,[]byte("ping...ping...ping"))

	err := request.GetConnection().SendBuffMsg(msg)
	if err != nil {
		fmt.Println("SendMsg failed", err)
	}
}
