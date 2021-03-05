package main

import (
	"tcpwchat/tcpserver/handle"
	"tcpwchat/tcpserver/server"
)

func main(){
	server.Run(&handle.Handle{})
}