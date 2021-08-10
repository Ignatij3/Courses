package main

import "fmt"

const size = 4

func search(sum int) {
	for x1:= 0; x1*x1*4 <= sum; x1++ {
		for x2:= x1; x2*x2*3 <= sum - x1*x1; x2++ {
			for x3:= x2; x3*x3*2 <= sum - x1*x1 - x2*x2; x3++ {
				for x4:= x3; x4*x4 <= sum - x1*x1 - x2*x2 - x3*x3; x4++ {
					if x1*x1 + x2*x2 + x3*x3 + x4*x4 == sum {
						fmt.Println(x1, x2, x3, x4)
					}	
				}
			}
		}			
	}	
}	

func main() {
	n:= 50
	search(n)
}	
	