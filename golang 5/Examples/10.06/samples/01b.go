package main

import  (
	"fmt"
	"unsafe"
)	

type (
	stringHeader struct {
		start   unsafe.Pointer
		length  uint
	}
)	

func getString(s string) *stringHeader {
    return (*stringHeader)(unsafe.Pointer(&s))
}

func main() {  
    str:= "ABCDEFG"
    fmt.Println(str)      // ABCDEFG
    fmt.Println(&str)     // 0xc00003c1c0
    sh:= getString(str)
    fmt.Println(*sh)      // {0x4bf578 7}
    b:= (*byte)(unsafe.Pointer(uintptr((*sh).start)+4))
    fmt.Println(&b, b)    // 0xc000070020 0x4bf57c
    *b = '+'
    // unexpected fault address 0x4bf57c
    // fatal error: fault
    // ...
    fmt.Println(str)                             
}
