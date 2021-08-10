package main

import (
	"fmt"
	"unsafe"
)

type (
	Slice struct {
        start unsafe.Pointer
        len int
        cap int
	}
)

func SliceCh(s []int) *Slice {
	return (*Slice)(unsafe.Pointer(&s))
}
/*
func lenS(sl Slice) {
	(*sl).len += 15
}*/

func main() {
	slice := make([]int, 0)
	slice = append(slice, 25, 14, 74, 2, 9)
	
	sl := SliceCh(slice)
	new := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(sl))+8))
	fmt.Println(sl.start)
	fmt.Println(new)
	fmt.Println(*new)
	fmt.Printf("slice - %v\nlen - %d\ncap - %d\n\n", slice, len(slice), cap(slice))
	fmt.Println(sl.len)
	fmt.Println((*sl).len)
	//lenS(sliceS)
	*new -= 1
	fmt.Println(*sl)
	
	fmt.Printf("slice - %v\nlen - %d\ncap - %d\n\n", slice, len(slice), cap(slice))
	//slice = capS(sl)
	//fmt.Printf("slice - %v\nlen - %d\ncap - %d\n\n", slice, len(slice), cap(slice))
}
