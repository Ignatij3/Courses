package main

import "fmt"

const size = 4

func search(sum int, amount int, result []int) {
	if amount == 0  {
		if sum == 0 {
			fmt.Println(result)
		}	
		return
	}	
	for x:= 0; x*x <= sum; x++ {
		search(sum - x*x, amount - 1, append(result, x))
	}	
}	

func main() {
	n:= 10
	search( n, size, make([]int, 0) )
}	
	
