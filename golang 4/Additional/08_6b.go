package main

import "fmt"

func f(d []int) {
	for i, _ := range d {
		d[i] = i + 1
	}
	d = append(d, 1, 2, 3, 4, 5)
	fmt.Println(d, len(d), cap(d))
}

func main() {
	var a [10]int
	fmt.Println(len(a), cap(a), a)    //   10 10 [0 0 0 0 0 0 0 0 0 0]
	c := a[2:8]                       //
	f(c)                              //   11 16 [1 2 3 4 5 6 1 2 3 4 5]
	fmt.Println(len(a), cap(a), a)    //   10 10 [0 0 1 2 3 4 5 6 0 0]
	fmt.Println(len(c), cap(c), c)    //   6 8 [1 2 3 4 5 6]
	cc := []int{11, 12, 13, 14}       //
	fmt.Println(len(cc), cap(cc), cc) //   4 4 [11 12 13 14]
	f(cc)                             //   9 10 [1 2 3 4 1 2 3 4 5]
	fmt.Println(len(cc), cap(cc), cc) //   4 4 [1 2 3 4]
}
