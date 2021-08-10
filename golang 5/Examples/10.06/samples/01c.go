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
    fmt.Println(*sh)      // {0x4bf579 7}

    str1:= "ABCDEFG"
    fmt.Println(str1)     // ABCDEFG
    fmt.Println(&str1)    // 0xc00003c1c0
    sh1:= getString(str1) 
    fmt.Println(*sh1)     // {0x4bf579 7}
    
    str1 += "!"
    fmt.Println(str1)     // ABCDEFG!
    fmt.Println(&str1)    // 0xc00003c1f0
    sh1 = getString(str1) 
    fmt.Println(*sh1)     // {0xc0000480b0 8}
    b:= (*byte)(unsafe.Pointer(uintptr((*sh1).start)+4))
    fmt.Println(&b, b)    // 0xc000070020 0xc0000480b4
    *b = '+'
	fmt.Println(str1)     // ABCD+FG!                    
}
