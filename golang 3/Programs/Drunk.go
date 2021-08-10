package main

import  (
	"fmt"
	"math/rand"
)


func Walking(width, man, random, sober, sim int) (string[], int[]) {
	var (
		i, temp, temp2, cI, check int
		side []string
		iter []int
	)
	man = width / 2
	for j := 0; j < sim; j++ {
		for i = 0; i >= 0; i++ {
			temp = 1 + rand.Intn(100)
			if random - temp >= 0 && check != 0 {man--} else {man++}
			iter[j] = i
			if man == 0 {side[j] = "левой"; break}
			if man == width {side[j] = "правой"; break}
			temp2 = 1 + rand.Intn(100)
			if sober - temp2 >= 0 {check = 0; cI = i}
			if cI + 2 == i {check = 1}
		}
		i = 0
	}
	return side, iter
}

func main() {
	var (
		man, width, random, sober, simulations int
		/*side []string
		steps []int*/
	)
	fmt.Print("Введите кол-во симуляций: ")
	fmt.Scan(&simulations)
	for simulations <= 0 {
		fmt.Println("Данные введены неправильно, попробуйте ещё раз: ")
		fmt.Scan(&simulations)
	}
	fmt.Print("Введите ширину дороги: ")
	fmt.Scan(&width)
	for width <= 0 {
		fmt.Println("Данные введены неправильно, попробуйте ещё раз: ")
		fmt.Scan(&width)
	}
	fmt.Print("Введите % отклонения в левую сторону: ")
	fmt.Scan(&random)
	for random < 0 || random > 100 {
		fmt.Println("Данные введены неправильно, попробуйте ещё раз (0 - 100): ")
		fmt.Scan(&random)
	}
	fmt.Print("Введите % трезвения: ")
	fmt.Scan(&sober)
	for sober < 0 || sober > 100 {
		fmt.Println("Данные введены неправильно, попробуйте ещё раз (0 - 100): ")
		fmt.Scan(&sober)
	}
	side, steps := Walking(width, man, random, sober, simulations)
	for i := 0; steps[i] != 0; i++ {
		fmt.Println("Во время", i, "-й", "симуляции, пьяница свалился с пирса с", side[i], "стороны, через", steps[i], "шагов")
	}
}
