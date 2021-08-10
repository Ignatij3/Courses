package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int)
	fmt.Println(cap(ch1), len(ch1))		//  0 0
	ch2 := make(chan int, 5)
	ch2 <- 1
	ch2 <- 2
	ch2 <- 1
	fmt.Println(cap(ch2), len(ch2))		//  5 3
	fmt.Println(<-ch2)							//  1
	fmt.Println(<-ch2)							//  2
	fmt.Println(cap(ch2), len(ch2))		//  5 1
	ch2 <- 3
	fmt.Println(cap(ch2), len(ch2))		//  5 2
	close(ch2)
	fmt.Println(cap(ch2), len(ch2))		//  5 2
	for x:= range ch2  {
		fmt.Printf("%d ",x)						//  1 3
	}	
	fmt.Println()
	fmt.Println(cap(ch2), len(ch2))		//  5 0
}
