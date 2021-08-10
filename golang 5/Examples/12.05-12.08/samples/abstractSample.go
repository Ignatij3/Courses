package main

import "fmt"

func P(n byte) {
    if n>0 {
        P(n / 2)
        fmt.Printf("%d.", n)
        P(n / 3)
    }
}

func main() {
    P(20)   
}
