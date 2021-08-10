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

func printSlice(s []uint)  {
    ps:= (*Slice)(unsafe.Pointer(&s))
    fmt.Println((*ps).start, (*ps).len, (*ps).cap)
}

func main() {  
    var p []uint
    fmt.Println(p)						// []
    printSlice(p)                       // <nil> 0 0
    p = append(p, 12345)
    fmt.Println(p)
    printSlice(p)                       // 0xc000048090 2 2
    p = make([]uint, 2)
    fmt.Println(p)						// [0 0]
    printSlice(p)                       // 0xc000048090 2 2
}
