package main

import "fmt"

func square(num int) int {
    return num * num
}

func mapper(f func(int) int, alist []int) []int { //f func(int) int - Величина функционального типа
	var maplist = make([]int, len(alist), len(alist))
	for i, value := range alist {
		maplist[i] = f(value)
	}
	return maplist
}

func main() {
	alist := []int{3, 11, 6, 4, 14, 9}
	result := mapper(square, alist)
	fmt.Println(result)		// 	[9 121 36 16 196 81]
}
