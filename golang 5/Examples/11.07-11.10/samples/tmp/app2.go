package main 

import  (
    "fmt"
    "learnpackage/power"
)

func main() {
    fmt.Println("Power calculation")
    var (
        a int64 = 3
        n uint = 5
    )
    p:= power.Power(a, n)
    fmt.Println(a, "^", n, "=", p)
    fmt.Println(power.Power(power.Two, 5))
    fmt.Println(power.Power(power.Three, 2))
    fmt.Println(power.Power(power.three, 2))
}