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
	fmt.Println("\n\n\n-------------------------- -------------------------")

	someInts := []int{1, 0, 2, 3, 4}
	for _, v := range someInts {
		divideTenBy(v)
	}

	fmt.Println("-------------------------- -------------------------\n\n\n")

	fmt.Println("Debugger")
}
