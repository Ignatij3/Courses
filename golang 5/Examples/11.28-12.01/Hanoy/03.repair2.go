package main

import (
    "fmt"
)

func Hanoy(n int, start, transit, finish int) {
    if n > 0 {
        if (start == 1  && finish == 3) || (start == 3  && finish == 1) {
        // variant: if transit == 2
            Hanoy(n-1, start, transit, finish)
            fmt.Println(start, transit)
            Hanoy(n-1, finish, transit, start)
            fmt.Println(transit, finish)
            Hanoy(n-1, start, transit, finish)
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
