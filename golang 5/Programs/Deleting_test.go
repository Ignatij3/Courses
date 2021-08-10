package main

import "testing"

func MyCopy(receiver, sender []int) []int {
	for k := 0; k < len(receiver) && k < len(sender); k++ {
		receiver[k] = sender[k]
	}
	return receiver
}

func BenchmarkDelete01(b *testing.B) {
	var i int
	a := make([]int, 1000000)
	for k := 0; k < b.N; k++ {
		if len(a) == 0 {
			a = make([]int, 1000000)
		} else {
			a = append(a[:i], a[i+1:]...)
		}
	}
}

func BenchmarkDelete02(b *testing.B) {
	var i int
	a := make([]int, 1000000)
	for k := 0; k < b.N; k++ {
		if len(a) == 0 {
			a = make([]int, 1000000)
		} else {
			a = a[:i+copy(a[i:], a[i+1:])]
		}
	}
}

func BenchmarkDelete02MyCopy(b *testing.B) {
	var (
		i int
		i2 []int
	)
	a := make([]int, 1000000)
	for k := 0; k < b.N; k++ {
		if len(a) == 0 {
			a = make([]int, 1000000)
		} else {
			i2 = MyCopy(a[i:], a[i+1:])
			a = a[:i+i2[0]]
		}
	}
}

func BenchmarkSafeDelete(b *testing.B) {
	var i int
	a := make([]int, 1000000)
	for k := 0; k < b.N; k++ {
		if len(a) == 0 {
			a = make([]int, 1000000)
		} else {
			copy(a[i:], a[i+1:])
			a = a[:len(a) - 1]
		}
	}
}

func BenchmarkSafeDeleteMyCopy(b *testing.B) {
	var i int
	a := make([]int, 1000000)
	for k := 0; k < b.N; k++ {
		if len(a) == 0 {
			a = make([]int, 1000000)
		} else {
			a = MyCopy(a[i:], a[i+1:])
			a = a[:len(a) - 1]
		}
	}
}
