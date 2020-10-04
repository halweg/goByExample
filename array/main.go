package main

import (
	"fmt"
)

func main() {

	var arr [4]int

	arr[0] = 101

	fmt.Println("数组的o号元素是", arr[0])
	fmt.Println("数组的1号元素是", arr[1])

	arr2 := [4]int{1,2,3,4}

	fmt.Println(arr2)

	var two [2][3]int
	for i := 0; i < 2; i++ {
		for j:=0; j< 3; j++ {
			two[i][j] = i+j
		}
	}
	fmt.Println(two)
}