package main

import "fmt"

func main() {

    switch i := 10; i % 2 {
    case 0:
        fmt.Println(i, "is even integer")
    case 1:
        fmt.Println(i, "is odd integer")
    default:
        fmt.Println("Wow! Amazing", i)
    }

}
