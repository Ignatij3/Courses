package main

import (
    "fmt"
)

func Hanoy(n int, start, transit, finish int) {
    if n > 0 {
        Hanoy(n-1, start, finish, transit)
        fmt.Printf("%d: from %d to %d\n", n, start, finish)	
        Hanoy(n-1, transit, start, finish)
    }			
}	

func main() {
    Hanoy(4, 1, 2, 3)	
}
