// Two slices may be considered equal if they have the same elements, 
// regardless of the order that they appear in. 
// Transformations can be used to sort the slice.

//This example is for demonstrative purposes; use cmpopts.SortSlices instead. 
package main

import (
	"fmt"
	"sort"

	"github.com/google/go-cmp/cmp"
)

func main() {
	// This Transformer sorts a []int.
	trans := cmp.Transformer("Sort", func(in []int) []int {
		out := append([]int(nil), in...) // Copy input to avoid mutating it
		sort.Ints(out)
		return out
	})

	x := struct{ Ints []int }{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}}
	y := struct{ Ints []int }{[]int{2, 8, 0, 9, 6, 1, 4, 7, 3, 5}}
	z := struct{ Ints []int }{[]int{0, 0, 1, 2, 3, 4, 5, 6, 7, 8}}

	fmt.Println(cmp.Equal(x, y, trans))
	fmt.Println(cmp.Equal(y, z, trans))
	fmt.Println(cmp.Equal(z, x, trans))

}

// true
// false
// false
