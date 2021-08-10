package main

import "fmt"

const N = 5

type queen struct {
	          col int
	          row int
	       }     

func abs(x int) int {
	if x >= 0  {
		return x 
	} else {
		return -x
	}
}			

func Connected (q1, q2 queen) bool {
	return q1.col == q2.col ||
	       q1.row == q2.row ||
	       abs(q1.col - q2.col) == abs (q1.row - q2.row) 
}

func Conflict (qs []queen, q queen) bool {
	for _, q2 := range qs {
		if Connected (q, q2) { return true }	
	}	
	return false
}	

var Queens [N]queen 

func Search(n int) {
	if n == 0  {
		for _, q := range Queens {
			fmt.Printf("%c%d ", q.col+'a', q.row + 1)
		}
		fmt.Println()	
		return
	}	
	for col:= 0; col < N; col++ {
		Queens[N-n] = queen{col, N-n} 
		if Conflict(Queens[:N-n], Queens[N-n]) { continue } 
		Search(n-1)
	}	
}	

func main() {
	Search(N)
}
