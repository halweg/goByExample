package main

import "fmt"

type point struct {
	x, y int
}

func main() {

	p := point{9,10}
	//以 value 形式
	fmt.Printf("%v\n", p)

	//以 k:v 形式
	fmt.Printf("%+v\n", p)

	//相应值的Go语法表示
	fmt.Printf("%#v\n", p)

	//类型
	fmt.Printf("%T\n", p)

	//布尔占位符
	fmt.Printf("%t\n", true)

	//整型形式
	fmt.Printf("%d\n", p.x)

	//二进制形式
	fmt.Printf("%b\n", p)

	//输出给定整数的字符
	fmt.Printf("%c\n", 52)

	//16进制形式
	fmt.Printf("%x\n", 26)

	//浮点形式
	fmt.Printf("%f\n", 0.9)

	// `%e` 和 `%E` 科学计数法
	fmt.Printf("%e\n", 100000.0)
	fmt.Printf("%E\n", 100000.0)

	//字符串形式
	fmt.Printf("%s\n", "\"string\"")

	// 和上面的整型数一样，`%x` 输出使用 base-16 编码的字符串，每个字节使用 2 个字符表示。
	fmt.Printf("%x\n", "Hex this")

	//输出一个指针的值
	fmt.Printf("%p\n", &p)

	//控制宽度
	fmt.Printf("|%6d|%6d|\n", 12, 345)

	//控制宽度和精度
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)

	//左对其
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)


	s := fmt.Sprintf("[[%s]]", "this")
	fmt.Printf(s)

}
