package main

import  (
	"fmt"
	"math/rand"
)

func Modeling(people, sim int)  {
	var (
		day int
		result [366]int
	)
	wasBornInOneYear := 0.0
	for i := 0; i < sim; i++ {
		for i := 0; i < 366; i++ {
			result[i] = 0
		}
		for j := 0; j < people; j++ {
			day = rand.Intn(366)
		    result[day]++
		    if (result[day]) > 1 {
				wasBornInOneYear++
				break
			}
		}
	}
	fmt.Println("Кол-во совпадений дней рожденья: ", wasBornInOneYear)
	fmt.Println("Процент совпадений:", 100 * (wasBornInOneYear / float64(sim)), "%")
}

func main() {
	var (
		people, simulations int
	)
	fmt.Print("Введите кол-во людей: ")
	fmt.Scan(&people)
	fmt.Print("Введите кол-во симуляций: ")
	fmt.Scan(&simulations)
	Modeling(people, simulations)
	
}
