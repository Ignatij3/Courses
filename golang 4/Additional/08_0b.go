package main

import "fmt"

func main() {
	coins := [...]int{1, 2, 5, 10, 20, 50, 100, 200}
	fmt.Println(len(coins), cap(coins), coins) //   8 8 [1 2 5 10 20 50 100 200]
	cent := coins[0:6]                         //
	fmt.Println(len(cent), cap(cent), cent)    //   6 8 [1 2 5 10 20 50]
	cent[6] = -1                               //   panic: runtime error: index out of range
}
