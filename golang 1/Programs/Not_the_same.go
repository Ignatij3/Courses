package main

import "fmt"

func main() {
	var a, b, c, d, e float64
	fmt.Println("Какое число среднее?")
	fmt.Print("А=")
	fmt.Scan(&a)
	fmt.Print("B=")
	fmt.Scan(&b)
	fmt.Print("C=")
	fmt.Scan(&c)
	fmt.Print("D=")
	fmt.Scan(&d)
	fmt.Print("E=")
	fmt.Scan(&e)
	if b <= a && c <= a && a <= d && a <= e ||
		d <= a && c <= a && a <= b && a <= e ||
		e <= a && c <= a && a <= b && a <= d ||
		b <= a && d <= a && a <= c && a <= e ||
		b <= a && e <= a && a <= d && a <= c ||
		e <= a && d <= a && a <= b && a <= c {
		fmt.Println("Среднее-", a)

	}
	if a <= b && c <= b && b <= d && b <= e ||
		d <= b && c <= b && b <= a && b <= e ||
		e <= b && c <= b && b <= a && b <= d ||
		a <= b && d <= b && b <= c && b <= e ||
		a <= b && e <= b && b <= d && b <= c ||
		e <= b && d <= b && b <= a && b <= c {
		fmt.Println("Среднее-", b)

	}
	if b <= c && a <= c && c <= d && c <= e ||
		d <= c && a <= c && c <= b && c <= e ||
		e <= c && a <= c && c <= b && c <= d ||
		b <= c && d <= c && c <= a && c <= e ||
		b <= c && e <= c && c <= d && c <= a ||
		e <= c && d <= c && c <= b && c <= a {
		fmt.Println("Среднее-", c)

	}
	if b <= d && c <= d && d <= a && d <= e ||
		a <= d && c <= d && d <= b && d <= e ||
		e <= d && c <= d && d <= b && d <= a ||
		b <= d && a <= d && d <= c && d <= e ||
		b <= d && e <= d && d <= a && d <= c ||
		e <= d && a <= d && d <= b && d <= c {
		fmt.Println("Среднее-", d)

	}
	if b <= e && c <= e && e <= d && e <= a ||
		d <= e && c <= e && e <= b && e <= a ||
		a <= e && c <= e && e <= b && e <= d ||
		b <= e && d <= e && e <= c && e <= a ||
		b <= e && a <= e && e <= d && e <= c ||
		a <= e && d <= e && e <= b && e <= c {
		fmt.Println("Среднее-", e)
	}
	if b == e && c == e && e == d && e == a ||
		d == e && c == e && e == b && e == a ||
		a == e && c == e && e == b && e == d ||
		b == e && d == e && e == c && e == a ||
		b == e && a == e && e == d && e == c ||
		a == e && d == e && e == b && e == c {

	}

}
