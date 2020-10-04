package main

import "fmt"

func main () {


	map01 := make(map[string]string)

	map01["name"] = "Go By Example!"
	map01["job"] = "临时工"
	map01["头衔"] = "高级零时工"

	fmt.Println(map01["name"])

	fmt.Println(map01)

	delete(map01, "头衔")

	fmt.Println(map01["头衔"])

	_, jods := map01["头衔"]

	fmt.Println(jods)

	fmt.Println(len(map01))


}
