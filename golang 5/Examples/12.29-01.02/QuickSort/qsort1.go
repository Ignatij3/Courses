package main

import (
	"fmt"
	"math/rand"
)

type list []int

func InitList(size int, max int) list {
	a:= make([]int, size, size)
	for i:= 0; i < size; i++ {
		a[i] = rand.Intn(max)
	}	
	return a
}	

func (l list) QSort() {
	if len(l) <= 1 { return }
	pivot := l[0]
	division:= 0
	for i, x := range(l) { 
		if x < pivot {
			division++ 
			l[i], l[division] = l[division], l[i]
		}	
	}
	l[0], l[division] = l[division], l[0]
	l[:division].QSort()
	l[division+1:]. QSort()
}	

func (l list) Sort() {
	l.QSort()
}	

func main() {
	list := InitList(10, 100)
	fmt.Println(list)
	list.Sort()
	fmt.Println(list)
}
