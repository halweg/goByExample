package main

import (
	"fmt"
	"time"
)

func main() {

	p := fmt.Println

	t := time.Now()

	p(t.Format(time.RFC3339))

	ft, e := time.Parse(time.RFC3339, "2020-10-06T14:50:09+00:00")

	if e != nil {
		panic(e)
	}
	fmt.Println(ft)

}
