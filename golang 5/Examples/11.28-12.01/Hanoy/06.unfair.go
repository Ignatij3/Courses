package main

import (
    "fmt"
)

func Hanoy(n int, start, finish int) {
    if n == 1 {
        fmt.Println(start, finish)	
    }		
    if n > 1 {
        Hanoy(n-1, start, finish)
        fmt.Println(start, 2)	
        Hanoy(n-1, finish, start)
        fmt.Println(2, finish)	
        Hanoy(n-1, start, finish)
    }			
}	

func main() {
    Hanoy(2, 1, 3)	
}
