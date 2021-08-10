package main

import (
	"rational/arithmetics"
	"rational/change"
	"math/rand"
	"time"
	"fmt"
)

func main() {
	var (
		x, backup *arithmetics.Fraction
		y arithmetics.Fraction
		x2, backup2 *change.Fraction
		y2 change.Fraction
		result int
	)
	
	rand.Seed(time.Now().UnixNano())
	x = &arithmetics.Fraction{rand.Intn(1000), rand.Intn(1000)}
	backup = &arithmetics.Fraction{(*x).A, (*x).B}
	y = arithmetics.Fraction{rand.Intn(1000), rand.Intn(1000)}
	
	x2 = &change.Fraction{(*x).A, (*x).B}
	backup2 = &change.Fraction{(*x).A, (*x).B}
	y2 = change.Fraction{y.A, y.B}
	
	fmt.Printf("Fraction 1 - %d/%d\nFraction 2 - %d/%d\n\n", (*x).A, (*x).B, y.A, y.B)
	
	(*backup) = (*x)
	fmt.Printf("%d/%d + %d/%d = ", (*x).A, (*x).B, y.A, y.B)
	x.Inc(y)
	fmt.Printf("%d/%d\n", (*x).A, (*x).B)
	
	(*x) = (*backup)
	fmt.Printf("%d/%d - %d/%d = ", (*x).A, (*x).B, y.A, y.B)
	x.Dec(y)
	fmt.Printf("%d/%d\n", (*x).A, (*x).B)
	
	(*x) = (*backup)
	fmt.Printf("%d/%d * %d/%d = ", (*x).A, (*x).B, y.A, y.B)
	x.Mult(y)
	fmt.Printf("%d/%d\n", (*x).A, (*x).B)
	
	(*x) = (*backup)
	fmt.Printf("%d/%d / %d/%d = ", (*x).A, (*x).B, y.A, y.B)
	x.Div(y)
	fmt.Printf("%d/%d\n\n", (*x).A, (*x).B)
	
	(*x2) = (*backup2)
	result = x2.Compare(y2)
	fmt.Printf("Result of compare - %d\n", result)
	
	(*x2) = (*backup2)
	fmt.Printf("Fraction reduced from %d/%d ", (*x2).A, (*x2).B)
	x2.Reduce()
	fmt.Printf("to %d/%d\n", (*x2).A, (*x2).B)
	
	(*x2) = (*backup2)
	fmt.Printf("Fraction %d/%d ", (*x2).A, (*x2).B)
	result = x2.Round()
	fmt.Printf("rounded to %d\n", result)
	
	(*x2) = (*backup2)
	fmt.Printf("Fraction %d/%d ", (*x2).A, (*x2).B)
	result = x2.Floor()
	fmt.Printf("rounded to %d (lower value)\n", result)
	
	(*x2) = (*backup2)
	fmt.Printf("Fraction %d/%d ", (*x2).A, (*x2).B)
	result = x2.Ceil()
	fmt.Printf("rounded to %d (higher value)\n", result)
	
	(*x2) = (*backup2)
	fmt.Printf("Fraction %d/%d ", (*x2).A, (*x2).B)
	(*x2) = x2.Frac()
	fmt.Printf("decreased to %d/%d\n", (*x2).A, (*x2).B)
}
