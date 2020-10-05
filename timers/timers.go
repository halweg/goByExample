package main

import (
	"fmt"
	"time"
)

func main() {

	//提供一个明确的等待时间，定时器会产生一个 channel 用于通知
	timer1 := time.NewTimer(time.Second * 2)

	// 如果什么都不做，通道 timer.C 会一直等待（阻塞）消息
	// 定时器失效的值之前，将一直阻塞。
	<-timer1.C
	fmt.Println("timer1 is expired")

	//定时器和sleep的区别，定时器的灵活之处在于可以自由取消定时器
	timer2 := time.NewTimer(time.Second * 5)
	go func() {
		<-timer2.C
		fmt.Println("timer2 结束了")
	}()

	//timer2 被提前 stop 会发现 “timer2 结束了" 这一行并没有被运行
	stopTime2 := timer2.Stop()
	if stopTime2 {
		fmt.Println("timer2 被 stop")
	}



}
