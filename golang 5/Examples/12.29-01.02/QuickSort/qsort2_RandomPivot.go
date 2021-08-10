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
	
	k:= rand.Intn(len(l))
	l[0], l[k] = l[k], l[0]
	
	pivot := l[0]
	small, large := 1, len(l)-1
	for { 
		for small < len(l) && l[small] < pivot { small++ }
		for l[large] > pivot { large-- } 
		if small >= large { break }
		l[small], l[large] = l[large], l[small]
		small++
		large--
	}
	l[0], l[large] = l[large], l[0]
	l[:large].QSort()
	l[large+1:]. QSort()
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
