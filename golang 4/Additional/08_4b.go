package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	s := a[:6]
	s2 := s
	fmt.Println(len(a), cap(a), a)    //   8 8 [1 2 3 4 5 6 7 8]
	fmt.Println(len(s), cap(s), s)    //   6 8 [1 2 3 4 5 6]
	fmt.Println(len(s2), cap(s2), s2) //   6 8 [1 2 3 4 5 6]
	s = append(s, -1, -2, -3, -4)     //
	fmt.Println(len(a), cap(a), a)    //   8 8 [1 2 3 4 5 6 7 8]
	fmt.Println(len(s), cap(s), s)    //   10 16 [1 2 3 4 5 6 -1 -2 -3 -4]
	s2 = append(s2, 0)                //
	fmt.Println(len(a), cap(a), a)    //   8 8 [1 2 3 4 5 6 0 8]
	fmt.Println(len(s2), cap(s2), s2) //   7 8 [1 2 3 4 5 6 0]
	s2 = append(s2, -11, 12, 13, 14, 15)
	fmt.Println(len(a), cap(a), a)    //   8 8 [1 2 3 4 5 6 0 8]
	fmt.Println(len(s2), cap(s2), s2) //   12 16 [1 2 3 4 5 6 0 -11 12 13 14 15]
}
