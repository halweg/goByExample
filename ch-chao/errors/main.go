package main

import "fmt"

func main() {
    recoverTest("1")
    fmt.Println("end")
}

func recoverTest(s string) {

    defer func() {
        if e := recover(); e != nil {
            fmt.Println(e)
        }
    }()

    panic("no args")
}
