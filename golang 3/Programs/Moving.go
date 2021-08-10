package main
import (
		"fmt"
		)

func Moving(array []int) {
	Temp := 0
	n := 0
	fmt.Println("Введите n:")
	fmt.Scan(&n)
	for i := 0; i < len(array); i++ {
		if i == len(array) - n {
			break
		} else {
			Temp = array[i + n]
			array[i + n] = array[i]
			array[i] = Temp
		}
	}
}

func main() {
	Array := []int {5, 12, 9, 2, 34, 22, 17, 8, 18, 48, 31, 20, 11, 19, 6, 51}
	Moving(Array)
	fmt.Println(Array)
}
