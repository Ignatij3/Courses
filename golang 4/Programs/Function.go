package main

import (
	"fmt"
	"math"
)

var (
	iFunc int
)

func f(x float64) float64 {
	if iFunc == 0 {
		return x * x * x - 2 * x - 5
	} else if iFunc == 1 {
		return x * x * x - 6 * x - 20
	} else if iFunc == 2 {
		return 0.004 * x * x * x + 0.07 * x * x + 0.3 * x - 1.47
	} else if iFunc == 3 {
		return x - 5 * math.Sin(x) - 3.5
	} else if iFunc == 4 {
		return x * x / 5 + math.Sin(2 * x) - 7
	}
	return 0.0
}

func Arrange(left, right, epsilon float64) {
	var a, b, c, resA, resB float64
	epsilon, a, b = 1 / epsilon, left, left + 0.01
	
	for {
		resA = f(a)
		resB = f(b)
		if (resA < 0 && resB > 0) || (resA > 0 && resB < 0) {
			c = (a + b) / 2
			DivideInHalf(a, b, c, epsilon)
			c = (a * f(b) - b * f(a)) / (f(b) - f(a))
			Chord(a, b, c, epsilon)
		}
		
		if b >= right {break}
		a += 0.01
		b = a + 0.01
	}
}

func DivideInHalf(a, b, c, epsilon float64) float64 {
	var lastRes float64
	
	y := f(c)
	res := f(a)
	if math.Abs(y) > epsilon {
		if (y < 0 && res > 0) || (y > 0 && res < 0) {
			DivideInHalf(a, c, (a + c) / 2, epsilon)//Left zero
		} else if (y < 0 && res < 0) || (y > 0 && res > 0) {
			DivideInHalf(c, b, (b + c) / 2, epsilon)//Right zero
		}
	} else if math.Abs(y) <= epsilon {
			lastRes = y
			fmt.Println("\nX -", c, "\nY -", lastRes)
		}
	return y
}

func Chord(a, b, c, epsilon float64) float64 {
	var lastRes float64
	
	y := f(c)
	if math.Abs(y) > epsilon {
		c = (a * f(b) - b * f(a)) / (f(b) - f(a))
		Chord(c, b, c, epsilon)
	} else if math.Abs(y) <= epsilon {
			lastRes = y
			fmt.Println("\nX -", c, "\nY -", lastRes, "Chord")
		}
	return y
}

func main() {
	var epsilon, left, right float64
	
	fmt.Println("0) x^3 - 2x - 5")
	fmt.Println("1) x^3 - 6x - 20")
	fmt.Println("2) 0.004x^3 + 0.07x^2 + 0.3x - 1.47")
	fmt.Println("3) x - 5sin(x) - 3.5")
	fmt.Println("4) x^2 / 5 + sin(2x) - 7\n")
	
	fmt.Print("Which function do you want to see? ")
	fmt.Scan(&iFunc)
	for iFunc < 0 || iFunc > 4 {
		fmt.Print("Try again: ")
		fmt.Scan(&iFunc)
	}
	
	fmt.Println()
	fmt.Println("Enter the range on which the roots should be found")
	fmt.Print("Left - ")
	fmt.Scan(&left)
	for left < -1000000000|| left > 1000000000 {
		fmt.Print("Try again: ")
		fmt.Scan(&left)
	}
	
	fmt.Println()
	fmt.Print("Right - ")
	fmt.Scan(&right)
	for right < -1000000000|| right > 1000000000 {
		fmt.Print("Try again: ")
		fmt.Scan(&right)
	}
	
	fmt.Println()
	fmt.Print("Enter how accurate should calculations be: 1/")
	fmt.Scan(&epsilon)
	for epsilon <= 0 || epsilon > 1000000000 {
		fmt.Print("Try again: 1/")
		fmt.Scan(&epsilon)
	}
	Arrange(left, right, epsilon)
}
/*
X - 2.094551481306553
Y - -2.6315731815884646e-09

X - 2.094551480699449
Y - -9.40772260094036e-09 Chord
*/
