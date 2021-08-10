package main
import "fmt"
func main() {
	const V = 7 //*Максимально возможное число "140.737.488.355.327"*//
	var c [V] int
	var i, q, t, s int
	for i = 0; i < V; i++ {
		fmt.Print("Введите ", i, "-е число: ")
		fmt.Scan(&c[i])
	}
		for  i = 0; i < V; i++ {
			fmt.Print("Я смотрю ", i)
			fmt.Println("-е по счёту число. Оно равно", c[i])
		
		if c[i] > 0 {
			fmt.Println("           ...и оно больше нуля.")
			q++
		}
	}
			for s = 0; s < V; s++ {
			fmt.Print("Я смотрю ", s)
			fmt.Println("-е по счёту число. Оно равно", c[s])
				if c[s] <= 0 {
					fmt.Println("           ...и оно меньше, или равно нулю.")
					t++
				}
			}
	fmt.Println("Здесь", q, "положительных числа, и", t, "отрицательных числа")
}
