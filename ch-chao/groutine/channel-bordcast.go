package main

import (
    "fmt"
    "sync"
)

func DataProduce(ch chan int, wg *sync.WaitGroup) {
    go func() {
        for i:=0; i<10; i++ {
            ch <- i
        }
        close(ch)
        wg.Done()
    }()
}

func DataReceive(cha chan int, wg *sync.WaitGroup)  {
    go func() {
        for {
            if d, ok := <-cha; !ok {
                fmt.Println("channel has close")
                break
            } else {
                fmt.Println(d)
            }
        }
        wg.Done()

    }()
}

func main() {
    var wg sync.WaitGroup
    ch := make(chan int)

    wg.Add(1)
    DataProduce(ch, &wg)
    wg.Add(1)
    DataReceive(ch, &wg)
   // time.Sleep(time.Second * 1)
    wg.Wait()
}
