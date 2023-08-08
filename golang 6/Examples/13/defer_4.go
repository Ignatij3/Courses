package main

import "fmt"

func c(k int) (list []int) {
    defer func(i int) { list = append(list, i) } (k)
	for ; k > 0 ; k-- {
		list = append(list, k)
	}	
    return append(list, 100)
}

func main() {
	fmt.Println(c(3))	// [3 2 1 100 3]
}
