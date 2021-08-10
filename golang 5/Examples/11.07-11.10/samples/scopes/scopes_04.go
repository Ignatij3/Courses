package main

import "fmt"

func main() {

    for i := 1; i <= 5; i++ {
        fmt.Print(i)          // 12345
    }
    fmt.Println()
    //  fmt.Println(i)    - undefined: i

    var i int                 // no redeclaring 
    for i = 1; i <= 5; i++ {
        fmt.Print(i)          // 12345
    }
    /// var i int         -   i redeclared in this block
    fmt.Println()
    fmt.Println(i)            // 6

}
