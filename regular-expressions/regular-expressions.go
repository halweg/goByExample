package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {

	//bool
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	r, _ := regexp.Compile("p([a-z]+)ch")
	fmt.Println(r.MatchString("peach"))

	//返回/查找字符串
	fmt.Println(r.FindString("peach puch"))
	fmt.Println(r.FindAllString("peach  puch  puch puchpuchpuchpuch", -1))

	//返回第一个符合的Index
	fmt.Println(r.FindStringIndex("peach  puch  puch puchpuchpuchpuch"))

	//返回完匹配和局部匹配的字符串, 此处会返回 peach 和 ea
	fmt.Println(r.FindStringSubmatch("peach"))

	// 类似的，这个会返回完全匹配和局部匹配的索引位置。
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	// 这个函数提供一个正整数来限制匹配次数。
	fmt.Println(r.FindAllString("peach punch pinch", 2))

	// 上面的例子中，我们使用了字符串作为参数，并使用了如 `MatchString` 这样的方法。
	// 我们也可以提供 `[]byte` 参数并将 `String` 从函数命中去掉。
	fmt.Println(r.Match([]byte("peach")))

	// 创建正则表达式常量时，可以使用 `Compile` 的变体 `MustCompile` 。
	// 因为 `Compile` 返回两个值，不能用于常量。
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)

	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	// `Func` 变量允许传递匹配内容到一个给定的函数中，
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))

}
