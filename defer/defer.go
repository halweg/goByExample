package main

import (
	"fmt"
	"os"
)

func main()  {

	f := createFile("./tmp.txt")
	writeFile(f)

	defer closeFile(f)

}

func createFile(fileName string) *os.File {

	f, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	return f
}

func writeFile(f *os.File) {

	fmt.Println("writing")
	fmt.Fprintln(f, "data")

}

func closeFile(f *os.File) {
	fmt.Println("closing")
	f.Close()
}


