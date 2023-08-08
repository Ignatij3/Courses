// To have floating-point comparisons combine both properties of NaN 
// being equal to itself and also approximate equality of values, 
// filters are needed to restrict the scope of the comparison so that 
// they are composable.

// This example is for demonstrative purposes; use cmpopts.EquateNaNs and cmpopts.EquateApprox instead.
package main

import (
	"fmt"
	"math"

	"github.com/google/go-cmp/cmp"
)

func main() {
	alwaysEqual := cmp.Comparer(func(_, _ interface{}) bool { return true })

	opts := cmp.Options{
		// This option declares that a float64 comparison is equal only if
		// both inputs are NaN.
		cmp.FilterValues(func(x, y float64) bool {
			return math.IsNaN(x) && math.IsNaN(y)
		}, alwaysEqual),

		// This option declares approximate equality on float64s only if
		// both inputs are not NaN.
		cmp.FilterValues(func(x, y float64) bool {
			return !math.IsNaN(x) && !math.IsNaN(y)
		}, cmp.Comparer(func(x, y float64) bool {
			delta := math.Abs(x - y)
			mean := math.Abs(x+y) / 2.0
			return delta/mean < 0.00001
		})),
	}

	x := []float64{math.NaN(), 1.0, 1.1, 1.2, math.Pi}
	y := []float64{math.NaN(), 1.0, 1.1, 1.2, 3.14159265359} // Accurate enough to Pi
	z := []float64{math.NaN(), 1.0, 1.1, 1.2, 3.1415}        // Diverges too far from Pi

	fmt.Println(cmp.Equal(x, y, opts))
	fmt.Println(cmp.Equal(y, z, opts))
	fmt.Println(cmp.Equal(z, x, opts))

}
// true
// false
// false
