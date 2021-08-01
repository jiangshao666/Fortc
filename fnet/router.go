package fnet

import
(
	"github.com/jiangshao666/Fortc/fiface"
)

// 抽象的类，需要根据业务重写这部分接口
type Router struct {}


func (router *Router)PreHandle(request fiface.IRequest) {}

func (router *Router)Handle(request fiface.IRequest) {}

func (router *Router)PostHandle(request fiface.IRequest) {}