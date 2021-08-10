package main

import (
	"fmt"
)

func main() {
	var x, y float32
	x = 1.2
	fmt.Println(x)		// 1.2
	y = x / 1.0E25
	y = y / 1.0E20
	x = y * 1.0E20
	x = x * 1.0E25	
	fmt.Println(x)		// 1.4012984
}
 
