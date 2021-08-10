package main

import (
	"fmt"
	"time"
	"math/rand"
)

type data struct {
	x float64
	y float64
}

func Fill(a *[20]data) {
	rand.Seed(time.Now().UnixNano())
	for k := 0; k < 20; k++ {
		(*a)[k] = data{float64(k), float64(rand.Intn(200) + 1)}
	}
}

/*func Test(a [20]data) {
	var (
		numer, denom, frac float64 = 1, 1, 0
		result [20]float64
	)
	
	for x := 0; x < 20; x++ {
		for p := 0; p < 20; p++ {
			for k := 0; k < 20; k++ {
				if k != p {
					numer *= float64(x) - a[k].x
					denom *= a[p].x - a[k].x
				}
			}
			frac = numer/denom
			result[x] += a[p].y * frac
			numer, denom = 1, 1
		}
		fmt.Printf("y[%d] = %g\n", x, result[x])
	}
}*/

func CalcX(a [20]data, k int, x float64) float64 {
	res := 1.0
	for p := 0; p < k; p++ {res *= x - a[p].x}
	return res
}

func GetC(a [20]data, k int, x float64) float64 {
	return (a[k].y - GetQ(a, k-1, a[k].x)) / CalcX(a, k, a[k].x)
}

func GetQ(a [20]data, k int, x float64) float64 {
	if k > 0 {return GetQ(a, k-1, x) + GetC(a, k, x) * CalcX(a, k, x)}
	return a[0].y
}

func main() {
	var a [20]data
	Fill(&a)
	//Test(a)
	for p, n := range a {fmt.Printf("data[%d] - %v\n", p, n)}
	fmt.Printf("\nInterpolation of 20 = %g\n", GetQ(a, 19, 20.0))
	fmt.Printf("\nInterpolation of 20 = %f\n", GetQ(a, 19, 20.0))
}
