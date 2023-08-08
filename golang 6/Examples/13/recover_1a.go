package main

import "fmt"

func recoverInvalidAccess() {  
    if r := recover(); r != nil {
        fmt.Println("Recovered \n\t", r)
    }
}

func invalidSliceAccess() {  
    defer recoverInvalidAccess()
    n := []int{5, 7, 4}
    fmt.Println(n[4])
    fmt.Println("normally returned from invalidSliceAccess")
}

func main() {  
    invalidSliceAccess()
    fmt.Println("normally returned from main")
}
//	Recovered
//  	     runtime error: index out of range [4] with length 3
//	normally returned from main
