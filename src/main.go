package main

import (
	"fmt"
)

func divideTenBy(i int) {

	defer func() {
		if v := recover(); v != nil {
			fmt.Println(v)
		}
	}()

	fmt.Println(10 / i)
}

func main() {

	someInts := []int{1, 2, 0, 3}
	for _, v := range someInts {
		divideTenBy(v)
	}

	fmt.Println("Debugger")
}
