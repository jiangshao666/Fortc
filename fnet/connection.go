package fnet
import(
	"net"
	"sync"
	"errors"
	"fmt"

	"github.com/jiangshao666/Fortc/fiface"
	"github.com/jiangshao666/Fortc/utils"
)

type Connection struct {
	TCPServer fiface.IServer
	connId uint32
	Conn *net.TCPConn

	MsgHandler fiface.IMsgHandler
	msgChan chan []byte
	msgBuffChan chan []byte
	
	properties map[string]interface{}
	lockProp sync.RWMutex

	isClosed bool
}

func NewConnection(s fiface.IServer, connId uint32, conn *net.TCPConn, msgHandler fiface.IMsgHandler) *Connection {
	connection := &Connection {
		TCPServer: s,
		connId: connId,
		Conn: conn,
		MsgHandler: msgHandler,

		msgChan: make(chan []byte),
		msgBuffChan: make(chan []byte, utils.GlobalConfig.MaxMsgChanLen),
	}
	return connection
}


func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is running")
	defer c.Stop()
	for {
		msg, err := c.TCPServer.Packet().Unpack(c.GetConn())
		if err != nil {
			fmt.Println("unpack error ", err)
			return
		}

		r := NewRequest(c, msg)
		go c.MsgHandler.DoMsgHandler(r)
	}
}

func (c *Connection) StartWriter() {
	fmt.Println("Writer Goroutine is running")
	defer fmt.Println(c.Conn.RemoteAddr().String(), "[conn Writer exit!]")

	for {
		select {
			case data:= <- c.msgChan:
				//有数据要写给客户端
				if _, err := c.Conn.Write(data); err != nil {
					fmt.Println("Send Data error:, ", err, " Conn Writer exit")
					return
				}
			case data, ok := <-c.msgBuffChan:
				if ok {
					//有数据要写给客户端
					if _, err := c.Conn.Write(data); err != nil {
						fmt.Println("Send Buff Data error:, ", err, " Conn Writer exit")
						return
					}
				} else {
					fmt.Println("msgBuffChan is Closed")
					break
				}
		}
	}
}

func(c *Connection) Start() {
	fmt.Println("Connection start", c.connId)
	c.TCPServer.CallOnConnStart(c)
	go c.StartReader()

	go c.StartWriter()
}

func(c *Connection) Stop() {
	
	if c.isClosed {
		return
	}

	c.GetConn().Close()
	c.TCPServer.CallOnConnStop(c)
	close(c.msgBuffChan)
	c.isClosed = true
}

func(c *Connection)GetConnId() uint32 {
	return c.connId
}

func(c *Connection)GetConn() *net.TCPConn {
	return c.Conn
}

func(c *Connection)GetRemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func(c *Connection)SetProperty(key string, value interface{}) {
	c.lockProp.Lock()
	defer c.lockProp.Unlock()
	if c.properties == nil {
		c.properties = make(map[string]interface{})
	}
	c.properties[key] = value

}

func(c *Connection)GetProperty(key string) (interface{}, error) {
	c.lockProp.RLock()
	defer c.lockProp.RUnlock()
	v, ok := c.properties[key]
	if ok  {
		return v, nil
	}
	return nil, errors.New("property is not exist")
}

func(c *Connection)RemoveProperty(key string) {
	c.lockProp.Lock()
	defer c.lockProp.Unlock()
	delete(c.properties, key)
}

func(c *Connection)SendMsg(msg fiface.IMessage) error {
	if c.isClosed {
		return errors.New("SendMsg to closed connection")
	}
	bytes, err := c.TCPServer.Packet().Pack(msg)
	if err != nil {
		fmt.Println("SendMsg Pack failed")
		return err
	}
	c.msgChan <- bytes
	return nil
}

func(c *Connection)SendBuffMsg(msg fiface.IMessage) error {
	if c.isClosed {
		return errors.New("SendMsg to closed connection")
	}
	bytes, err := c.TCPServer.Packet().Pack(msg)
	if err != nil {
		fmt.Println("SendBuffMsg Pack failed")
		return err
	}
	c.msgBuffChan <- bytes
	return nil
}