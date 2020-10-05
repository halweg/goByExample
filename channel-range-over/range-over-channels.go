package main

import "fmt"

func main()  {

	queue := make(chan  string, 2)

	queue<- "第一个任务"
	queue<- "第二个任务"
	close(queue)

	for q := range queue {
		fmt.Println(q)
	}

}
