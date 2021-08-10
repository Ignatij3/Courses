package main

import "fmt"

func main() {
	a := [7]int{0, 1, 2, 5, 10, 20, 50}
	var b []int = a[1:4] //    b - слайс (отрезок) массива a от a[1] до a[3]
	fmt.Println(b)       //    [1 2 5]
	for i := 0; i < 3; i++ {
		b[i] *= i
	}
	fmt.Println(b) //   [0 2 10]
	fmt.Println(a) //   [0 0 2 10 10 20 50]
	for i, x := range a {
		a[i] = x + 15
	}
	var c []int
	c = a[2:6]     //   c - слайс (отрезок) массива a от a[2] до a[5]
	fmt.Println(a) //   [15 15 17 25 25 35 65]
	fmt.Println(b) //   [15 17 25]
	fmt.Println(c) //   [17 25 25 35]
	for i, x := range c {
		c[i] = x - 35
	}
	fmt.Println(a) //   [15 15 -18 -10 -10 0 65]
	fmt.Println(b) //   [15 -18 -10]
	fmt.Println(c) //   [-18 -10 -10 0]
	d := a[2:3]
	fmt.Println(d) //   [-18]
	d = a[:4]
	fmt.Println(d) //   [15 15 -18 -10]
	d = a[3:]
	fmt.Println(d) //   [-10 -10 0 65]
}
