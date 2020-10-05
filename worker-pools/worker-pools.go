package main

import (
	"fmt"
	"strconv"
	"time"
)

func worker(workerId int, jobs <-chan int, response chan<- string) {

	for j := range jobs {
		time.Sleep(time.Second * 10)
		//fmt.Println("jobs ：", j, "被遍历,此时jobs长度为", len(jobs), "response 长度为", len(response))
		response<-  "response结果 ： worker id ："+ strconv.Itoa(workerId) + "消费了" + strconv.Itoa(j)
	}

}

func main() {

	jobs := make(chan int, 100)
	response := make(chan string, 100)

	for i :=0; i<9; i++ {
		//fmt.Println("worker： ", i, "被投放")
		go worker(i, jobs, response)   //由于jobs 是空的，此时 worker 中的 range 会阻塞住
	}

	for i := 1; i<=9; i++ {
		jobs<- i
		//fmt.Println("job : ", i , "被投放")
	}
	//fmt.Println("jobs 要被关闭了......")
	close(jobs)
	//fmt.Println("jobs 被关闭了......")

	//time.Sleep(time.Second * 10)
	for i :=1; i<=9; i++ {
		fmt.Println(<-response)
	}
}
