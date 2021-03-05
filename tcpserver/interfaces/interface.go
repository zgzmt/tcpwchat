package handler

import "net"
//定义处理接口
type Handler interface {
	 Handlefunc(con net.Conn)
	 Close()error
}