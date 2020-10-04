package main

import "fmt"

func main() {

	nums := []int{2,3,4}
	sum := 0

	for  _, num := range nums {
		sum += num
	}

	fmt.Println(sum)

	for i, num := range nums {
		fmt.Println(nums[i]==num)
	}

	map02 := make(map[string]string)
	map02["name"] = "程咬金"
	map02["job"]  = "士兵"
	map02["age"] = "66"

	var map03 map[string]string
	fmt.Println(map03)

	map04 := map[string]string{"name":"老六", "age":"74"}

	for k, v := range map04 {
		fmt.Printf("%s:%s,", k, v)
	}
	fmt.Println("")
	for k := range map02 {
		fmt.Println(k)
	}

	for i, c := range "This is a simple words 中 人 并 " {
		fmt.Println(i, c)
	}


}
