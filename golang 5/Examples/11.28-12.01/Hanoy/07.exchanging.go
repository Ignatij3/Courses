package main

import (
    "fmt"
)

func Hanoy(n int, start, transit, finish int) {
    if n == 1 {
        fmt.Println("  1:", start, finish)	
    }		
    if n > 1 {
        Hanoy(n-1, start, transit, finish)
        Hanoy(n-2, finish, start, transit)
        fmt.Println("swap", start, finish)	
        Hanoy(n-2, transit, finish, start)
        Hanoy(n-1, start, transit, finish)
    }			
}	

func main() {
    Hanoy(4, 1, 2, 3)	
}
