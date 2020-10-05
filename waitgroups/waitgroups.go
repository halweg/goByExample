package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(workerId int, wg *sync.WaitGroup) {

	fmt.Println("worker ", workerId, " is starting...")
	time.Sleep(time.Second)
	fmt.Println("worker ", workerId, " is ending")

	wg.Done()

}

func main() {

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()
}
