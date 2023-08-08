package square
import (
	"math"
	"sort"
)	

// Returns the square of x 
func Square(x int) int {
	return x*x
}

// If n is perfect square, then PerfectSquare1 returns
// true and value of square root of n, else PerfectSquare1
// returns false and some undefined integer
func PerfectSquare1 (n int) (ok bool, sqrt int) {
	sqrt = int(math.Round(math.Sqrt(float64(n))))
	ok = Square(sqrt) == n
	return
}	

// If n is perfect square, then PerfectSquare2 returns
// true and value of square root of n, else PerfectSquare2
// returns false and some undefined integer
func PerfectSquare2 (n int) (ok bool, sqrt int) {
	sum, delta := 0, 1
	for sum < n {
		// k^2 - (k-1)^2 = 2*k - 1
		sum += delta
		delta += 2
	}
	ok, sqrt = sum == n, delta/2
	return
}	

// If n is perfect square, then PerfectSquare3 returns
// true and value of square root of n, else PerfectSquare3
// returns false and some undefined integer
func PerfectSquare3 (n int) (ok bool, sqrt int) {
	c := func(i int) bool {
		return i*i - n >= 0
	}
	sqrt = sort.Search(n+1, c)
	ok = Square(sqrt) == n
	return
}
