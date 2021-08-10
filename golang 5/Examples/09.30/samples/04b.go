package main

import  (
	"fmt"
	"unsafe"
)	

type (
	sliceHeader struct {
        start   uintptr
        len     int
        cap     int
	}
)	

func main() {  
    a:= [...]uint {0,1,2,3,4,5,6,7,8,9}
    b:= a[2:7]
    fmt.Println(unsafe.Pointer(&b))                    // 0xc0000044c0
    fmt.Println(unsafe.Sizeof(b))                      // 24
	barr := (*[3]int)(unsafe.Pointer(&b))
	fmt.Println(*barr)	                               // [824633795296 5 8]
    // #0
    fmt.Printf("%x %p %p\n", (*barr)[0], &b[0], &a[2]) // c0000122e0 0xc0000122e0 0xc0000122e0
    // #1
    fmt.Printf("%d %d\n", (*barr)[1], len(b))          // 5 5
    // #2
    fmt.Printf("%d %d\n", (*barr)[2], cap(b))          // 8 8
	// all together - struct
	fmt.Println(*(*sliceHeader)(unsafe.Pointer(&b)))   // {824633795296 5 8}
}
