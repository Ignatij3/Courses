package main

import "fmt"

const size = 4

func search(sum int) {
	for x1:= 0; x1*x1 <= sum; x1++ {
		for x2:= 0; x2*x2 <= sum; x2++ {
			for x3:= 0; x3*x3 <= sum; x3++ {
				for x4:= 0; x4*x4 <= sum; x4++ {
					if x1*x1 + x2*x2 + x3*x3 + x4*x4 == sum {
						fmt.Println(x1, x2, x3, x4)
					}	
				}
			}
		}			
	}	
}	

func main() {
	n:= 10
	search(n)
}	
	
