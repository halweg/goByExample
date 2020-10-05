package main

import (
	"fmt"
	"time"
)

func main()  {

 	//select 的超时 case

	c1 := make(chan string, 1)

	go func() {
		time.Sleep(time.Second * 3)
		c1 <- "往c1里放东西"
	}()

	select {
	case r1 := <-c1 :
		fmt.Println(r1)
	case <-time.After(time.Second * 2 ) :
		fmt.Println("任务超时了。。。")
	}

}
