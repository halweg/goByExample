package main

import (
	"fmt"
	"path"
	"path/filepath"
	"strings"
)

func main() {

	//拼接目录的方法，代替手动
	p := path.Join("tmp", "rec", "bin.ext")
	fmt.Println(p)

	// `Dir` 和 `Base` 可以被用于分割路径中的目录和文件, 就是分割最后一级和前面的多级
	fmt.Println("Dir(p):", filepath.Dir(p))
	fmt.Println("Base(p):", filepath.Base(p))

	fmt.Println(filepath.IsAbs("dir/relative/path"))
	fmt.Println(filepath.IsAbs("p:\\ProgramData"))

	fmt.Println(path.Ext("config.ini"))

	ext := path.Ext("config.ini")  //.ini
	fmt.Println(strings.TrimSuffix("config.ini", ext )) //config

								//bastpath 和 targpath 之间的相对路径
	n, _ := filepath.Rel("ab/c", "ab/c/d/ff")
	fmt.Println(n)

}
