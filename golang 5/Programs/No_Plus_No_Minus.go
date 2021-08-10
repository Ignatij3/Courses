package main

import (
	"fmt"
)

func PP(a uint32) uint32 {
	return a+1
}

func MM(a uint32) uint32 {
	if a > 0 {return a-1} else {return 0}
}

func Plus(a, b uint32) uint32 {
	/*for i := uint32(0); i < b; i = PP(i) {
		a = PP(a)
	}*/
	
	/*for b > 0 {
		a = PP(a)
		b = MM(b)
	}*/
	
	if b == 0 {
		return a
	} else {
		return Plus(PP(a), MM(b))
	}
}

func main() {
	fmt.Println(Plus(5, 7))
	fmt.Println(MM(0))
}
