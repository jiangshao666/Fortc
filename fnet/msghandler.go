package fnet
import(
	"fmt"
	"github.com/jiangshao666/Fortc/fiface"
	"github.com/jiangshao666/Fortc/utils"
)


type MsgHandler struct {
	routers map[uint32] fiface.IRouter
	WorkPoolSize uint16
	TaskQueue []chan fiface.IRequest
}


func NewMsgHandler() *MsgHandler{
	return &MsgHandler{
		routers: make(map[uint32]fiface.IRouter),
		WorkPoolSize: utils.GlobalConfig.WorkPoolSize,
		TaskQueue: make([]chan fiface.IRequest, utils.GlobalConfig.WorkPoolSize),
	}
}

func (mh *MsgHandler)DoMsgHandler(request fiface.IRequest) {
	msgId := request.GetMsgId()
	router, ok := mh.routers[msgId]
	if !ok {
		fmt.Println("No router found of msgId", msgId)
	} else {
		router.PreHandle(request)
		router.Handle(request)
		router.PostHandle(request)
	}
}

func (mh *MsgHandler)AddRouter(msgId uint32, router fiface.IRouter) {
	if mh.routers == nil {
		mh.routers = make(map[uint32]fiface.IRouter)
	}
	mh.routers[msgId] = router
	fmt.Println("Add Router to MsgHandler", msgId)
}

func (mh *MsgHandler) StartWorker(workId uint16, workChan chan fiface.IRequest) {
	fmt.Printf("[StartWorker %d successful] \n", workId)
	for {
		select {
			case request := <- workChan :
				mh.DoMsgHandler(request)
		}
	}
}

// 启动工作goroutine池
func (mh *MsgHandler)StartWorkPool() {
	for i:=0; i<int(mh.WorkPoolSize); i++ {
		mh.TaskQueue[i] = make(chan fiface.IRequest, utils.GlobalConfig.MaxMsgChanLen)
		go mh.StartWorker(uint16(i), mh.TaskQueue[i])
	}
}

// 均衡的将请求放到队列中, 根据连接Id取余数
func (mh *MsgHandler)PutReqToTaskQueue(request fiface.IRequest) {
	taskId := request.GetConnection().GetConnId() % uint32(mh.WorkPoolSize)
	taskChan := mh.TaskQueue[taskId]
	taskChan <- request
}