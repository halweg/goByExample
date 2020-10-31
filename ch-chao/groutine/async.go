package main

import (
    "fmt"
    "time"
)

func service()  string {
    time.Sleep(time.Microsecond * 100)
    fmt.Println("service done")
    return "Done"
}

func OtherTask()  {
    fmt.Println("This is other task!")
}

func AsyncService() chan string {
    chant := make(chan string)

    go func() {
        ret := service()
        chant <- ret
    }()
    return chant
}

func main()  {
    r := AsyncService()
    OtherTask()
    fmt.Println(<-r)
}
