package main

import (
	"fmt"
	"webia1/MyGoProject/src/treeexample"
)

func main() {

	var it *treeexample.IntTree

	var vals = []int{2, 1, 3, 3, 3, 4, 2, 7, 1, 5, 3}

	for _, v := range vals {
		it = it.Insert(v)
	}

	fmt.Println(it.Has(7))

	fmt.Println("Before Programm End")

}
