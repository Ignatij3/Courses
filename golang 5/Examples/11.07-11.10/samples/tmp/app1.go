package main

import  (
    "fmt"
    "learnpackage/power"
)

func main() {
    var (
        a int64 = 3
        n uint = 5
    )
    p:= power.Power(a, n)
    fmt.Println(a, "^", n, "=", p)
}
