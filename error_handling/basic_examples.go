package main

import (
	"errors"
	"fmt"
)

func main() {
	example1()
	example2()
}

func example1() {
	err := doSomething()

	if (err != nil) {
		fmt.Println("err: ", err)
		return
	}
}

func example2() {
	slice := []int{1, 2}
	element, err := getKthElement(slice, 2)

	if (err != nil) {
		fmt.Println("err: ", err)
		return
	}

	fmt.Println("element: ", element)
}

func doSomething() error {
    if true {
        return errors.New("error la")
    }

    return nil
}

func getKthElement(slice []int, k int) (element int, err error) {
	if (k < 0 || k >= len(slice)) {
		err = errors.New("index out of range!")
		return
	}

	element = slice[k]
	return
}
