package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	fmt.Println("2 groutine is starting...")

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		time.Sleep(time.Second * 2)
		wg.Done()
		fmt.Println("groutine 1 is end")
	}()

	wg.Add(1)
	go func() {
		time.Sleep(time.Second * 1)
		wg.Done()
		fmt.Println("groutine 2 is end")
	}()

	wg.Wait()
	fmt.Println("all groutine is end")
}

