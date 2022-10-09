package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func divideTenBy(i int) (int, error) {

	defer func() {
		if v := recover(); v != nil {
			fmt.Println(v)

		}
	}()

	return (10 / i), nil
}

func main() {

	type stackTracer interface {
		StackTrace() errors.StackTrace
	}

	fmt.Println("\n\n\n--------------------------")
	someInts := []int{1, 0, 2, 3, 4}
	for _, v := range someInts {
		_, err := divideTenBy(v)

		if err, ok := err.(stackTracer); ok {
			for _, f := range err.StackTrace() {
				fmt.Printf("%+s:%d\n", f, f)
			}
		}

	}
	fmt.Println("--------------------------\n\n\n")

	fmt.Println("Debugger")
}
