package main

import (
	"fmt"
)

func main() {
	panicExample1()

	buf := []int{1, 2}
	fmt.Println(panicExample2(buf, 2))
}

func panicExample1() {
	defer func() {
		recover()

		// err := recover();
		// if err != nil {
		// 	fmt.Println(err)
		// }
	}()

	panic("rrrrrrrrrr")
}


// ref: https://stackoverflow.com/questions/25025467/catching-panics-in-golang
func panicExample2(buf []int, i int) (x int, err interface{}) {
	defer func() {
		err = recover()
	}()

	x = buf[i]
	return
}
