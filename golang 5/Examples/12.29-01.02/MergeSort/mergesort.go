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

func (l list) Merge(start2 int) {
	res:= make([]int, len(l), len(l))
	i1, i2, ires:= 0, start2, 0
	for i1 < start2 && i2 < len(l) {
		if l[i1] < l[i2] {
			res[ires] = l[i1]
			i1++
		} else {
			res[ires] = l[i2]
			i2++
		}
		ires++
	}
	if i2 == len(l) {
		copy (l[ires:], l[i1:])
	}		 
	copy (l, res[:ires])
}	

func (l list) BinaryMergeSort() {
	if len(l) <= 1 { return }
	l[:len(l)/2].BinaryMergeSort()
	l[len(l)/2:].BinaryMergeSort()
	l.Merge(len(l)/2)
}	

func main() {
	list := InitList(10, 100)
	fmt.Println(list)
	list.BinaryMergeSort()
	fmt.Println(list)
}
