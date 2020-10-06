package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {

	state := make(map[int]int)

	var ops int64

	var mutex  sync.Mutex


	for i := 0; i<100; i++ {
		go func() {

			for {
				total := 0
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)

				runtime.Gosched()
			}

		}()
	}


	for i:=0; i < 10; i ++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)
				runtime.Gosched()
			}
		}()
	}

	time.Sleep(time.Second)

	loadedOps := atomic.LoadInt64(&ops)

	fmt.Println("操作了", loadedOps, "次")

	mutex.Lock()
	fmt.Println(state)
	mutex.Unlock()

}
