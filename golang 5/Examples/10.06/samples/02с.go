package main

import "fmt"

func square(d []int) {
	for i, x := range d {
		d[i] = x * x
	}
}

func fill(d []int) {
	for i, _ := range d {
		d[i] = i + 1
	}
}

func main() {
	var a [10]int
	fmt.Println(len(a), cap(a), a) //   10 10 [0 0 0 0 0 0 0 0 0 0]
	d := a[2:8]
	fill(d)
	fmt.Println(len(a), cap(a), a) //   10 10 [0 0 1 2 3 4 5 6 0 0]
	fmt.Println(len(d), cap(d), d) //   6 8 [1 2 3 4 5 6]
	square(a[4:7])
	fmt.Println(len(a), cap(a), a) //   10 10 [0 0 1 2 9 16 25 6 0 0]
	fmt.Println(len(d), cap(d), d) //   6 8 [1 2 9 16 25 6]
}
