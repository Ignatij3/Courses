package main

import (
	"sorted"
)

type Integer int

func (a Integer) Before(b sorted.Ordered) bool {
	return a < b.(Integer)
}

func main() {
	data := []Integer{5, 8, 2, 4, 3, 2, 9, 7}
	var sortdata sorted.SortedCollection
	var x sorted.Ordered
	for _, x = range data {
		sortdata.Insert(x)
	}
	sortdata.Print()
}
