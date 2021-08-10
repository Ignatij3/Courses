package main

import "fmt"

func main() {
    {
        v := 1
        {
            fmt.Println(v)        // 1
        }
        fmt.Println(v)
    }
    //  fmt.Println(v)
    //             compilation failed - undefined v
}
