package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	type person struct {
		Fullname string `json: "fullname"`
		Age      int    `json: "age"`
	}

	p := person{}

	err := json.Unmarshal([]byte(`{"fullname": "Michael Jackson", "age": 55}`), &p)

	fmt.Println(err) // <nil>
	fmt.Println(p)   // {Michael Jackson 55}

	fmt.Println("Before Programm End")

}
