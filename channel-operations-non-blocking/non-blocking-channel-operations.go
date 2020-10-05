package main

import (
	"fmt"
	"time"
)

func main() {

	//使用带有 default 分支的select 可以达到非阻塞的效果

	c1 := make(chan string)

	go func() {
		time.Sleep(time.Second * 5)
		c1 <- "Hello World!"
	}()

	select {
	case msg := <-c1:
		fmt.Println(msg)
	default:
		fmt.Println("c1没有收到消息")
	}
}
