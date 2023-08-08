package main

import (  
    "fmt"
)

type Ordered interface {
	Before (b Ordered) bool
}

type Integer int

func (a Integer) Before(b Ordered) bool {
	return a < b.(Integer)
}

func main() { 
	d:= []Integer{3,1,0}
	var x Ordered
	for _, x = range d {
		fmt.Println(x.(float64))
	}
// 02b.go:21:16: impossible type assertion:
//	float64 does not implement Ordered (missing Before method)
}
