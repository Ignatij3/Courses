package main

import (
	"fmt"
)

type Lmnt struct {
	Next  *Lmnt
	Value int
}

type SortedList struct {
	Head Lmnt // sentinel list Lmnt;
	// it's located at the top of the list
	Len  int // current list length excluding sentinel
	less func(*Lmnt, *Lmnt) bool
}

func NewSortedList(less func(*Lmnt, *Lmnt) bool) SortedList {
	return SortedList{less: less}
}

func (l *SortedList) Add(x int) {
	p := &(l.Head)
	v := &Lmnt{Value: x}
	for i := 0; i < l.Len; i++ {
		if l.less(v, (*p).Next) {
			break
		}
		p = (*p).Next
	}
	(*v).Next = (*p).Next
	p.Next = v
	l.Len++
}

func main() {
	before := func(p, q *Lmnt) bool {
		return (*p).Value > (*q).Value
	}
	l := NewSortedList(before)
	l.Add(2)
	l.Add(4)
	l.Add(1)
	l.Add(8)
	l.Add(5)
	p := l.Head
	for i := 0; i < l.Len; i++ {
		p = *(p.Next)
		fmt.Printf("%v\n", p.Value)
	}
}
