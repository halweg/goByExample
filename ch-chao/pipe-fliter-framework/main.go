package main

import "fmt"

func main() {
    slipt := NewSplitFilter(",")
    toInt := NewToIntFilter()
    sum := NewSumFilter()

    pipeline := NewStraightPipeline("pip", slipt, toInt, sum)

    ret, e := pipeline.Process("1,2,3,4")
    if e != nil {
        fmt.Println(e)
    } else {
        fmt.Println("the ret is ", ret)
    }

}
