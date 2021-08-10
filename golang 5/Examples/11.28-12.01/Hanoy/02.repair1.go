package main

import (
    "fmt"
)

func Hanoy(n int, start, transit, finish int) {
    if n > 0 {
        if start == 1 && finish == 3{
            Hanoy(n-1, 1, 2, 3)
            fmt.Println("1 2")
            Hanoy(n-1, 3, 2, 1)
            fmt.Println("2 3")
            Hanoy(n-1, 1, 2, 3)
        } else {
            Hanoy(n-1, start, finish, transit)
            fmt.Println(start, finish)	
            Hanoy(n-1, transit, start, finish)
        }
    }			
}	

func main() {
    Hanoy(3, 1, 2, 3)	
}
