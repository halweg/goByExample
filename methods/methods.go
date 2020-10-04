package main

import "fmt"

type Employee struct {
	name string
	birthDay string
	job string
}

func (e *Employee) getName() string {
	return e.name
}

func (e *Employee) setJob(job string) {
	e.job = job
}

func (e Employee) setBirthDay(b string) {
	e.birthDay = b
}

func main() {

	e := new(Employee)

	e.setBirthDay("2020-12-20")

	fmt.Println(e.birthDay)

	e.setJob("Job Less")

	fmt.Println(e.job)

	e2 := Employee{
		name:     "halweg",
		birthDay: "",
		job:      "",
	}

	fmt.Println(e2.birthDay)

}


