package main

import (
	"fmt"
	"sort"
)

func main() {
	a, b := 0.0, 3.0 // the root is separated on a segment [a; b]
	// such as f(a) < 0, f(b) > 0
	epsilon := 1.0e-6 // accuracy
	n := int((b-a)/(epsilon*2.0)) + 1
	epsilon2 := (b - a) / float64(n)
	c := func(i int) bool {
		x := a + float64(i)*epsilon2
		return x*x-2.0 >= 0.0
	}
	index := sort.Search(n+1, c)
	fmt.Printf("The first point with a positive function value is %d\n", index)
	fmt.Printf("Accordingly, the solution of the equation is %f\n", a+(float64(index)-0.5)*epsilon2)
}
