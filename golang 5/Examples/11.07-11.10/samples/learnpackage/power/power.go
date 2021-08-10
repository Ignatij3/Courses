//power.go
package power

const (
  Two = 2
  three = 3
)

func Power(a int64, n uint) int64 {
    var (
        i uint
        res int64
    )    
    for i, res = 0, 1; i < n; i++ {
        res *= a
    }
    return res
}
