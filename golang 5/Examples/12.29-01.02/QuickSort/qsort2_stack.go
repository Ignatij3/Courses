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

func (l list) QSortRecursiveless() {
	
	type segment struct {left, right int}
	// начальный вызов "рекурсии" - начинаем со всего массива
	stack := []segment { segment{0, len(l)-1} }
	// в stack храним параметры, которые в рекурсивном варианте мы 
	// передавали бы отложенным вызовам, когда до них дойдёт очередь
	for len(stack) > 0 {
		// снимаем с верхушки стека верхний отрезок - вход в "рекурсию"
		workingSegment:= stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// обработка "рекурсивного" вызова
		if workingSegment.left >= workingSegment.right { continue }
		pivot := l[workingSegment.left]
		small, large := workingSegment.left+1, workingSegment.right
		for { 
			for small <= workingSegment.right && l[small] < pivot { small++ }
			for l[large] > pivot { large-- } 
			if small >= large { break }
			l[small], l[large] = l[large], l[small]
			small++
			large--
		}
		l[workingSegment.left], l[large] = l[large], l[workingSegment.left]
		// Кладём в стек сначала больший из вдух отрезков, 
		// потом меньший, его мы сразу же вытащим из стека 
		// в начале большого цикла ( который for len(stack) > 0 )
		if large - workingSegment.left > workingSegment.right - large {
			stack = append(stack, segment{workingSegment.left, large - 1},
								  segment{large + 1, workingSegment.right})	
		} else {	
			stack = append(stack, segment{large + 1, workingSegment.right},
								  segment{workingSegment.left, large - 1} )	
		}
	}	
}	

func (l list) Sort() {
	l.QSortRecursiveless()
}	

func main() {
	list := InitList(10, 100)
	fmt.Println(list)
	list.Sort()
	fmt.Println(list)
}
