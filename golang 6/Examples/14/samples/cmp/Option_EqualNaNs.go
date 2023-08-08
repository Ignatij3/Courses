// Normal floating-point arithmetic defines == to be false when comparing 
// NaN with itself. In certain cases, this is not the desired property.

// This example is for demonstrative purposes; use cmpopts.EquateNaNs instead.

package main

import (
	"fmt"
	"math"

	"github.com/google/go-cmp/cmp"
)

func main() {
	// This Comparer only operates on float64.
	// To handle float32s, either define a similar function for that type
	// or use a Transformer to convert float32s into float64s.
	opt := cmp.Comparer(func(x, y float64) bool {
		return (math.IsNaN(x) && math.IsNaN(y)) || x == y
	})

	x := []float64{1.0, math.NaN(), math.E, -0.0, +0.0}
	y := []float64{1.0, math.NaN(), math.E, -0.0, +0.0}
	z := []float64{1.0, math.NaN(), math.Pi, -0.0, +0.0} // Pi constant instead of E

	fmt.Println(cmp.Equal(x, y, opt))
	fmt.Println(cmp.Equal(y, z, opt))
	fmt.Println(cmp.Equal(z, x, opt))

}
// true
// false
// false

