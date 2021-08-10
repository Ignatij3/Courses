package main

import (
	"fmt"
	"time"
	"math/rand"
)

func FillSlice() []int {
	var (
		x int
		a []int
	)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 20000; i++ {
		x = rand.Intn(10001)
		switch rand.Intn(2) {
			case 0:
				a = append(a, x)
			case 1:
				a = append(a, -x)
		}
	}
	return a
}

func Sort(a []int) []int {
	//fmt.Println(comp)
	if len(a) < 2 {return a}
	left, right := 0, len(a) - 1
	pivotIndex := rand.Int() % len(a)
	a[pivotIndex], a[right] = a[right], a[pivotIndex]
	// Pile elements smaller than the pivot on the left	
	for i := range a {
		if a[i] < a[right] {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}
	// Place the pivot after the last smaller element
	a[left], a[right] = a[right], a[left]
	Sort(a[:left])
	Sort(a[left + 1:])
	return a
}

func Find(s []int, x int) bool {
	a, b := 0, len(s) - 1
	for {
		if a == b || a + 1 == b {break}
		if s[(a+b)/2] == x {
			return true
		} else if s[(a+b)/2] > x {
			b = (a+b)/2
		} else if s[(a+b)/2] < x {
			a = (a+b)/2
		}
	}
	return false
}

func main() {
	var (
		found bool
		x int
		a []int
	)
	
	a = FillSlice()
	a = Sort(a)
	for {
		fmt.Print("Введите число, которое хотите найти:")
		fmt.Scan(&x)
		fmt.Println()
		found = Find(a, x)
		if found {
			fmt.Println("Число было успешно найдено")
		} else {
			fmt.Println("Число не было найдено")
		}
	}
}
