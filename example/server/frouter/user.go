package frouter

import(
	"fmt"
	"github.com/jiangshao666/Fortc/fnet"
	"github.com/jiangshao666/Fortc/fiface"
)

type UserRouter struct {
	fnet.Router
}

func (pr *UserRouter)Handle(request fiface.IRequest) {
	fmt.Println("Call UserRouter Handle")
	//先读取客户端的数据，再回写ping...ping...ping
	fmt.Println("recv from client : msgId=", request.GetMsgId(), ", data=", string(request.GetMsgData()))

	msg := fnet.NewMessage(1,[]byte("name: Jiangshao"))

	err := request.GetConnection().SendBuffMsg(msg)
	if err != nil {
		fmt.Println("SendMsg failed", err)
	}
}