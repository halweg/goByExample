package main

import "fmt"

func main() {

		//若无缓冲buffer, 则必须有相应的channel接受者，不然编译无法通过
		bufChan := make(chan string, 2)

		bufChan<- "hello world"
		bufChan<- "this is "
		// 由于此通道是缓冲的，因此我们可以将这些值发送到通道中
		// 而不需要相应的并发接收。

		fmt.Println(<-bufChan)
		fmt.Println(<-bufChan)

}
