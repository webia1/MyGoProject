package main

import (
	"fmt"
	"webia1/MyGoProject/src/counterexample"
)

func main() {

	var c counterexample.Counter
	fmt.Println(c.Log())
	c.Inc()
	fmt.Println(c.Log())

	fmt.Println("Before Programm End")

}
