package main

import (  
    "fmt"
)

func main() {  
    b := 9999
    a := &b
    fmt.Println("address of b is", a)
    fmt.Println("value of b is", *a)
    *a++
    fmt.Println("new value of b is", b)
/*
address of b is 0xc00000e0b0
value of b is 9999
new value of b is 10000
*/ 
}
