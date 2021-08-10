package main

import  (
	"fmt"
	"unsafe"
)	

func main() {  
    a:= [...]uint {0,1,2,3,4,5,6,7,8,9}
    b:= a[2:7]
    fmt.Println(b, len(b), cap(b))                // [2 3 4 5 6] 5 8
    fmt.Printf("%d %x %x\n", b[0], &b[0], &a[2])  // 2 c000122100 c000122100
    fmt.Printf("%d %x %x\n", b[1], &b[1], &a[3])  // 3 c000122108 c000122108
    fmt.Printf("%d %x %x\n", b[2], &b[2], &a[4])  // 4 c000122110 c000122110
    a[5] = 97
    fmt.Println(b)                                // [2 3 4 97 6]
    fmt.Println(&b)                               // &[2 3 4 97 6]
    fmt.Printf("%p\n", &b)                        // 0xc0000044c0
    fmt.Println(unsafe.Pointer(&b))               // 0xc0000044c0
}
