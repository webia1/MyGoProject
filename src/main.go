package main

import (
	"fmt"
	_ "math/rand"
	_ "reflect"
	_ "time"
)

func main() {

	initialValues := [][]int{{1, 2, 3}, {2, 4, 6}}

myLabelLevel0:
	for i, level0Val := range initialValues {
		for _, level1Val := range level0Val {
			if level1Val == 3 {
				fmt.Println("Iteration", i+1, "Level 1: Skipped value", level1Val)
				continue myLabelLevel0
			}
		}
		fmt.Println("Iteration", i+1, "Level 1: No values skipped")
	}

	fmt.Println("Before Programm End")

}
