package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, D, x1, x2 float64
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	D = b*b - 4*a*c
	fmt.Println("D =", b, "^2 - 4 *", a, "*", c, "=", D)
	D = math.Sqrt(D)
	x1 = (-b - D) / (2 * a)
	x2 = (-b + D) / (2 * a)
	fmt.Println("x1 = (", -b, "-", D, ")/(", 2, "*", a, ")=", x1)
	fmt.Println("x2 = (", -b, "+", D, ")/(", 2, "*", a, ")=", x2)
}
