package main

import (
    "fmt"
    "time"
)

func IsCanceled(ch chan int) bool{
    select {
    case <-ch:
        return true
    default:
        return false
    }
}

func Cancel1(ch chan int)  {
    ch<-1
}
func Cancel2(ch chan int)  {
    close(ch)
}

func main() {
    ch := make(chan int)
    for i:=0; i<5; i++ {
        go func(i int, ch chan int) {

            for {
                if IsCanceled(ch) {
                    fmt.Println("channel ", i, "has been close")
                    break
                }

                time.Sleep(time.Millisecond * 100)
            }

        }(i, ch)
    }
    Cancel2(ch)
    time.Sleep(time.Second * 1)
}