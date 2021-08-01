package fnet

import (
	"net"
	"fmt"
	"github.com/jiangshao666/Fortc/utils"
	"github.com/jiangshao666/Fortc/fiface"
)

type Server struct {
	Name string
	IPVersion string
	IP string
	Port uint16

	connMgr fiface.IConnMgr
	msgHandler fiface.IMsgHandler
	onConnStart func(conn fiface.IConnection)
	onConnStop func(conn fiface.IConnection)

	packet fiface.IPacket
}

func NewServer() fiface.IServer {
	return &Server {
		Name: utils.GlobalConfig.Name,
		IPVersion: "tcp4",
		IP: utils.GlobalConfig.IP,
		Port: utils.GlobalConfig.Port,

		msgHandler: NewMsgHandler(),
		connMgr: NewConnMgr(),
		packet: NewPacket(),

	}
}

func (s *Server) Start() {
	fmt.Printf("Fortc Server Start Listen on ip: %s port: %d \n", s.IP, s.Port)

	fmt.Printf("[Zinx] Version: %s, MaxConn: %d, MaxPacketSize: %d\n",
		utils.GlobalConfig.Version,
		utils.GlobalConfig.MaxConn,
		utils.GlobalConfig.MaxPacketSize)
	
	go func() {
		s.msgHandler.StartWorkPool()

		addr, err := net.ResolveTCPAddr(s.IPVersion,fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("ip and port resolve failed", err)
			return
		}
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("ListenTCP failed", err)
			return
		}
		
		var connID uint32
		connID = 0
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("AcceptTCP failed", err)
				continue
			}
			fmt.Println("Get conn remote addr = ", conn.RemoteAddr().String())

			//3.2 设置服务器最大连接控制,如果超过最大连接，那么则关闭此新的连接
			if s.connMgr.GetLen() >= utils.GlobalConfig.MaxConn {
				fmt.Println("Conn Over MaxConn!! ")
				conn.Close()
				continue
			}
			connection := NewConnection(s, connID, conn, s.msgHandler)
			s.connMgr.Add(connID, connection)
			connID++
			go connection.Start()
		}
	}()
}

func (s *Server) Stop() {
	fmt.Println("[STOP] Fortc server , name ", s.Name)

	//将其他需要清理的连接信息或者其他信息 也要一并停止或者清理
	s.connMgr.ClearConn()
}

func (s *Server) Serve() {
	s.Start()
	// 阻塞住
	select {}
}

func (s *Server) AddRouter(msgId uint32, router fiface.IRouter) {
	s.msgHandler.AddRouter(msgId, router)
}

func (s *Server) GetConnMgr() fiface.IConnMgr {
	return s.connMgr
}

func (s *Server) GetMsgHandler() fiface.IMsgHandler {
	return s.msgHandler
}

func (s *Server) Packet() fiface.IPacket{
	return s.packet
}

func (s *Server) SetOnConnStart( funcHandle func(conn fiface.IConnection)) {
	s.onConnStart = funcHandle
}

func (s *Server) SetOnConnStop( funcHandle func(conn fiface.IConnection)) {
	s.onConnStop = funcHandle
}
	
func (s *Server)  CallOnConnStart(conn fiface.IConnection) {
	fmt.Println("Server CallOnConnStart", conn.GetConnId(), s.onConnStart ==nil)
	if s.onConnStart != nil {
		s.onConnStart(conn)
	}
}
	
func (s *Server) CallOnConnStop(conn fiface.IConnection) {
	if s.onConnStop != nil {
		s.onConnStop(conn)
	}
}