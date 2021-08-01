package main

import(
	"fmt"
	"github.com/jiangshao666/Fortc/fnet"
	"github.com/jiangshao666/Fortc/fiface"
	"github.com/jiangshao666/Fortc/example/server/frouter"
)

func doConnStart(conn fiface.IConnection) {
	fmt.Println("onConnectionStart is called", conn.GetConnId())
	conn.SetProperty("Name", "Jiangshao")
	conn.SetProperty("Token", "token12345677")
}

func doConnStop(conn fiface.IConnection) {
	fmt.Println("onConnectionStop is called", conn.GetConnId())
	if Name, err := conn.GetProperty("Name"); err ==nil {
		fmt.Println("Conn Name", Name)
	}
	if Token, err := conn.GetProperty("Token"); err ==nil {
		fmt.Println("Conn Token", Token)
	}
}


func main() {
	s := fnet.NewServer()

	s.SetOnConnStart(doConnStart)
	s.SetOnConnStop(doConnStop)

	s.AddRouter(0, &frouter.PingRouter{})
	s.AddRouter(1, &frouter.UserRouter{})

	s.Serve()
}