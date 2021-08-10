package main

import (
    "fmt"
)

func Hanoy(n int, start, transit, finish int) {
    if n > 0 {
        Hanoy(n-1, start, finish, transit)
        fmt.Println(start, finish)	
        Hanoy(n-1, transit, start, finish)
    }			
}	

func Hanoy2(n int, start int) {
    if n > 0 {
        switch {
        case start == 1 && n%2 == 0:
            Hanoy(n-1, 1, 2, 3)
            fmt.Println(1, 2)
            Hanoy2(n-1, 3)
        case start == 1 && n%2 == 1:
            Hanoy(n-1, 1, 3, 2)
            fmt.Println(1, 3)
            Hanoy2(n-1, 2)
        case start == 2  && n % 2 == 0:
            Hanoy2(n-1, 2)
        case start == 2  && n % 2 == 1:
            Hanoy(n-1, 2, 3, 1)
            fmt.Println(2, 3)
            Hanoy2(n-1, 1)
        case start == 3  && n % 2 == 0:
            Hanoy(n-1, 3, 2, 1)
            fmt.Println(3, 2)
            Hanoy2(n-1, 1)
        case start == 3  && n % 2 == 1:
            Hanoy2(n-1, 3)
        }
        
    }			
}	

func main() {
    Hanoy2(4, 1)	
}
