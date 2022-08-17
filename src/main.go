package main

import (
	"fmt"
	"sort"
)

func main() {

	type Person struct {
		Fullname string
		Age      int
	}

	people := []Person{
		{"Michael Jackson", 55},
		{"George Michael", 56},
	}

	sort.Slice(people, func(i int, j int) bool {
		return people[i].Fullname < people[j].Fullname
	})
	fmt.Println("People: ", people)

	sort.Slice(people, func(i int, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println("People: ", people)

	fmt.Println("Before Programm End")
}
