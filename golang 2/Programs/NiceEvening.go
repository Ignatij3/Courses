package main
import (
		"fmt"
)

func xor(a, b bool) bool { return a != b }
func impl(a, b bool) bool { return !a || b }
func equiv(a, b bool) bool { return a == b }

func Evening01(a, b, c, d bool) bool {return impl(a, xor(b, c)) && impl(d, xor(a, b)) && impl(c, !b) && equiv(d, !c) && impl(b, c) }

func main() {
	BArr  := [2]bool {false, true}
	control := true
	
	fmt.Println("      A      B      C      D      (A => (B ^ C)) & (D => (A ^ B)) & (C => !B) & (D <=> !C) & (B => C)")
	
	for _, a := range BArr {
		for _, b := range BArr {
			for _, c := range BArr {
				for _, d := range BArr {
					fmt.Printf("%7v%7v%7v%7v%73v\n", a, b, c, d, Evening01(a, b, c, d))
					control = control && Evening01(a, b, c, d)
				}
			}
		}
	}
	if control {
		fmt.Println("Данные верны")
	} else {
		fmt.Println("Данные ошибочны")
	}
}
/*
 * (A => (B ^ C)) & (D => (A ^ B)) & (C => !B) & (D <=> !C) & (B => C)
 */
