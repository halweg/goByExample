package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {

	fmt.Println("🐕汝町开始工作.......")
	time.Sleep(time.Second * 10)
	fmt.Println("🐕🐕🐕ending........")

	done <- true

}

func main() {

	done := make(chan bool,1)

	go worker(done)

	fmt.Println("在此等候 狗汝汀")

	// 程序将在接收到通道中 worker 发出的通知前一直阻塞。
	// 利用通道的默认阻塞 来同步
	<-done
}
