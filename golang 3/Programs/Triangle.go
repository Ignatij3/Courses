package main

import  (
	"fmt"
)

func Modeling(a, b, c float64) string {
	var cases string
	if (a == b  && a + b > c) || (a == c && a + c > b) || (b == c && b + c > a) {cases = "равнобедренный"}
	if (a != b && b != c && c != a) && (a + b > c || a + c > b || b + c > a) {cases = "неравносторонний"}
	if a + b <= c || a + c <= b || b + c <= a {cases = "несуществует"}
	if a == b && b == c && c == a {cases = "равносторонний"}
	return cases
}

func main()  {
	var (
		a, b, c float64
	)
	fmt.Print("Введите первую сторону треугольника: ")
	fmt.Scan(&a)
	for a <= 0 {
		fmt.Print("Данные введены неправильно, попробуйте ещё раз: ")
		fmt.Scan(&a)
	}
	fmt.Print("Введите вторую сторону треугольника: ")
	fmt.Scan(&b)
	for b <= 0 {
		fmt.Print("Данные введены неправильно, попробуйте ещё раз: ")
		fmt.Scan(&b)
	}
	fmt.Print("Введите третью сторону треугольника: ")
	fmt.Scan(&c)
	for c <= 0 {
		fmt.Print("Данные введены неправильно, попробуйте ещё раз: ")
		fmt.Scan(&c)
	}
	result := Modeling(a, b, c)
	if result == "несуществует" {
		fmt.Println("Данный треугольник невозможно составить")
	} else {
		fmt.Println("У вас получился", result, "треугольник")
	}
}
