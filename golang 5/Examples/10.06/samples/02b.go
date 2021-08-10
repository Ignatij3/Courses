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
    a:= [...]uint {0,1,2,3,4,5,6}
    fmt.Println(unsafe.Pointer(&a))     // 0xc000078040
    b:= a[2:5]
    fmt.Println(b)                      // [2 3 4]
    fmt.Println(unsafe.Pointer(&a[2]))  // 0xc000078050
	printSlice(b)                       // 0xc000078050 3 5
	c:= b[1:2]
    fmt.Println(c)                      // [3]
	printSlice(c)                       // 0xc000078058 1 4
}
