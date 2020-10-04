package main

import "fmt"

func vals() (int, bool) {
	return 8, false
}

func main () {
	res1, res2 := vals()

	fmt.Println(res1)

	fmt.Println(res2)

	_, res2 = vals()

	fmt.Println(res2)
}
