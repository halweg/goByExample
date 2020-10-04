package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Print("write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Today is weekend")
	default:
		fmt.Println("Today is not weekend")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("现在是早上")
	default:
		fmt.Println("现在是下午")
	}

	varType := func(i interface{}) {
		switch i.(type) {
		case bool:
			fmt.Println("变量类型为bool")
		case byte:
			fmt.Println("变量类型为byte")
		case int:
			fmt.Println("变量类型为int")
		case string:
			fmt.Println("变量类型为string")
		}
	}

	varType(1)
	varType("what")
	varType(true)
	var bt byte = 1
	varType(bt)
}
