package main
import (
	"fmt"
)

func toBinary(array []int) []string {
	var (
		num, lenn int
		nArray []string
	)
	lenn = len(array)
	for j := 0; j < lenn; j++ {
		nArray = append(nArray, "")
	}
	for i := 0; i < len(array); {
		num = array[i]
		for ; num > 0; {
			if num % 2 == 1 {
				nArray[i] += "1" 	
			} else {
				nArray[i] += "0"
			}
			num = num / 2
		}
		i++
	}
	return nArray
}

func main() {
	var (
		num int
		array []int
		cont bool = true
		quest string
	)
	for cont {
		for num >= 0 {
			fmt.Print("Введите число для перевода (\"0\" - что бы закончить):")
			fmt.Scan(&num)
			for num < 0 {
				fmt.Println("Данные введены неправильно, попробуйте ещё раз: ")
				fmt.Scan(&num)
			}
			if num == 0 {break}
			array = append(array, num)
		}
		nArray := toBinary(array)
		for i := 0; i < len(nArray); i++ {
			fmt.Printf("Число %-9v переведено в %s\n", array[i], nArray[i] )
		}
		fmt.Print("Продолжить? (y/n):")
		fmt.Scan(&quest)
		for quest != "y" && quest != "n" {
			fmt.Println("Данные введены неправильно, попробуйте ещё раз: ")
			fmt.Scan(&quest)
		}
		if quest == "n" {cont = false}
	}
	
}
