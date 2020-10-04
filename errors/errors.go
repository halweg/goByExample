package main

import (
	"errors"
	"fmt"
)

func f1(i int) (int, error) {

	if (i == 42) {
		err := errors.New("can`t work with it")
		return   -1,  err
	}

	return i+3,nil
}

type argError struct {
	arg int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(i int) (int, error) {

	if (i == 42) {
		return -1, &argError{
			arg:  i,
			prob: "can`t work with it",
		}
	}
	return i+3, nil
}

func main() {

	i, err := f1(42)

	if err != nil {
		fmt.Println(err)
	}

	i , err2 := f2(42)

	if err != nil {
		fmt.Println(err2)
	}

	fmt.Println(i)

	_, e := f2(42)

	ae, ok := e.(*argError)

	fmt.Println(ae)
	fmt.Println(ok)
	
	if ae, ok = e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}

}
