package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {

	//os.RemoveAll("临时目录")
	err := os.Mkdir("临时目录", 755)
	if err != nil {
		//panic(err)
	}

	createEmptyDir := func(name string) {
		d := []byte("")
		err := ioutil.WriteFile(name, d, 0644)
		if err != nil {
			panic(err)
		}
	}

	createEmptyDir("临时目录/dir.exa")

	err = os.MkdirAll("临时目录/level1/tmp/dir/子目录", 755)
	if err != nil {
		panic(err)
	}

	dirs , er	:= ioutil.ReadDir("临时目录")
	if er != nil {
		panic(er)
	}

	for  _, dir := range dirs {
		fmt.Println(dir.Name(), dir.IsDir())
		if dir.IsDir() {

		}
	}
	c, _ := ioutil.ReadDir(".")
	for _, cc := range c{
		fmt.Println(cc.Name())
	}
	//defer os.RemoveAll("临时目录")

	// `cd` 回到最开始的地方。
	err = os.Chdir("../../..")


	// 当然，我们也可以遍历一个目录及其所有子目录。
	// `Walk` 接受一个路径和回调函数，用于处理访问到的每个目录和文件。
	fmt.Println("Visiting subdir")
	err = filepath.Walk("subdir", visit)

	// 创建临时文件最简单的方法是调用 `ioutil.TempFile` 函数。
	// 它会创建并打开文件，我们可以对文件进行读写。
	// 函数的第一个参数传 `""`，`ioutil.TempFile` 会在操作系统的默认位置下创建该文件。
	f, err := ioutil.TempFile("", "sample")
	check(err)

	// 打印临时文件的名称。
	// 文件名以 `ioutil.TempFile` 函数的第二个参数作为前缀，
	// 剩余的部分会自动生成，以确保并发调用时，生成不重复的文件名。
	// 在类 Unix 操作系统下，临时目录一般是 `/tmp`。
	fmt.Println("Temp file name:", f.Name())

	// defer 删除该文件。
	// 尽管操作系统会自动在某个时间清理临时文件，但手动清理是一个好习惯。
	defer os.Remove(f.Name())

	// 我们可以向文件写入一些数据。
	_, err = f.Write([]byte{1, 2, 3, 4})
	check(err)

	// 如果需要写入多个临时文件，最好是为其创建一个临时 *目录*。
	// `ioutil.TempDir` 的参数与 `TempFile` 相同，
	// 但是它返回的是一个 *目录名*，而不是一个打开的文件。
	dname, err := ioutil.TempDir("", "sampledir")
	fmt.Println("Temp dir name:", dname)

	defer os.RemoveAll(dname)

	// 现在，我们可以通过拼接临时目录和临时文件合成完整的临时文件路径，并写入数据。
	fname := filepath.Join(dname, "file1")
	err = ioutil.WriteFile(fname, []byte{1, 2}, 0666)
	check(err)
}

func visit(p string, info os.FileInfo, err error) error {
	// `cd` 回到最开始的地方。
	err = os.Chdir("../../..")

	// 当然，我们也可以遍历一个目录及其所有子目录。
	// `Walk` 接受一个路径和回调函数，用于处理访问到的每个目录和文件。
	fmt.Println("Visiting subdir")
	err = filepath.Walk("subdir", visit)
	return nil
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

