package main

import (  
    "fmt"
)

func main() {  
    someInt := new(int)
    fmt.Printf("someInt value is %d, type is %T, address is %v\n", *someInt, someInt, someInt)
    *someInt = 85
    fmt.Println("New someInt value is", *someInt)
/*
someInt value is 0, type is *int, address is 0xc00000e0b0
New someInt value is 85
*/
}
