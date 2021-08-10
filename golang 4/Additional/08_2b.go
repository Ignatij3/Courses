package main

import "fmt"

func main() {
	a := make([]int, 3, 7)          //
	fmt.Println(len(a), cap(a), a)  //   3 7 [0 0 0]
	var b []int = make([]int, 5, 7) //
	fmt.Println(len(b), cap(b), b)  //   5 7 [0 0 0 0 0]
	c := a[:6]                      //
	fmt.Println(len(c), cap(c), c)  //   6 7 [0 0 0 0 0 0]
	d := make([]int, 7)             //
	fmt.Println(len(d), cap(d), d)  //   7 7 [0 0 0 0 0 0 0]
}
