package main

import (
	"fmt"
	"strings"
)

var p = fmt.Println

func main()  {

	p("Contains: ", strings.Contains("testsasd", "d"))

	p("Index: ", strings.Index("TESTESsA", "s"))

	p("Join: ", strings.Join([]string{"a", "b"}, "-"))

	//是否以 指定字符串开头
	p("haspre :", strings.HasPrefix("halsdgsdg", "ha"))

	//是否以 指定字符串结尾
	p("hasSuf: ", strings.HasSuffix("hasdfasdfddd", "ddd"))

	p("Index: ", string([]rune("Hel你o")[3]))

	//出现的次数
	p("Count: ", strings.Count("Helloe", "e"))
}
