package main

import (
	"fmt"
	"order/sorted"
)

func main() {
	l := sorted.NewSortedList(func(p, q *sorted.Element) bool {
		return (*p).Value < (*q).Value
	})
	l.Add(2)
	l.Add(4)
	l.Add(1)
	l.Add(8)
	l.Add(5)
	p := l.Head
	for i := 0; i < l.Len; i++ {
		p = *(p.Next)
		fmt.Printf("%5d", p.Value)
	}
}
