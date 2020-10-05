package main

import "fmt"

func main() {

	jobs := make(chan string)
	done := make(chan bool)

	go func() {
		for {
			j, hasMore := <-jobs
			if hasMore {
				fmt.Println("消费了 ： "+j)
			} else {
				fmt.Println("没有更多的job了")
				done<-true
				return
			}
		}
	}()

	for i:=0; i <= 3; i ++ {
		jobs<-"一个任务"
		fmt.Println("jobs 被安排了一个任务")
	}
	close(jobs)

	<-done
}
