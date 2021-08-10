package main

type plottedFunction func(float64) float64//Функциональный тип

func Plot(left, right float64, F plottedFunction) {//Графопостроимая функция
// . . .
}

func cube(x float64) float64 {
	return x * x * x
}

func main() {
	Plot(-1.5, 2.0, cube)
}
