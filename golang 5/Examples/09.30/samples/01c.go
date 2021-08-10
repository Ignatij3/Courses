package main

import "fmt"

func main() {  
	a := 25
	var b *int
	if b == nil {
		fmt.Println("b is", b)
		fmt.Printf("b is %p or %v\n", b, b)
		b = &a
	}
	fmt.Println("b after initialization is", b)
/*
b is <nil>
b is 0x0
b after initialization is 0xc00000e0b0
*/
}

