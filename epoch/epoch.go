package main

import (
	"fmt"
	"time"
)

func main()  {

	now := time.Now()

	sec := now.Unix()

	naoc := now.UnixNano()

	millisec := naoc / 1000000

	fmt.Println(sec)
	fmt.Println(naoc)
	fmt.Println(millisec)

	fmt.Println(time.Unix(sec, 0))
	fmt.Println(time.Unix(0, naoc))

}
