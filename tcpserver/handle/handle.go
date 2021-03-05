package handle

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"sync"
	"sync/atomic"
	"tcpwchat/tcpserver/user"
)

type Handle struct{
	User *user.User
	Con    net.Conn  //用户管理连接通道，供发送协程发送数据
	Closeflag uint32 //原子性关闭标志
	wg *sync.WaitGroup
}
var Handlemap map[int32]*Handle

func init() {
	Handlemap = make(map[int32]*Handle,1024)
}
func (h *Handle)initwg(){
	h.wg = &sync.WaitGroup{}
}
func (h *Handle)Handlefunc(con net.Conn){
	h.initwg()
	atomic.StoreUint32(&h.Closeflag,1)
	msg := make([]byte,4)
	n, err := con.Read(msg)
	if err != nil{
		fmt.Println("读数据错误:",err)
		return
	}
	fmt.Println("初始读了",n,"位")
	bytebuff :=bytes.NewBuffer(msg[:4])
	var userid int32
	binary.Read(bytebuff,binary.BigEndian,&userid)
	h.User = &user.User{
		Id:     userid,
		Name:   "",
		Passwd: "",
		Addr:   con.RemoteAddr().String(),
	}
	h.Con = con
	Handlemap[userid] = h
	h.wg.Add(2)
	go func(){
		defer h.wg.Done()
		h.Con.Write([]byte("登录成功!"))
	}()
	go h.keepAlive()

	h.wg.Wait()
}

func (h *Handle)Close()error {
	if atomic.LoadUint32(&h.Closeflag) == 0 {
		fmt.Println("connation is clossing or closed!")
		return nil
	}else{
		atomic.StoreUint32(&h.Closeflag,0)
		h.Con.Close()
		return nil
	}
}

func (h *Handle)keepAlive(){
	defer h.wg.Done()
	sendch := make(chan byte)
	for {
		buf:= make([]byte,2048)
		n, err := h.Con.Read(buf)
		if err != nil {
			fmt.Println(h.User.Addr,":读取错误",err)
		}
		useridbuff := bytes.NewBuffer(buf[0:4])
		var touserid int32
		binary.Read(useridbuff,binary.BigEndian,&touserid)
		toh ,ok := Handlemap[touserid]
		if ok != true {
			useridstring := fmt.Sprintf("用户%d未登陆!")
			go h.send(sendch ,[]byte(useridstring))
			//h.Con.Write([]byte(useridstring))
		}else{
			tempid :=h.User.Id
			tempbuff := bytes.NewBuffer([]byte{})
			binary.Write(tempbuff,binary.BigEndian,&tempid)
			sendbuff := append(tempbuff.Bytes()[:4],buf[4:n]...)
			go toh.send(sendch,sendbuff)
		}
		sendch <- '1'
	}
}

func (h *Handle)send(ch <-chan byte,msg []byte){
	for{
		 <- ch
		h.Con.Write(msg)
	}
}
