package main

import "fmt"

type person struct {
	name string
	age int
	job string
}

func main()  {
	p1 := person{
		name: "halweg",
		age:  18,
		job:  "果农",
	}

	fmt.Println(p1.name)
	fmt.Println(p1)

	p2 := person{
		name: "halweg2",
		job:  "jobLess",
	}
	fmt.Println(p2)

	p3 := &person{
		name: "halweg3",
		age:  5,
		job:  "homeLess and jobLess",
	}
	fmt.Println(p3)
	fmt.Println(p3.name)

	p4 := new(person)

	fmt.Println(p4)

	p4.name = "hhh"

	fmt.Println(p4.name)
}
