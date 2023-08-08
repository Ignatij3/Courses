package main

import "fmt"

func reversePrintList(nums []int) {
	defer fmt.Println("<- head of list")
	fmt.Print("tail of list -> ")
	for _, x := range nums {
		defer fmt.Printf("%d  ", x)
	}
}

func main() {
	a := []int{11, 22, 33, 44, 55}
	reversePrintList(a)
	fmt.Println("The end")
}
//	tail of list -> 55  44  33  22  11  <- head of list
//	The end
