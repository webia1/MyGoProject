package main

import (
	"fmt"
	_ "reflect"
)

func main() {

	var x = []int{2, 3, 5, 7, 9}
	var a = x[:3]
	var b = x[2:]
	var c = x[1:4]
	var d = x[:]
	var e = x

	fmt.Println(x, a, b, c, d, e)
	fmt.Println("Before Programm End")

}
