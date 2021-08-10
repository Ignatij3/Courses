package main

import "fmt"

func main() {
	a := []int{3, 24, 17, 93, 31, 44, 60}
	fmt.Println(len(a), cap(a), a) //   7 7 [3 24 17 93 31 44 60]
	b := a[2:5]                    //
	fmt.Println(len(b), cap(b), b) //   3 5 [17 93 31]
}
