package main

import "fmt"

func main() {

	var x []int
	fmt.Println(x, len(x), cap(x))
	x = append(x, 10)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 11)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 12)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 13)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 14)
	fmt.Println(x, len(x), cap(x))

}
