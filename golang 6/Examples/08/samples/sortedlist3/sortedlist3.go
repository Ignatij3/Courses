package main

import (
	"fmt"
	"order/sorted"
)

func main() {
	l := sorted.NewSortedList(func(p, q *sorted.Element) bool {
		return (*p).Value.(int) < (*q).Value.(int)
	})
	l.Add(2)
	l.Add(4)
	l.Add(1)
	l.Add(8)
	l.Add(5)

	l.Do(func(v interface{}) {
		fmt.Printf("%5d", v.(int))
	})

}
