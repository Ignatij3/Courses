package main

import  (
	"fmt"
	"unsafe"
)

func main() {  
    a:= [...]uint {0,1,2,3,4,5,6,7,8,9}
    fmt.Println(a, "length =", len(a), "capacity =", cap(a))
    fmt.Printf("array a. address: %p  size: %d\n", &a, unsafe.Sizeof(a))
    fmt.Printf("   a[0]. address: %p  size: %d\n", &a[0], unsafe.Sizeof(a[0]))
    fmt.Printf("   a[2]. address: %p  size: %d\n", &a[2], unsafe.Sizeof(a[2]))
    fmt.Println("   a[2] =", *(*int)(unsafe.Pointer((uintptr(unsafe.Pointer(&a)) + unsafe.Sizeof(a[0]) + unsafe.Sizeof(a[1])))))
    fmt.Printf("   a[9]. address: %p  size: %d\n", &a[9], unsafe.Sizeof(a[9]))
    b:= a
    fmt.Printf("array b. address: %p  size: %d\n", &b, unsafe.Sizeof(b))
    b[0] = 100
    fmt.Println("   a:", a)
    fmt.Println("   b:", b)
    
/*
[0 1 2 3 4 5 6 7 8 9] length = 10 capacity = 10
array a. address: 0xc0000122d0  size: 80
   a[0]. address: 0xc0000122d0  size: 8
   a[2]. address: 0xc0000122e0  size: 8
   a[2] = 2
   a[9]. address: 0xc000012318  size: 8
array b. address: 0xc000086050  size: 80
   a: [0 1 2 3 4 5 6 7 8 9]
   b: [100 1 2 3 4 5 6 7 8 9]
*/
}
