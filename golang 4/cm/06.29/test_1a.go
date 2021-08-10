package main

import (  
    "fmt"
)

func hello() {  
    fmt.Println("Hello, goroutine")
}
func main() {  
    go hello()
    fmt.Println("Hello from main")
}

// Hello from main
