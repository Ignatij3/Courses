package main

import (
	"fmt"
	"os"
)

func Reverse (x []int)  {
	j := len(x) - 1
	for i := 0; i < j; i++ {
		x[i], x[j] = x[j], x[i]
		j--
	}
}

func main() {

	fin, _ := os.Open("a1.dat")
	defer fin.Close()

	// Input data from a1.dat
	var n int
	fmt.Fscanln(fin, &n)
	a := make([]int, n, n)
	for i:= 0; i < n; i++ {
		if _, err := fmt.Fscan(fin, &a[i]); err != nil {
			break
		}
	}
	//Before
	fmt.Println(a)

	Reverse(a)

	//After
	fmt.Println(a)
	
	// Output result into a1.res
	fout, _ := os.Create("a1.res")
	defer fout.Close()
	for _, c := range a {
		fmt.Fprint(fout, c, " ")
	}
	fmt.Fprintln(fout)
}
