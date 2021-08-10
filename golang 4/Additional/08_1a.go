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
	fmt.Println(a) //  [0 0 0 0 0 0 0 0 0 0]
	fill(a[2:8])
	fmt.Println(a) //  [0 0 1 2 3 4 5 6 0 0]
	square(a[4:7])
	fmt.Println(a) //  [0 0 1 2 9 16 25 6 0 0]
}
