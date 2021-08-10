//logarithm.go
package aroundpower

import  (
    "fmt"
    _ "learnpackage/power"
)

func init()  {
   fmt.Println("learnpackage/aroundpower package initializing... - logarithm")
}

func Logarithm(a int64, base uint) int64 {
    if base <= 1 || a <= 0 { return -1 }
    var  (
        res int64 = 0
        p int64 = 1
    )
    for p <= a  { 
        p *= int64(base)
        res++
    }
    return res-1
}
