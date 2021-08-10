//root.go
package aroundpower

import (
    "fmt"
    "learnpackage/power"
)

func init()  {
   fmt.Println("learnpackage/aroundpower package initializing... - root")
}

func Root(a int64, n uint) int64 {
    var i int64
    for i = 0; power.Power(i, n) <= a; i++ {  }
    return i-1
}
