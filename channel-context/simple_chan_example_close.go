package main

import (
	"fmt"
	"time"
)

func main() {

	stop := make(chan bool, 1)

	fmt.Println("监控开启中。。。")
	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("groutine is stopped!")
				return
			default:
				time.Sleep(time.Second *1)
				fmt.Println("groutine is working...")
			}
		}

	}()

	time.Sleep(10 * time.Second)

	fmt.Println("可以了，通知监控停止")

	stop <- true

	time.Sleep(5 * time.Second)

}
