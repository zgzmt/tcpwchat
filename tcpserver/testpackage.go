package testpackage

import (
	"fmt"
	"net"
)

func Test() {
	fmt.Println("zheshiyigeceshi!!")
}

func Runtest() {
	listent, err := net.Listen("tcp", ":9090")
	if err != nil {
		fmt.Printf("listen err:%s", err)
		return
	}
	defer listent.Close()
	for {
		con, err := listent.Accept()
		if err != nil {
			fmt.Println("accept err :", err)
			continue
		}
		go handleconfun(con)
	}
}

func handleconfun(con net.Conn) {
	defer con.Close()

}
