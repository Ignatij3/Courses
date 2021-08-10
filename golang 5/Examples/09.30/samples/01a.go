package main

import "fmt"

func main() {
	k:= 12345
	fmt.Printf("Type of k is %T, value of k is %d\n\n", k, k)
	p1:= &k
	fmt.Printf("Type of p1 is %T\n", p1)
	fmt.Println("address of k is", p1)
	fmt.Printf("the number %d is located at address %p\n\n", *p1, p1)
	var p2 *int = &k
	fmt.Printf("Type of p2 is %T\n", p2)
	fmt.Printf("address of k is %p\n", p2)
	fmt.Printf("numerical value of address of k is %x\nor %b or %d\n\n", p2, p2, p2)

/*
Type of k is int, value of k is 12345

Type of p1 is *int
address of k is 0xc00000e0b0
the number 12345 is located at address 0xc00000e0b0

Type of p2 is *int
address of k is 0xc00000e0b0
numerical value of address of k is c00000e0b0
or 1100000000000000000000001110000010110000 or 824633778352
*/

}
