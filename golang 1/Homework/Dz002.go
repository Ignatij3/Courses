package main
import ("fmt"
			  )
func main() {
var w1, h1, l1, w2, h2, l2 float64
fmt.Println("Хотите узнать влезет ли ваш крипич в дырку? Вперёд!")
fmt.Println("Данные о кирпиче:")
fmt.Println("Ширина")
fmt.Scan(&w1)
fmt.Println("Высота")
fmt.Scan(&h1)
fmt.Println("Длина")
fmt.Scan(&l1)
fmt.Println("Данные о дырке:")
fmt.Println("Ширина")
fmt.Scan(&w2)
fmt.Println("Высота")
fmt.Scan(&h2)
fmt.Println("Длина")
fmt.Scan(&l2)
if w1>w2 {
	fmt.Println("Не влезет")
	} else {
		fmt.Println("Влезет")
		}
if h1>h2 {
	fmt.Println("Не влезет")
	} else {
		fmt.Println("Влезет")
		}




}
/*
 *  w1 - ширина (кирпича)
 *  h1 - высота 
 *  l1 - длина
 *  w2 - ширина (дырки)
 *  h2 - высота
 *  l2 - длина
 */	 
