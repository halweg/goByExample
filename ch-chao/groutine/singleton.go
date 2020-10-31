package main

import (
    "fmt"
    "sync"
    "unsafe"
)

type singleton struct {

}

var sing *singleton
var once sync.Once

func GetSingleton()  *singleton {
    once.Do(func() {
        fmt.Println("just run once")
        sing = new(singleton)
    })
    return sing
}

func main() {
    var wg sync.WaitGroup

    for i:=0; i<10; i++ {
        wg.Add(1)
        go func() {
            obj := GetSingleton()
            fmt.Printf( "%x\n", unsafe.Pointer(obj))
            wg.Done()
        }()
    }

    wg.Wait()
}
