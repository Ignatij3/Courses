package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3, 4, 5}
	s := a[:3]
	fmt.Println(len(s), cap(s), s) //   3 5 [1 2 3]
	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(len(s), cap(s), s) //   13 14 [1 2 3 1 2 3 4 5 6 7 8 9 10]
	s = append(s, 1, 2, 3)
	fmt.Println(len(s), cap(s), s) //   16 28 [1 2 3 1 2 3 4 5 6 7 8 9 10 1 2 3]
}
