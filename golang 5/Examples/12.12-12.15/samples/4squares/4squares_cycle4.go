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
	
func search(sum int) {
	for x1:= 0; x1*x1*4 <= sum; x1++ {
		for x2:= x1; x2*x2*3 <= sum - x1*x1; x2++ {
			for x3:= x2; x3*x3*2 <= sum - x1*x1 - x2*x2; x3++ {
				if x4, ok:= PerfectSquare(sum - x1*x1 - x2*x2 - x3*x3); ok {
						fmt.Println(x1, x2, x3, x4)
				}
			}
		}			
	}	
}	

func main() {
	n:= 500
	search(n)
}	
	
