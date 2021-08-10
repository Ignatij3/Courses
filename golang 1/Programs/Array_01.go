package main
import "fmt"
func main() {
	var k[14090] int
	var c, x, cx int
	
	fmt.Println("Введите числа, что-бы закончить и получить их количество и сумму напишите 0")
	fmt.Scan(&x)
	for c = 0; x != 0; c++ {
		k[c] = x
		cx = cx + x
		fmt.Scan(&x)
	}
	fmt.Println("Всего было:"," ", c," ", "чисел.")
	fmt.Println("Их сумма:"," ",cx)
	
	
}
