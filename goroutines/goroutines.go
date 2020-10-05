package main

import (
	"math/rand"
	"fmt"
	"time"
)

func f(form string) {

	for i := 0; i <= 10; i++ {
		fmt.Println(form, ":", i)
		rand.Seed(int64(time.Now().UnixNano()) + int64(i))
		num := rand.Int31n(int32(15 + i))
		fmt.Println(num)
		time.Sleep(time.Duration(num)* time.Microsecond)
	}

}

func main() {

		f("同步调用")

		go f("狗汝汀")

		go func(s string) {
			fmt.Println(s)
		}("going")

		fmt.Scanln()
		fmt.Println("done")
}
