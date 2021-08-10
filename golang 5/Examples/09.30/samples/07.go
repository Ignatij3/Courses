package main

import  (
	"fmt"
	"unsafe"
)	

type (
	stringHeader struct {
		start   uintptr
		length  uint
	}
)	

func main() {  
    str:= "ABCD"
    str += ".abcd" 
	fmt.Println(str)                                    // ABCD.abcd
    fmt.Println(&str)                                   // 0xc00002c1f0
    fmt.Println(unsafe.Pointer(&str))                   // 0xc00002c1f0
    fmt.Println(unsafe.Sizeof(str))                     // 16
	strarr := (*[2]uintptr)(unsafe.Pointer(&str))       // &[824633778352 9]
	fmt.Println(*strarr)                                // c00000e0b0 9
	fmt.Printf("%x %d\n", (*strarr)[0], (*strarr)[1])
	*(*byte)(unsafe.Pointer((*strarr)[0]+4)) = '+'
	fmt.Println(str)                                    // ABCD+abcd
	// all together - struct
	fmt.Println(*(*stringHeader)(unsafe.Pointer(&str))) // {824633778352 9}
	str2:= str
	fmt.Println(*(*stringHeader)(unsafe.Pointer(&str2)))// {824633778352 9}
}
