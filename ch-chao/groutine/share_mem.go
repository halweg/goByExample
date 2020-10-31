package main

import (
    "fmt"
    "sync"
)

func main() {
    c := 0
    var mlock sync.Mutex
    var wg sync.WaitGroup
    for i :=0; i < 100000; i++ {
        wg.Add(1)
        go func() {
            defer func() {
                mlock.Unlock()
            }()
            mlock.Lock()
            c++
            wg.Done()
        }()
    }
    wg.Wait()
    fmt.Println(c)
}