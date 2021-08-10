package main

import (
    "fmt"
)

func TurnOff (n int) {
    if n > 1 {
        TurnOff(n-2)
        fmt.Println(n)	
        TurnOn(n-2)
        TurnOff(n-1)
    } else 
    if n == 1 {
        fmt.Println(1)	
    }	
	
}
	
func TurnOn (n int) {
    if n > 1 {
        TurnOn(n-1)
        TurnOff(n-2)
        fmt.Println(n)	
        TurnOn(n-2)
    } else 
    if n == 1 {
        fmt.Println(1)	
    }	
}	

func main() {
    TurnOn(6)	
}
