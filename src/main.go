package main

import (
	"fmt"
	_ "reflect"
)

func main() {

	s1 := Singer{
		name: "Michael Jackson",
		age:  55,
	}

	PrintSinger(s1)

	fmt.Println("Before Programm End")

}

func PrintSinger(singer Singer) {
	fmt.Println(singer)
}

type Singer struct {
	name string
	age  int
}
