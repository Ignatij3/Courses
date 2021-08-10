package main 

import  (
    "fmt"
    "learnpackage/power"
    "learnpackage/aroundpower"
)

var title = "main.init() function processing..."

func init() {
  fmt.Println(title)
}

func main() {
    var (
        a int64 = 3
        n uint = 5
    )    
    p:= power.Power(a, n)
    fmt.Println(a, "^", n, "=", p)
    fmt.Println(power.Two, "^ 5 =", power.Power(power.Two, 5))
    fmt.Println("cubic root of 250 = ", aroundpower.Root(250, 3))
    // 6^3 = 216 < 250 < 7^3 = 343
    fmt.Println("fifth degree root of 50000 = ", aroundpower.Root(50000, 5))
    // 8^5 = 32768 < 50000 < 9^5 = 59049
    fmt.Print("the logarithm of 250 to base 3 = ")
    fmt.Println(aroundpower.Logarithm(250, 3))
    // 3^5 = 243 < 250 < 3^6 = 729
}
