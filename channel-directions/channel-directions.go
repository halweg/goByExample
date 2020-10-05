package main

import "fmt"

func ping (pings chan<- string, msg string ) {
	pings <- msg
}

//单向通道，只进不出/只出不进，多用于形参， 直接声明成单向通道的情况，暂时没有想到用途
func pong(pongs chan<- string,  pings <-chan string) {
	msg := <-pings
	pongs<- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)


	ping(pings, "hello")

	pong(pongs, pings)

	fmt.Println(<-pongs)

}
