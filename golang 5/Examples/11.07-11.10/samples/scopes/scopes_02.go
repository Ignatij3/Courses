package main

import "fmt"

func main() {
    { // start outer block
        a := 1
        fmt.Println(a)         // 1

        { // start inner block
            b := 2
            fmt.Println(a, b)  // 1 2
        } // end inner block
        fmt.Println(a)         // 1
        // fmt.Println(b)       - undefined: b

    } // end outer block
    // fmt.Println(a)       - undefined: a
}
