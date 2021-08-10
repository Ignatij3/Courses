package main

import "fmt"

func main() {
	coins := [...]int{1, 2, 5, 10, 20, 50, 100, 200}
	fmt.Println(len(coins), cap(coins), coins) //   8 8 [1 2 5 10 20 50 100 200]
	cent := coins[0:6]                         //
	fmt.Println(len(cent), cap(cent), cent)    //   6 8 [1 2 5 10 20 50]
	cent = cent[3:5]                           //
	fmt.Println(len(cent), cap(cent), cent)    //   2 5 [10 20]
	cent = cent[0:5]                           //
	fmt.Println(len(cent), cap(cent), cent)    //   5 5 [10 20 50 100 200]
	cent = cent[3:5]                           //
	fmt.Println(len(cent), cap(cent), cent)    //   2 2 [100 200]
	cent = cent[0:3]                           //   panic: runtime error: slice bounds out of range
}
