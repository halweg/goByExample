package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func fileExist(filePath string) bool {

	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}

	return true

}

func main() {

	//这里列出3种方式
	//1.细粒度写
	//2.细粒度写 ： f.Write 写byte ，f.WriteString 写string, 配合 f.Sync 写入磁盘
	//3.buffion方式 : bufio.NewWriter(f), 配合 w.Flush


	//1.粗粒度写
	if  hasf := fileExist("粗粒度写.txt"); hasf == true {
		err := os.Remove("粗粒度写.txt")
		if err != nil {
			panic(err)
		}
	}
	err := ioutil.WriteFile("粗粒度写.txt", []byte("粗粒度写一段东西"), 644)
	if err != nil {
		panic(err)
	}




	//2.细粒度写
	if  hasf := fileExist("细粒度写.txt"); hasf == true {
		err := os.Remove("细粒度写.txt")
			if err != nil {
				panic(err)
			}
	}
	f1, e := os.Create("细粒度写.txt")
	if e != nil {
		panic(e)
	}
	defer func() {
		f1.Close()
		fmt.Println("关闭了句柄 f1")
	}()
	wcnt, e := f1.Write([]byte("使用细粒度一次性写入文件 \" 细粒度写.txt \"里..."))
	if e != nil {
		panic(e)
	}
	wcnt2, _ := f1.WriteString("\n使用 WriteString 再写一段话 ")
	fmt.Println("向文件[细粒度写.txt]中写入了一段话，共 ", wcnt+wcnt2, "个 byte")

	//3.带缓冲的写法
	if  hasf := fileExist("bufio方式写文件.txt"); hasf == true {
		err := os.Remove("bufio方式写文件.txt")
		if err != nil {
			panic(err)
		}
	}

	f3, err := os.Create("bufio方式写文件.txt")
	if err != nil {
		panic(err)
	}
	defer f3.Close()

	bf := bufio.NewWriter(f3)
	bf.WriteString("bufio方式写入文件")

	bf.Flush()
}
