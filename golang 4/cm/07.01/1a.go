package main
import  (
	"fmt"
	"time"
)	

func firstDivisor(n int64) int64 {
	// Возвращает наименьший собственный 
	// (отличающийся от 1 и n) делитель числа n. 
	// Если n - простое число, то возвращает 0.
	var i int64
    for i = 2; i*i <= n; i++ {
        if n%i == 0 {
            return i
        } 
    }
    return 0
}

func main() {
    var n int64
    n = 1000000087*1120000093
	t0 := time.Now()
	fmt.Println( firstDivisor(n) )
	t1 := time.Now()
	fmt.Printf("The call took %v to run.\n", t1.Sub(t0))
}
