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

func main() {
	var Queens [N]queen 
	for col0:= 0; col0 < N; col0++ {
		Queens[0] = queen{col0, 0} 
		if Conflict(Queens[:0], Queens[0]) { continue } 
		for col1:= 0; col1 < N; col1++ {
			Queens[1] = queen{col1, 1} 
			if Conflict(Queens[:1], Queens[1]) { continue } 
			for col2:= 0; col2 < N; col2++ {
				Queens[2] = queen{col2, 2} 
				if Conflict(Queens[:2], Queens[2]) { continue } 
				for col3:= 0; col3 < N; col3++ {
					Queens[3] = queen{col3, 3} 
					if Conflict(Queens[:3], Queens[3]) { continue } 
					for col4:= 0; col4 < N; col4++ {
						Queens[4] = queen{col4, 4} 
						if Conflict(Queens[:4], Queens[4]) { continue } 
						for _, q := range Queens {
							fmt.Printf("%c%d ", q.col+'a', q.row + 1)
						}
						fmt.Println()	
					}	
				}	
			}	
		}	
	}
}
