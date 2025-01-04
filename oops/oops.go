package oops

import (
	"encoding/json"
	"fmt"
)

func Oops() {
	type Name struct {
		First string `json:"first"`
		Last  string `json:"last"`
	}
	n := new(Name)
	n.First = "hello"
	n.Last = "world"
	bytes, err := json.Marshal(&n)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bytes))

	type Shape interface {
		Area() float64
		Perimeter() float64
	}
	type Circle struct {
	}
	type Rectangle struct {
	}
}
