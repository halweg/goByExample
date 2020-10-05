package main

import (
	"fmt"
	"time"
)

func main() {

	ticker := time.NewTicker(time.Millisecond * 100)

	//定时器和打点器都是利用了 channel 的阻塞来实现的
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	time.Sleep(time.Millisecond * 1200)
	// 一旦一个打点停止了，将不能再从它的通道中接收到值。
	ticker.Stop()
	fmt.Println("Ticker stopped")

}
