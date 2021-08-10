package main

import (  
    "fmt"
)

func hello(quit chan bool) {  
    fmt.Println("Hello, goroutine")
    quit <- true
}
func main() {  
    waiting := make(chan bool)
    go hello(waiting)
    <-waiting
    fmt.Println("Hello from main")
}
/*
Hello, goroutine
Hello from main
*/
