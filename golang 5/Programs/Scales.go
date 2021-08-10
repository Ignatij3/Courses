package main

import "fmt"

func Copy(a []int, n int) []int {
	b := make([]int, len(a))
	copy(b, a)
	Delete(&b, n)
	return b
}

func Delete(b *[]int, n int) {
	copy((*b)[n:], (*b)[n+1:])
	(*b)[len((*b))-1] = 0
	(*b) = (*b)[:len((*b))-1]
}

func Same(rep *[]int, num int) bool {
	for _, n := range *rep {
		if n == num {return true}
	}
	return false
}

func Scales(left, right int, weights []int, repeat *[]int) {
	if left == right {
		if !Same(repeat, left) {fmt.Printf("Total weight on 1 cup - %v\n", left)}
		*repeat = append(*repeat, left)
		return
	}
	if len(weights) == 0 {return}
	for n, k := range weights {
		Scales(left, right + k, Copy(weights, n), repeat)
		Scales(left + k, right, Copy(weights, n), repeat)
	}
}

func main() {
	weights := []int{10, 100, 3, 1, 5, 20, 4, 50, 15, 2}
	repeat := make([]int, 0)
	item := 25
	Scales(item, 0, weights, &repeat)
}
