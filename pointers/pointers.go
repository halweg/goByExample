package main

import "fmt"

func zeroval(i int) {
	i = 0
}

func zeroptr(i *int) {
	*i = 0
}

func main() {

	i := 4399
	zeroval(i)
	fmt.Println(i)

	zeroptr(&i)
	fmt.Println(i)
	fmt.Println("i address is ", &i)

}
