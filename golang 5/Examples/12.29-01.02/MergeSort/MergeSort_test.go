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
 
func InitList() {
// Возвращает слайс длины len, заполненный
// случайными числами от 1 до max
	numbers = make([]int, size, size)
	for i:= 0; i < size; i++ {
		numbers[i] = rand.Intn(MaxValue) + 1
	}
}


func (l list) Sort0() {
// Библиотечная сортировка 	
	sort.Slice(l, func(i, j int) bool { return l[i] < l[j] })
}	

func (l list) Merge(start2 int) {
// Слайс l состоит из двух возрастающих серий: l[:start2]] и l[start2:]
// В результате получаем отсортированный слайс l	
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

func (l list) BinaryMergeSortRecursive() {
	if len(l) <= 1 { return }
	l[:len(l)/2].BinaryMergeSortRecursive()
	l[len(l)/2:].BinaryMergeSortRecursive()
	l.Merge(len(l)/2)
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

func (l list) BinaryMergeSortUnRecursive() {
	// инициализация - все серии состоят из одного элемента
	var start []int
	for i := 0; i <= len(l); i++ {
		start = append(start, i)
	}	
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
// Названия тестируемых функций должны начинаться на Benchmark,
// за которым идёт название, начинающееся с большой буквы

func BenchmarkSort0(b *testing.B) {
    for i := 0; i < b.N; i++ {
        numbers.Sort0()
    }
}

func BenchmarkBinaryMergeSortRecursive(b *testing.B) {
    for i := 0; i < b.N; i++ {
        numbers.BinaryMergeSortRecursive()
    }
}

func BenchmarkNaturalMergeSort(b *testing.B) {
    for i := 0; i < b.N; i++ {
        numbers.NaturalMergeSort()
    }
}

func BenchmarkBinaryMergeSortUnRecursive(b *testing.B) {
    for i := 0; i < b.N; i++ {
        numbers.BinaryMergeSortUnRecursive()
    }
}

