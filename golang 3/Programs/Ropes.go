package main

import  (
	"fmt"
)

func Counting(Rcut []int, cut int) int {
	var (
		cutTotal, temp, ch int
		check bool
	)
	check = true
	for i := 0; check == true; {
		if ch == 0 {temp = Rcut[i]; ch = 1}
		temp -= cut
		if temp >= 1 {cutTotal++} else {i++; temp = Rcut[i]}
		if i == len(Rcut) - 1 {check = false}
	}
	return cutTotal
}

func main()  {
	var (
		ropeCut, app int
		cont bool = true
		continues string
		ropes []int
	)
	for i := 0; cont == true; i++ {
		fmt.Print("Введите размер верёвки (для прекращения введите \"0\"): ")
		fmt.Scan(&app)
		ropes = append(ropes, app)
		if ropes[i] == 0 {cont = false}
	}
	for {
		fmt.Print("Введите размер отрезков (что бы получить 11 равных частей): ")
		fmt.Scan(&ropeCut)
		for ropeCut <= 0 {
			fmt.Println("Данные введены неправильно, попробуйте ещё раз: ")
			fmt.Scan(&ropeCut)
		}
		amount := Counting(ropes, ropeCut)
		fmt.Println("Получилось", amount)
		if amount < 11 {
			fmt.Println("Число слишком большое!")
		}
		if amount > 11 {
			fmt.Println("Число слишком маленькое!")
		}
		if amount == 11 {
			fmt.Println("Поздравляем, число найдено!")
			break
		}
		fmt.Print("Желаете продолжить? (y/n)")
		fmt.Scan(&continues)
		for continues != "y" && continues != "n" {
			fmt.Print("Ошибка, введите ещё раз (y/n)")
			fmt.Scan(&continues)
		}
		if continues == "n" {break}
	}
}
