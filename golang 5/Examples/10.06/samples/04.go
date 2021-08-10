package main

import  (
	"fmt"
	"unsafe"
)	

type (
	Slice struct {
        start   unsafe.Pointer
        len     int
        cap     int
	}
)	

func printSlice(s []int)  {
    sInfo:= *(*Slice)(unsafe.Pointer(&s))
    fmt.Println(sInfo.start, sInfo.len, sInfo.cap)
}

func main() {
	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("a:", len(a), a)             // a: 8 [1 2 3 4 5 6 7 8]
	fmt.Println(unsafe.Pointer(&a))          // 0xc000078040
	s := a[:5]
	fmt.Println(unsafe.Pointer(&s))          // 0xc000042420
	fmt.Println("s:", len(s), cap(s), s)     // s: 5 8 [1 2 3 4 5]
	printSlice(s)                            // 0xc000078040 5 8
	s2 := s
	fmt.Println(unsafe.Pointer(&s2))         // 0xc000042460
	fmt.Println("s2:", len(s2), cap(s2), s2) // s2: 5 8 [1 2 3 4 5]
	printSlice(s2)                           // 0xc000078040 5 8
	s = append(s, -1, -2, -3, -4)            //
	fmt.Println("a:", len(a), a)             // a: 8 [1 2 3 4 5 6 7 8]
	fmt.Println("s:", len(s), cap(s), s)     // s: 9 16 [1 2 3 4 5 -1 -2 -3 -4]
	printSlice(s)                            // 0xc00007a080 9 16
	s2 = append(s2, 0)                       //
	fmt.Println("a:", len(a), a)             // a: 8 [1 2 3 4 5 0 7 8]
	fmt.Println("s2:", len(s2), cap(s2), s2) // s2: 6 8 [1 2 3 4 5 0]
	printSlice(s2)                           // 0xc000078040 6 8
	s2 = append(s2, -11, 12, 13, 14, 15)
	fmt.Println("a:", len(a), a)             // a: 8 [1 2 3 4 5 0 7 8]
	fmt.Println("s2:", len(s2), cap(s2), s2) // s2: 11 16 [1 2 3 4 5 0 -11 12 13 14 15]
	printSlice(s2)                           // 0xc00007a100 11 16
	s = a[3:6]
	fmt.Println("s:", len(s), cap(s), s)     // s: 3 5 [4 5 0]
	printSlice(s)                            // 0xc000078058 3 5
	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8)
	fmt.Println("s:", len(s), cap(s), s)     // s: 11 12 [4 5 0 1 2 3 4 5 6 7 8]
	printSlice(s)                            // 0xc00003a060 11 12
	s = append(s, 1, 2, 3)
	fmt.Println("s:", len(s), cap(s), s)     // s: 14 24 [4 5 0 1 2 3 4 5 6 7 8 1 2 3]
	printSlice(s)                            // 0xc000086000 14 24
	s = s[:16]
	fmt.Println("s:", len(s), cap(s), s)     // s: 16 24 [4 5 0 1 2 3 4 5 6 7 8 1 2 3 0 0]
	printSlice(s)                            // 0xc000086000 16 24
}
