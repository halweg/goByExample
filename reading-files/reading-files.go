package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {

	if e != nil {
		panic(e)
	}

}

func main() {

	dat, err := ioutil.ReadFile("./go.mod")
	check(err)
	fmt.Printf("%T\n", dat)
	//fmt.Printf("%s\n", dat)

	f, err := os.Open("go.mod")
	check(err)

	buf := make([]byte, 20)
	hasReadLength, err := f.Read(buf)
	hasReadLength, err = f.Read(buf)
	hasReadLength, err = f.Read(buf)
	hasReadLength, err = f.Read(buf)
	check(err)
	fmt.Printf("%d,%s\n", hasReadLength, buf)

	//操作文件指针
	          //偏移量    //起始位置
	f.Seek(0, 0)
	hasReadLength, err = f.Read(buf)
	fmt.Printf("%d,%s\n", hasReadLength, buf)

	r4 := bufio.NewReader(f)

	f.Seek(0, 0)
	b4, err := r4.Peek(20)
	fmt.Println("------带缓冲的读写------")
	fmt.Printf("%s", b4)

	defer func() {
		f.Close()
	}()
}
