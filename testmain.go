package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	testpackage "tcpwchat/tcpserver"
	"time"
)

func main111() {
	ch := make(chan int)
	testpackage.Test()
	go func() {
		fmt.Println("读chan:", <-ch)
	}()
	go func() {
		ch <- 1
		fmt.Println("写chan:")
	}()
	time.Sleep(time.Second * 5)
	var wait sync.WaitGroup
	wait.Add(2)
	go func() {
		defer wait.Done()
		for {
			readers := bufio.NewReader(os.Stdin)
			text, _ := readers.ReadString('\n')
			fmt.Println(text)
		}
	}()
	go func() {
		defer wait.Done()
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second * 2)
			fmt.Println("输出,", i)
		}
	}()
	wait.Wait()
}
