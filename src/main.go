package main

import (
	"fmt"
	"strconv"
	"webia1/MyGoProject/src/calculator"
)

func main() {

	someExpressions := [][]string{
		{"4", "+", "2"},
		{"4", "-", "2"},
		{"4", "*", "2"},
		{"4", "/", "2"},
	}

	for _, se := range someExpressions {
		a, err := strconv.ParseFloat(se[0], 64)
		if err != nil {
			continue
		}

		op := se[1]
		opKind, ok := calculator.OpMap[op]
		if !ok {
			fmt.Println("Op not within OpMap")
		}

		b, err := strconv.ParseFloat(se[2], 64)
		if err != nil {
			continue
		}

		result := opKind(a, b)
		fmt.Println("Result: ", result)

	}

	fmt.Println("Before Programm End")
}
