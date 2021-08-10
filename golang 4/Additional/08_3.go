package main

import (
	"fmt"
	"math"
)

type TPolynom []float64

// P вычисляет значение многочлена в точке x "впрямую"
func P(a TPolynom, x float64) (y float64) {
	xi := 1.0
	for _, c := range a {
		y += c * xi
		xi *= x
	}
	return
}

// PHorner вычисляет значение многочлена в точке x по схеме Горнера
func PHorner(a TPolynom, x float64) (y float64) {
	for i := len(a) - 1; i >= 0; i-- {
		y = y*x + a[i]
	}
	return
}

func main() {
	// Ввод многочлена
	n := -1
	for n < 0 {
		fmt.Print("Введите степень многочлена: ")
		fmt.Scan(&n)
	}
	a := make(TPolynom, n+1)
	fmt.Print("Введите свободный член: ")
	fmt.Scan(&a[0])
	for i := 1; i <= n; i++ {
		fmt.Printf("Введите коэффициент при %d-й степени: ", i)
		fmt.Scan(&a[i])
	}

	// Ввод диапазона и шага изменения переменной
	var x0, xk float64
	fmt.Print("Введите диапазон значений x (через пробел): ")
	fmt.Scan(&x0, &xk)
	x0, xk = math.Min(x0, xk), math.Max(x0, xk)

	var h float64
	fmt.Print("Введите шаг приращения x: ")
	for h <= 0 {
		fmt.Scan(&h)
	}
	fmt.Println()

	// Вычисление многочлена двумя способами и печать результатов
	lines := 0
	for x := x0; x <= xk; x += h {
		fmt.Printf("%11.4f  %g  %g\n", x, P(a, x), PHorner(a, x))
		lines++
		if lines == 10 {
			fmt.Println("Для продолжения нажмите <ВВОД>")
			lines = 0
			fmt.Scanln()
		}
	}
}
