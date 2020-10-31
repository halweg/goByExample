package main

import (
    "fmt"
    "time"
)

func services() string{
    time.Sleep(time.Millisecond * 200)
    return "Done"
}

func ASyncService() chan string{
    retCh := make(chan string, 1)
    go func() {
        fmt.Println("async start")
        r := services()
        retCh<-r
        fmt.Println("async done")
    }()

    return retCh
}

func main()  {
    select {
    case ret := <-ASyncService():
        fmt.Println(ret)
    case <-time.After(time.Millisecond * 100):
        fmt.Println("time out")
    }
}
