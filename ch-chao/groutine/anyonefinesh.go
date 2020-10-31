package main

import (
    "fmt"
    "runtime"
    "time"
)

func Task(id int) string  {
    time.Sleep(time.Second * 1)
    return fmt.Sprintf("task %d is fin\n", id)
}

func ToResponse() string {
    chanflag := make(chan string, 10)

    for i := 0; i < 10; i++ {
        go func(i int) {
            ret := Task(i)
            chanflag <- ret
        }(i)
    }

    return <-chanflag
}


func main() {
    fmt.Println("before ", runtime.NumGoroutine())
    fmt.Println(ToResponse())
    time.Sleep(time.Second * 2)
    fmt.Println("after ", runtime.NumGoroutine())
}
