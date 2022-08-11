package main

import (
	"fmt"
	_ "reflect"
)

func main() {

	myMap1 := make(map[int][]string, 10)
	myMap2 := make(map[string]int, 10)

	myMap1[13] = []string{"Hi", "there"}
	myMap2["foo"] = 4716

	fmt.Println(myMap1, myMap2)

	fmt.Println("Before Programm End")

}
