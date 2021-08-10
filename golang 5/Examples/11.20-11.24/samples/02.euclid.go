package main

import "fmt"

func gcd1(a, b uint64) uint64  {
    if a == b  { return a }
    if a > b  { return gcd1(a-b, b) }
    /*if a < b */ return gcd1(a, b-a)
}

func gcd2(a, b uint64) uint64  {
    if b==0  {return a}
    return gcd2(b, a%b)
}

func main()  {
    var a, b uint64
    fmt.Print("Enter first number:  ")
    fmt.Scan(&a)
    fmt.Print("Enter second number:  ")
    fmt.Scan(&b)
    fmt.Println ( gcd1(a, b) )
    fmt.Println ( gcd2(a, b) )
}
