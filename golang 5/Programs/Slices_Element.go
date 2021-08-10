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

func main() {
	var (
		number int
		step uintptr
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
	first := getSlice(slice)
	cont = ""
	
	for cont != "n" {
		fmt.Print("Enter place of the number:")
		fmt.Scan(&number)
		
		for number >= len(slice) || number < 0 {
			fmt.Print("Error, inaproppriate number, try again:")
			fmt.Scan(&number)
		}
		step = uintptr(number) * unsafe.Sizeof(slice[0])
		ch := (*int)(unsafe.Pointer(uintptr((*first) .start) + step))
		
		fmt.Print("Enter the replacing number:")
		fmt.Scan(&number)
		*ch = number
		
		fmt.Printf("slice - %v\nlen - %d\ncap - %d\n\n", slice, len(slice), cap(slice))
		first = getSlice(slice)
		
		fmt.Print("Do you wish to continue? (y/n)")
		fmt.Scan(&cont)
		for cont != "n" && cont != "y" {
			fmt.Print("Error, try again (y/n)")
			fmt.Scan(&cont)
		}
	}
}
