package main

import (
	"fmt"
	_ "reflect"
)

func main() {

	x := make([]int, 0, 10)
	x = append(x, 3, 5, 7, 9)
	b := x[:2:2]
	c := x[2:4:4]

	fmt.Println("x:", x)
	fmt.Println("b:", b)
	fmt.Println("c:", c)

	b = append(b, 20, 30, 40)
	x = append(x, 11)
	c = append(c, 13)

	fmt.Println("x:", x)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
	fmt.Println("Before Programm End")

}
