package main
import "fmt"

type smallset = uint64

func emptyset() smallset {
	return 0
}
func isempty(s smallset) bool {
	return s == 0
}
func include(s smallset, n uint8) smallset {
	return s | (1 << n)
}
func exclude(s smallset, n uint8) smallset {
	return s | (0 << n)
}
func union(a, b smallset) smallset {
	return a ^ b
}
func intersect(a, b smallset) smallset {
	return a & b
}
func difference(a, b smallset) smallset {
	return a &^ b
}
func complement(s smallset) smallset {
	return ^s
}
func disjoint(a, b smallset) bool {
	return isempty(intersect(a, b))
}
func subset(a, b smallset) bool {
	return isempty(intersect(a, b))
}
func belongs(a smallset, n uint8) bool {
	return a & (1 << n) != 0
}

/*
func belongs(a smallset, n uint8) bool {
	var b smallset = emptyset()
	b = include(b, n)
	return !isempty(union(a, b))
}
*/
func show(s smallset) {
	var i uint8 = 0
	for i < 64 {
		if belongs(s, i) {
			fmt.Print(i, ", ")
		}
		i++
	}
	fmt.Println()
}
func size(s smallset) int {
	count := 0
	var i uint8 = 0
	for i < 64 {
		if belongs(s, i) { count++ }
		i++
	}
	return count
}

func main() {
	var a smallset = emptyset()
	a = include(a, 4)
	a = include(a, 7)
	a = include(a, 0)
	a = include(a, 6)
	show(a)
}
/*
 * if s & (1 << 4) != 0 {...}
 */
