package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	s := a[:5]
	fmt.Println(a) //   [1 2 3 4 5 6 7 8 9 10 11 12 13 14]
	fmt.Println(s) //   [1 2 3 4 5]
	s = append(s, 21)
	fmt.Println(a) //   [1 2 3 4 5 21 7 8 9 10 11 12 13 14]
	fmt.Println(s) //   [1 2 3 4 5 21]
	s = append(s, 41, 42, 43)
	fmt.Println(a) //   [1 2 3 4 5 21 41 42 43 10 11 12 13 14]
	fmt.Println(s) //   [1 2 3 4 5 21 41 42 43]
	s = append(s, a[12:]...)
	fmt.Println(a) //   [1 2 3 4 5 21 41 42 43 13 14 12 13 14]
	fmt.Println(s) //   [1 2 3 4 5 21 41 42 43 13 14]
}
