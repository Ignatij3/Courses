package main

import (
	"fmt"
)

type Ordered interface {
	Before(b Ordered) bool
}

type Integer int

func (a Integer) Before(b Ordered) bool {
	return a < b.(Integer)
}

type Float float64

func (a Float) Before(b Ordered) bool {
	return a < b.(Float)
}

func main() {
	d := []Ordered{Integer(310), Float(310.5)}
	for _, x := range d {
		if v, ok := x.(Integer); ok {
			fmt.Println(v)
		} else {
			fmt.Println(v, "impossible type assertion")
		}
	}
	//	310
	//	0 impossible type assertion
}
