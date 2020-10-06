package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

func main() {

	var incrementId uint64

	for i := 0; i < 50; i++ {

		go func() {

			for {
				atomic.AddUint64(&incrementId, 1)

				//把机会交给其他 狗汝汀。防止自己被惩罚/饿死
				runtime.Gosched()
			}

		}()

	}

	time.Sleep(time.Second * 1)

	loadedId := atomic.LoadUint64(&incrementId)

	fmt.Println(loadedId)

}
