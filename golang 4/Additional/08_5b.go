package main

import "fmt"

func main() {
	fmt.Println("Программа вычисляет такую последовательность:")
	fmt.Println("ее K-й член - это число, образованное двумя")
	fmt.Println("последними цифрами числа N^K.")
	fmt.Println()

	var n int
	for n <= 0 || n >= 100 {
		fmt.Print("Введите N (от 1 до 99): ")
		fmt.Scan(&n)
	}

	x := []int{1}
	var k, i int
	for k == i {
		x = append(x, (x[k]*n)%100)
		k++
		i = 0
		for x[i] != x[k] {
			i++
		}
	}
	if i > 0 {
		fmt.Printf("\nПолучили последовательность:\n")
		fmt.Println(x[:i])
		fmt.Printf("\nДалее последовательность периодически повторяется:\n")
	} else {
		fmt.Println("\nПоследовательность периодически повторяется:\n")
	}
	fmt.Println(x[i:k])
	fmt.Printf("\nДлина периода равна %d\n", k-i)
}
