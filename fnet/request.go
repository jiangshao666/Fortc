package fnet
import(
	"github.com/jiangshao666/Fortc/fiface"
)


type Request struct {
	conn fiface.IConnection
	msg fiface.IMessage
}

func NewRequest(conn fiface.IConnection, msg fiface.IMessage) *Request {
	return &Request{
		conn: conn,
		msg: msg,
	}
}


func (r *Request)GetConnection() fiface.IConnection {
	return r.conn;	
}

func (r *Request)GetMsgId() uint32 {
	return r.msg.GetMsgId();
}

func (r *Request)GetMsgData() []byte {
	return r.msg.GetMsgData()
}