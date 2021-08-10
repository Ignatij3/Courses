package main

import (
	"testing"
	"fmt"
)

func lala(s []int) {
	for i, x := range(s) {
		s[i] = (x-1)*x*(x+1)
	}
}

func main() {
	var s [100000]int
	res:= testing.Benchmark(
		func(b *testing.B) {
			for i:= 0; i < b.N; i++  {
				lala(s[:])
			}
		})
	fmt.Println(res)
}
