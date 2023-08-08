package main

// Запуск из командной строки:
//         go test -bench . benchmark_test.go
// В данном случае benchmark_test.go - это имя файла,
// в котором находится данная программа.

import (
	"container/skiplist"
	"math"
	"math/rand"
	"testing"
)

var d [2048]int

func init() {
	for i, _ := range d {
		d[i] = rand.Intn(1000000000)
	}
}

func benchmarkSort(b *testing.B, sort func([]int)) {
	for _, tc := range []struct {
		name string
		data []int
	}{
		{"64", d[:64]},
		{"128", d[:128]},
		{"256", d[:256]},
		{"512", d[:512]},
		{"1024", d[:1024]},
		{"2048", d[:2048]},
	} {
		b.Run(tc.name, func(b *testing.B) {
			d := make([]int, len(tc.data))
			for i := 0; i < b.N; i++ {
				copy(d, tc.data)
				sort(d)
			}
		})
	}
}

func BenchmarkSortLinkedList1(b *testing.B) { benchmarkSort(b, SortLinkedList1) }
func BenchmarkSortLinkedList2(b *testing.B) { benchmarkSort(b, SortLinkedList2) }
func BenchmarkSortSkipList(b *testing.B)    { benchmarkSort(b, SortSkipList) }

// implementations

const (
	PlusInfinity  = math.MaxInt64
	MinusInfinity = math.MinInt64
)

type (
	list struct {
		head *lmnt
	}
	lmnt struct {
		x    int
		next *lmnt
	}
)

func SortLinkedList1(data []int) {
	l := NewList()
	for _, x := range data {
		l.Insert1(x)
	}
}

func SortLinkedList2(data []int) {
	l := NewList()
	for _, x := range data {
		l.Insert2(x)
	}
}

func SortSkipList(data []int) {
	l := skiplist.NewSkipList(12, 0.5, func(a, b interface{}) bool { return a.(int) < b.(int) })
	for _, x := range data {
		l.Insert(x)
	}
}

func NewList() list {
	return list{&lmnt{MinusInfinity, &lmnt{PlusInfinity, nil}}}
}

func (s *list) Insert1(num int) {
	runner := (*s).head
	for (*(*runner).next).x < num {
		runner = (*runner).next
	}
	(*runner).next = &lmnt{num, (*runner).next}
}

func (s *list) Insert2(num int) {
	runner := (*s).head
	runner2 := (*runner).next
	for (*runner2).x < num {
		runner, runner2 = runner2, (*runner2).next
	}
	(*runner).next = &lmnt{num, runner2}
}
