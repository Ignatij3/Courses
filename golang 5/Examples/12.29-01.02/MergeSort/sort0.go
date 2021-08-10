package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type list []int

func InitList(size int, max int) list {
	a:= make([]int, size, size)
	for i:= 0; i < size; i++ {
		a[i] = rand.Intn(max)
	}	
	return a
}	

func (l list) Sort() {
	sort.Slice(l, func(i, j int) bool { return l[i] < l[j] })
}	

func main() {
	list := InitList(10, 100)
	fmt.Println(list)
	list.Sort()
	fmt.Println(list)
}
