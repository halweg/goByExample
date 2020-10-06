package main

import (
	"fmt"
	"strconv"
)

func main()  {
	//                          字符   0：自动推断进制   存储位数
	n1, _ := strconv.ParseInt("12", 0, 64)
	fmt.Println(n1)

	f1, _ := strconv.ParseFloat("12.3356533448", 64)
	fmt.Println(f1)

	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	// `ParseUint` 也是可用的。
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	// `Atoi` 是一个基础的 10 进制整型数转换函数。
	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	// 在输入错误时，解析函数会返回一个错误。
	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}