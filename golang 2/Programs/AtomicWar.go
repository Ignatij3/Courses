package main
import (
		"fmt"
)

func xor(a, b bool) bool { return a != b }
func impl(a, b bool) bool { return !a || b }

func Shelter01(a, b, c, d, e, f bool) bool {return impl(a, b && c) && impl(b, d) && impl(c, e) && impl(!a, f) }
func Shelter02(d, e, f bool) bool {return  xor(d && e, f) }

func main() {
	BArr  := [2]bool {false, true}
	control := true
	
	fmt.Println("      A      B      C      D      E      F      (A => (B & C)) & (B => D) & (C => E) & (!A => F)      (D & E) ^ F")
	
	for _, a := range BArr {
		for _, b := range BArr {
			for _, c := range BArr {
				for _, d := range BArr {
					for _, e := range BArr {
						for _, f := range BArr {
							fmt.Printf("%7v%7v%7v%7v%7v%7v%36v%17v\n", a, b, c, d, e, f, Shelter01(a, b, c, d, e, f), Shelter02(d, e, f))
							control = control && impl(Shelter02(d, e, f), Shelter01(a, b, c, d, e, f))
						}
					}
				}
			}
		}
	}
	if control {
		fmt.Println("Выводы верны")
	} else {
		fmt.Println("Выводы ошибочны")
	}
}
/* A => (B & C)
 * B => D
 * C => E
 * !A => F
 * 
impl(A, B & C) & impl(B, D) & impl(C, E) & impl(!A, F)
(A => (B & C)) & (B => D) & (C => E) & (!A => F)


 * (A & B & C & D & E) | (!A & F) - 36
 * 
 * (D & E) ^ F - 11
 */
