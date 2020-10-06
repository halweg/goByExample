package main

import (
	"crypto/sha1"
	"fmt"
)

func main()  {

	s := "sha1 this string 这是一个需要被哈希的字符串"

	sh := sha1.New()

	sh.Write([]byte(s))

	//nil 是需要追加的字符，不追加就置空
	bs := sh.Sum(nil)

	//sha1 一般以 16 进制表示，在 git 中极为常见
	fmt.Printf("%x",  bs)
}
