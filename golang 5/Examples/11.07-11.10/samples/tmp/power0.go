//power.go
package power

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
