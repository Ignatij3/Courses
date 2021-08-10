package main

import "fmt"

func lalala(c []int) {
	cc := make([]int, 0, 0)
	cc = append(cc, c...)
	for i, _ := range cc {
		cc[i] = i + 1
	}
	fmt.Println(len(c), cap(c), c)    //   4 7 [0 0 0 0]
	fmt.Println(len(cc), cap(cc), cc) //   4 4 [1 2 3 4]
}

func main() {
	var a [10]int
	fmt.Println(len(a), cap(a), a) //   10 10 [0 0 0 0 0 0 0 0 0 0]
	d := a[3:7]                    //
	lalala(d)                      //
	fmt.Println(len(a), cap(a), a) //   10 10 [0 0 0 0 0 0 0 0 0 0]
	fmt.Println(len(d), cap(d), d) //   4 7 [0 0 0 0]
}
