package main

import  (
	"fmt"
	"math/rand"
)

func DiceRolling(money, bet, betNum int) int {
	var (
		check, result int
		continues string
	)
	if money != 0 {check = 1}
	for i := 0; money != 0 && check == 1; i++ {
		result = 1 + rand.Intn(6) + 1 + rand.Intn(6) + 1 + rand.Intn(6)
		if betNum != result {
			money -= bet
			fmt.Println("Ваша ставка не сыграла, с вашего счёта было списано", bet)
		}
		if betNum == result {
			money += bet
			fmt.Println("Ваша ставка сыграла, на ваш счёт было зачислено", bet)
		}
		fmt.Println("Теперь на вашем счёте", money)
		if money == 0 {break}
		fmt.Print("Желаете продолжить? (y/n)")
		fmt.Scan(&continues)
		for continues != "y" && continues != "n" {
			fmt.Print("Ошибка, введите ещё раз (y/n)")
			fmt.Scan(&continues)
		}
		if continues == "n" {break}
		if continues == "y" {
			result = 0
			fmt.Print("Введите вашу ставку: ")
			fmt.Scan(&bet)
			for bet > money {
				fmt.Print("Ошибка, введите ещё раз: ")
				fmt.Scan(&bet)
			}
			fmt.Print("Введите число, на которое ставите(3 - 18): ")
			fmt.Scan(&betNum)
			for  betNum < 3 || betNum > 18 {
				fmt.Print("Ошибка, введите ещё раз: ")
				fmt.Scan(&betNum)
			}
		}
	}
	for i := 0; money != 0 && check == 0; i++ {
		result = 1 + rand.Intn(6) + 1 + rand.Intn(6) + 1 + rand.Intn(6)
		if betNum != result {
			bet -= bet
			fmt.Println("Ваша ставка не сыграла, с вашего счёта было списано", bet)
		}
		if betNum == result {
			bet += bet
			fmt.Println("Ваша ставка сыграла, на ваш счёт было зачислено", bet)
		}
		money = bet
		fmt.Println("Теперь на вашем счёте", money)
		if money == 0 {break}
		fmt.Print("Желаете продолжить? (y/n)")
		fmt.Scan(&continues)
		for continues != "y" && continues != "n" {
			fmt.Print("Ошибка, введите ещё раз (y/n)")
			fmt.Scan(&continues)
		}
		if continues == "n" {break}
		if continues == "y" {
			result = 0
			fmt.Print("Введите вашу ставку: ")
			fmt.Scan(&bet)
			for bet < 0 {
				fmt.Print("Ошибка, введите ещё раз: ")
				fmt.Scan(&bet)
			}
			fmt.Print("Введите число, на которое ставите(3 - 18): ")
			fmt.Scan(&betNum)
			for  betNum < 3 || betNum > 18 {
				fmt.Print("Ошибка, введите ещё раз: ")
				fmt.Scan(&betNum)
			}
		}
	}
	return money
}

func main() {
	var (
		bet, betNum, money int = 0, 0, 0
		question string
	)
	fmt.Print("Вы хотите иметь начальный капитал? (y/n)")
	fmt.Scan(&question)
	for question != "y" && question != "n" {
		fmt.Print("Ошибка, введите ещё раз (y/n)")
		fmt.Scan(&question)
	}
	if question == "y"{
		fmt.Print("Введите начальный капитал: ")
		fmt.Scan(&money)
		for money == 0 {
			fmt.Print("Ошибка, введите ещё раз: ")
			fmt.Scan(&money)
		}
	}
	fmt.Print("Введите вашу ставку: ")
	fmt.Scan(&bet)
	for (question != "y" && bet > money) || bet < 0 {
		fmt.Print("Ошибка, введите ещё раз: ")
		fmt.Scan(&bet)
	}
	fmt.Print("Введите число, на которое ставите(3 - 18): ")
	fmt.Scan(&betNum)
	for betNum < 3 || betNum > 18 {
		fmt.Print("Ошибка, введите ещё раз: ")
		fmt.Scan(&bet)
	}
	fMoney := DiceRolling(money, bet, betNum)
	fmt.Println("У вас на счёте", fMoney, "евро")
}
