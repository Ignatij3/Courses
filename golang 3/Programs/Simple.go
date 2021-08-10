package main

import (
	"fmt"
)

func main() {
	var (
		num int
		n int
		sum int
	)
	fmt.Print("Введите кол-во чисел:")
	fmt.Scan(&num)
	for i := 0; i < num; i++ {
		fmt.Print("Введите число:")
		fmt.Scan(&n)
		sum += n
	}
	fmt.Println("Вот их сумма:", sum)
}