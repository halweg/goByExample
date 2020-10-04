package main

import (
	"fmt"
)

func main() {
	s := []string{"h", "e", "l", "l", "0"}
	fmt.Println(s)

	var s2 = make([]string, 5)

	fmt.Println(s2)

	s2 = append(s2, "hello")

	fmt.Println(s2)
	fmt.Println("s2的长度是", len(s2))

	sc := make([]string, len(s2))
	copy(sc, s)
	fmt.Println(sc)

	l := sc[:4]
	fmt.Println("[:4]切片方法", l)

	fmt.Println("[2:]切片方法", sc[2:])

	fmt.Println("[2:5]的切片方法", sc[2:5])

	fmt.Println("")

	t := []string{"hell", "llo", "w", "ord"}
	fmt.Println(t)



}