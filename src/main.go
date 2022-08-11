package main

import (
	"fmt"
	_ "reflect"
)

func main() {

	myMap1 := map[string]int{
		"foo": 7,
		"bar": 11,
		"baz": 2014,
	}

	delete(myMap1, "baz")

	fmt.Println(myMap1)

	fmt.Println("Before Programm End")

}
