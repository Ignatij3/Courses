package main

import (
	"fmt"
	"math"
)	

const size = 4

func PerfectSquare (n int) (sqrt int, is bool) {
	sqrt = int( math.Round( math.Sqrt( float64(n) ) ) )
	return sqrt, sqrt*sqrt==n
}
	
func search(sum int, amount int, result []int) {
	if amount == 1  {
		if x, ok:= PerfectSquare(sum); ok {
			fmt.Println(append(result, x))
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
	search(500, size, make([]int,0))
}	
	
