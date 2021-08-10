package main

// Запуск из командной строки:
//         go test -bench . benchmark_4squares_test.go
// Имя файла обязательно должно заканиваться на _test

import (
	"testing"
	"math/rand"
	"sort"
)	

const (
	size = 1000000
	MaxValue = 2000000000
)
	
type list []int
 
var numbers list
 
func init() {
	numbers = make([]int, size, size)
	/*
	for i:= 0; i < len; i++ {
		numbers[i] = rand.Intn(MaxValue) + 1
	}
	*/
	for i:= 0; i < size; i++ {
		numbers[i] = i
	}
	for i:= 0; i<100; i++ {
		a, b:= rand.Intn(size), rand.Intn(size)
		numbers[a], numbers[b] = numbers[b], numbers[a]
	}
}

func (l list) Sort0() {
	sort.Slice(l, func(i, j int) bool { return l[i] < l[j] })
}	

func (l list) QSort1() {
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
	l[:division].QSort1()
	l[division+1:]. QSort1()
}	

func (l list) Sort1() {
	l.QSort1()
}	

func (l list) QSort2() {
	if len(l) <= 1 { return }
	
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
	l[:large].QSort2()
	l[large+1:]. QSort2()
}	

func (l list) Sort2() {
	l.QSort2()
}

func (l list) QSort2RandomPivot() {
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
	l[:large].QSort2RandomPivot()
	l[large+1:]. QSort2RandomPivot()
}	

func (l list) Sort2RandomPivot() {
	l.QSort2RandomPivot()
}

const cutoff = 42

func (l list) QSort1Sedgewick() {
	if len(l) <= cutoff { return }
	
	pivot := l[0]
	division:= 0
	for i, x := range(l) { 
		if x < pivot {
			division++ 
			l[i], l[division] = l[division], l[i]
		}	
	}
	l[0], l[division] = l[division], l[0]
	l[:division].QSort1Sedgewick()
	l[division+1:]. QSort1Sedgewick()
	
}	

func (l list) Sort1Sedgewick() {
	l.QSort1Sedgewick()
	// Постобработка - сортировка простыми вставками
	for i, x:= range(l) {
		j:= i
		for j > 0 && l[j-1] > x  { 
			l[j] = l[j-1]
			j-- 
		}	
		l[j] = x
	}	
	
}	

func (l list) QSort2Sedgewick() {
	if len(l) <= cutoff { return }
	
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
	l[:large].QSort2Sedgewick()
	l[large+1:]. QSort2Sedgewick()
}	

func (l list) Sort2Sedgewick() {
	l.QSort2Sedgewick()
	for i, x:= range(l) {
		j:= i
		for j > 0 && l[j-1] > x  { 
			l[j] = l[j-1]
			j-- 
		}	
		l[j] = x
	}	
}	

func (l list) QSort2Recursiveless() {
	type segment struct {division, right int}
	stack := []segment { segment{0, len(l)-1} }
	
	for len(stack) > 0 {
		workingSegment:= stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if workingSegment.division >= workingSegment.right { continue }
		pivot := l[workingSegment.division]
		small, large := workingSegment.division+1, workingSegment.right
		for { 
			for small <= workingSegment.right && l[small] < pivot { small++ }
			for l[large] > pivot { large-- } 
			if small >= large { break }
			l[small], l[large] = l[large], l[small]
			small++
			large--
		}
		l[workingSegment.division], l[large] = l[large], l[workingSegment.division]
		if large - workingSegment.division > workingSegment.right - large {
			stack = append(stack, segment{workingSegment.division, large - 1},
								  segment{large + 1, workingSegment.right})	
		} else {	
			stack = append(stack, segment{large + 1, workingSegment.right},
								  segment{workingSegment.division, large - 1} )	
		}
	}	
}	

func (l list) Sort2Recursiveless() {
	l.QSort2Recursiveless()
}	

// Названия тестируемых функций должны начинаться на Benchmark,
// за которым идёт название, начинающееся с большой буквы

func BenchmarkSort0(b *testing.B) {
    for i := 0; i < b.N; i++ {
        numbers.Sort0()
    }
}

func BenchmarkSort1(b *testing.B) {
    for i := 0; i < b.N; i++ {
        numbers.Sort1()
    }
}

func BenchmarkSort2(b *testing.B) {
    for i := 0; i < b.N; i++ {
        numbers.Sort2()
    }
}

func BenchmarkSort2RandomPivot(b *testing.B) {
    for i := 0; i < b.N; i++ {
        numbers.Sort2RandomPivot()
    }
}

func BenchmarkSort1Sedgewick(b *testing.B) {
    for i := 0; i < b.N; i++ {
        numbers.Sort1Sedgewick()
    }
}

func BenchmarkSort2Sedgewick(b *testing.B) {
    for i := 0; i < b.N; i++ {
        numbers.Sort2Sedgewick()
    }
}

func BenchmarkSort2Recursiveless(b *testing.B) {
    for i := 0; i < b.N; i++ {
        numbers.Sort2Recursiveless()
    }
}
