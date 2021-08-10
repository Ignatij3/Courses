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

func (l list) NaturalMergeSort() {
	// инициализация - заполняем слайс start
	start:= []int{0}
	for i:= 1; i < len(l); i++ {
		if l[i] < l[i-1] {
			start = append(start, i)
		}	
	}	
	start = append(start, len(l))
	// сортировка
	for len(start) > 2 {
		// проходим по всему массиву, склеивая пары соседних серий
		for k:= 0; k < len(start) - 2; k += 2 {
			l[start[k]:start[k+2]].Merge(start[k+1]-start[k])
		}	
		// преобразуем слайс start: start[2] -> start[1], 
		// start[4] -> start[2], start[6] -> start[3] и т.д.
		k:= 0; 	
		for {
			k += 2
			if k >= len(start) { break }
			start[k/2] = start[k]
		}
		start = start[:k/2]
		// если перед этим было нечётное количество серий, то 
		// надо добавить конец последней серии - len(l)
		if start[len(start)-1] < len(l) {
			start = append(start, len(l))
		}	
	}	
}	

func main() {
	list := InitList(25, 100)
	fmt.Println(list)
	list.NaturalMergeSort()
	fmt.Println(list)
}
