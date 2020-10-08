package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// 对 `os.Stdin` 使用一个带缓冲的 scanner，
	// 让我们可以直接使用方便的 `Scan` 方法来直接读取一行，
	// 每次调用该方法可以让 scanner 读取下一行。
	sanner := bufio.NewScanner(os.Stdin)

	for sanner.Scan() {
		s := "你输入了"+sanner.Text()
		fmt.Println(s)
		if sanner.Text() == "exit" {
			fmt.Println("Bye!")
			break
		}
	}

}
