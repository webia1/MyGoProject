package main

import (
	"fmt"
	"reflect"
	_ "reflect"
)

func main() {

	var x []int
	var y = []int{}

	fmt.Println(x, len(x)) // [] 0
	fmt.Println(y, len(y)) // [] 0

	fmt.Println(x == nil) // true
	fmt.Println(y == nil) // false

	fmt.Println(reflect.TypeOf(x))                      // []int
	fmt.Println(reflect.TypeOf(y))                      // []int
	fmt.Println(reflect.TypeOf(y) == reflect.TypeOf(x)) // true

	fmt.Println(reflect.ValueOf(x).Kind()) // slice
	fmt.Println(reflect.ValueOf(y).Kind()) // slice

	var v []int
	z := make([]int, 0, 10)

	fmt.Println(v, len(v), cap(v))
	fmt.Println(z, len(z), cap(z))

	v = append(v, 3, 5, 7)
	z = append(x, 3, 5, 7)

	fmt.Println(v, len(v), cap(v))
	fmt.Println(z, len(z), cap(z))

	fmt.Println("Before Programm End")
}
