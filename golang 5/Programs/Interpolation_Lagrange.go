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

func Test(a [20]data) {
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
}

func Calculate(a *[20]data) float64 {
	var result float64
	for n := 0; n < 20; n++ {
		result += (*a)[n].y * PnCalc((*a), n)
	}
	return result
}

func PnCalc(a [20]data, num int) float64 {
	var (
		numer, denom, res float64 = 1, 1, 0
	)
	
	for k := 0; k < 20; k++ {
		if k != num {
			numer *= 20 - a[k].x
			denom *= a[num].x - a[k].x
		}
	}
	
	res = numer/denom
	return res
}

func main() {
	var a [20]data
	Fill(&a)
	Test(a)
	for p, n := range a {fmt.Printf("data[%d] - %v\n", p, n)}
	fmt.Printf("\nInterpolation of 20 = %g\n", Calculate(&a))
	fmt.Printf("\nInterpolation of 20 = %f\n", Calculate(&a))
}
