package fiface

type IMsgHandler interface {
	DoMsgHandler(IRequest) 
	AddRouter(msgId uint32, router IRouter)
	StartWorkPool()
	PutReqToTaskQueue(request IRequest)
}