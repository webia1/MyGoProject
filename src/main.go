package main

import (
	"fmt"
	_ "reflect"
)

func main() {

	a := []int{1, 3, 5, 7, 9}
	b := make([]int, 7)

	x := copy(b, a)

	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("x:", x)

	fmt.Println("Before Programm End")

}
