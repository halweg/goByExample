package main

import (
	"fmt"
	"time"
)

func main() {

	//1.使用 ticker 进行速率限制,
	request := make(chan int, 5)
	for i :=0; i<5; i++ {
		request<-i
	}
	close(request)
	ticker := time.Tick(time.Millisecond * 200)

	for req := range request {
		<-ticker
		fmt.Println("request", req, time.Now())
	}

	fmt.Println("-------------------------------")
	//2.使用通道缓冲 进行速率限制
	// 有时候我们想临时进行速率限制，并且不影响整体的速率控制，
	// 我们可以通过[通道缓冲](channel-buffering.html)来实现。
	// 这个 `burstyLimiter` 通道用来进行 3 次临时的脉冲型速率限制。
	burstyLimiter := make(chan time.Time, 3)

	// 想将通道填充需要临时改变3次的值，做好准备。
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// 每 200 ms 我们将添加一个新的值到 `burstyLimiter`中，
	// 直到达到 3 个的限制。
	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			burstyLimiter <- t
		}
	}()

	// 现在模拟超过 5 个的接入请求，它们中刚开始的 3 个将
	// 由于受 `burstyLimiter` 的“脉冲”影响，可以快速的执行
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}

}
