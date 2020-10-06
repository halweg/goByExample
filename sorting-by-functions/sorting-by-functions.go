package main

import (
	"fmt"
	"sort"
)

func sorting() {

	nums := []int{70,0,1,5,55,0,2,12}

	sort.Ints(nums)

	fmt.Println(nums)

	strs := []string{"a", "b", "ss", "a" , "d", "c", "ea"}

	fmt.Println(sort.StringsAreSorted(strs))

	sort.Strings(strs)

	fmt.Println(strs)

	fmt.Println(sort.StringsAreSorted(strs))

}

type byteLength []string

func (s byteLength) Swap(i, j int)  {
	s[i], s[j] = s[j], s[i]
}

func (s byteLength) Less(i,j int) bool {
	return len(s[i]) < len(s[j])
}

func (s byteLength) Len() int  {
	return len(s)
}

func main()  {

	strs := []string{"hwll", "s", "ss", "ssss"}

	//        实现了 sort 接口， 按照 字符串长度排序
	sort.Sort(byteLength(strs))

	fmt.Println(strs)
}
