package main

import (
	"fmt"
)

func main() {

	var s Singer = Singers{
		Name: "Michael Jackson",
	}

	fmt.Println(s)
	s.Sing()
	s.Talk()

	fmt.Println("Before Programm End")

}

type Singer interface {
	Talk()
	Sing()
}

type Singers struct {
	Name string
}

func (s Singers) Talk() {
	fmt.Println("Talking")
}

func (s Singers) Sing() {
	fmt.Println("Singing")
}
