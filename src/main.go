package main

import "fmt"

func main() {

	var x = [...]int{1, 2, 3}
	x[0] = 10
	fmt.Println(x)
	fmt.Println(len(x))

}
