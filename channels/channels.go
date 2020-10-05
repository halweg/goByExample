package main

import "fmt"

func main()  {

	msg := make(chan string)


	go func() {
		msg <- "ping ping ping , are you online?"
	}()

	//默认发送和接收操作是阻塞的，直到发送方和接收方都准备完毕

	msging := <-msg

	fmt.Println(msging)

}
