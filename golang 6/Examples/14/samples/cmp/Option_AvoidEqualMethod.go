// If the Equal method defined on a type is not suitable, 
// the type can be dynamically transformed to be stripped 
// of the Equal method (or any method for that matter).
package main

import (
	"fmt"
	"strings"

	"github.com/google/go-cmp/cmp"
)

type otherString string

func (x otherString) Equal(y otherString) bool {
	return strings.ToLower(string(x)) == strings.ToLower(string(y))
}

func main() {
	// Suppose otherString.Equal performs a case-insensitive equality,
	// which is too loose for our needs.
	// We can avoid the methods of otherString by declaring a new type.
	type myString otherString

	// This transformer converts otherString to myString, allowing Equal to use
	// other Options to determine equality.
	trans := cmp.Transformer("", func(in otherString) myString {
		return myString(in)
	})

	x := []otherString{"foo", "bar", "baz"}
	y := []otherString{"fOO", "bAr", "Baz"} // Same as before, but with different case

	fmt.Println(cmp.Equal(x, y))        // Equal because of case-insensitivity
	fmt.Println(cmp.Equal(x, y, trans)) // Not equal because of more exact equality

}
// true
// false

