package server

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	handler "tcpwchat/tcpserver/interfaces"
	"tcpwchat/tcpserver/user"
)

func Run(h handler.Handler) {
	fmt.Println(len(user.GetUsers()))
	listent, err := net.Listen("tcp", "127.0.0.1:9090")
	if err != nil {
		fmt.Println("listen err:", err)
		return
	}
	defer listent.Close()
	signch := make(chan os.Signal,1)
	signal.Notify(signch,syscall.SIGHUP,syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		switch <-signch{
		case syscall.SIGHUP,syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT :
			h.Close()
		}
	}()
	for {
		con, err := listent.Accept()
		if err != nil {
			fmt.Println("accept err:", err)
			return
		}
		go h.Handlefunc(con)
	}
}
//func handleconfun(con net.Conn) {
//	defer con.Close()
//
//	u := user.User{
//		Id:     0,
//		Name:   "",
//		Passwd: "",
//		Addr:   con.RemoteAddr().String(),
//	}
//	//user.AddUsers(u)
//	for {
//		var b []byte
//		_, err := con.Read(b)
//		if err != nil {
//			fmt.Print(err)
//		}
//
//	}
//}
