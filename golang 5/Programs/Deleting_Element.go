package main

import (
	"fmt"
	"unsafe"
)

type (
	sliceHeader struct {
		start unsafe.Pointer
		len uint
		cap uint
	}
)

func getSlice(s []int) *sliceHeader {
	return (*sliceHeader)(unsafe.Pointer(&s))
}

func ReplaceSlice(slice []int, step int) []int {
	var rSlice []int
	for i, k := range slice {
		if i != step {rSlice = append(rSlice, k)}
	}
	return rSlice
}

func main() {
	var (
		number int
		slice []int
		cont string
	)
	
	fmt.Println("Enter numbers you want to include in slice")
	for cont != "n" {
		fmt.Print("Enter number:")
		fmt.Scan(&number)
		slice = append(slice, number)
		fmt.Print("Do you wish to continue? (y/n)")
		fmt.Scan(&cont)
		for cont != "n" && cont != "y" {
			fmt.Print("Error, try again (y/n)")
			fmt.Scan(&cont)
		}
	}
	
	fmt.Printf("slice - %v\nlen - %d\ncap - %d\n\n", slice, len(slice), cap(slice))
	cont = ""
	
	for cont != "n" {
		fmt.Print("Enter place of the number you want to erase:")
		fmt.Scan(&number)
		
		for number >= len(slice) || number < 0 {
			fmt.Print("Error, inaproppriate number, try again:")
			fmt.Scan(&number)
		}
		slice = ReplaceSlice(slice, number)
		fmt.Printf("slice - %v\nlen - %d\ncap - %d\n\n", slice, len(slice), cap(slice))
		
		if len(slice) == 0 {break}
		fmt.Print("Do you wish to continue? (y/n)")
		fmt.Scan(&cont)
		for cont != "n" && cont != "y" {
			fmt.Print("Error, try again (y/n)")
			fmt.Scan(&cont)
		}
	}
}
