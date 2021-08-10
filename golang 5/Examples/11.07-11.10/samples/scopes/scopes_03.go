package main

import "fmt"

func main() {
    type line struct {
        name string
        next *line
    }

    x := line{name: "Max", next: &line{name: "Alex", next: nil}}
    fmt.Println(x.name)           // Max
    fmt.Println(x.next.name)      // Alex
    fmt.Println(x.next.next)      // <nil> 
}
