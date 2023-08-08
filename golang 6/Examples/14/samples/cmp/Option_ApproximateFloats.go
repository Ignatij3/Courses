// Approximate equality for floats can be handled by defining a 
// custom comparer on floats that determines two values to be 
// equal if they are within some range of each other.

// This example is for demonstrative purposes; 
// use cmpopts.EquateApprox instead.

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
		delta := math.Abs(x - y)
		mean := math.Abs(x+y) / 2.0
		return delta/(1.0 + mean) < 0.00001
	})

	x := []float64{1.0, 1.1, 1.2, math.Pi}
	y := []float64{1.0, 1.1, 1.2, 3.14159265359} // Accurate enough to Pi
	z := []float64{1.0, 1.1, 1.2, 3.1415}        // Diverges too far from Pi

	fmt.Println(cmp.Equal(x, y, opt))
	fmt.Println(cmp.Equal(y, z, opt))
	fmt.Println(cmp.Equal(z, x, opt))

}
// true
// false
// false
