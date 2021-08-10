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
	var start int
	if len(result) == 0 {
		start = 0
	} else {
		start = result[len(result)-1]	
	}	
	for x:= start; x*x*amount <= sum; x++ {
		search(sum - x*x, amount - 1, append(result, x))
	}	
}	

func main() {
	search(500000, size, make([]int,0))
}	
	
