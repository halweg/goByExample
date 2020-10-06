package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Print(rand.Intn(100), ",")
	fmt.Print(rand.Intn(100))
	fmt.Println()

	//0~1之间的随机浮点数
	fmt.Println(rand.Float64())

	//利用 0~1 之间的随机浮点数可以获取其他范围的随机数
	fmt.Println(rand.Float64() * 10) //
	fmt.Println(rand.Float64() * 10) //
	fmt.Println(rand.Float64() * 10) //
	fmt.Println(rand.Float64() * 10) //
	fmt.Println(rand.Float64() * 5 + 2) //


	//0~100之间
	/*fmt.Println(int(rand.Float64() * 100) )
	fmt.Println(int(rand.Float64() * 100) )
	fmt.Println(int(rand.Float64() * 100) )
	fmt.Println(int(rand.Float64() * 100) )*/

	//默认的随机种子确定的，每次运行期产生的随机数都是一样的

	seed := time.Now().UnixNano()
	r1 := rand.New(rand.NewSource(seed))

	fmt.Println(r1.Intn(10))
	fmt.Println(r1.Intn(10))
	fmt.Println(r1.Intn(10))
	fmt.Println(r1.Intn(100))
	fmt.Println(r1.Intn(100))

	//若使随机数种子不变，每次编译运行出的随机数序列是一样的
	r2 := rand.New(rand.NewSource(2))
	fmt.Println(r2.Intn(100))
	fmt.Println(r2.Intn(100))
	fmt.Println(r2.Intn(100))
	fmt.Println(r2.Intn(100))
	fmt.Println(r2.Intn(100))
}
