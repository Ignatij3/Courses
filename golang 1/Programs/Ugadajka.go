package main
import "fmt"
func main() {
var a, b, c, d, n float64
	fmt.Print("Введите кол-во первых чисел Фиббоначи, для сложения:")
	fmt.Scan(&n)
	a, b = 0, 1
	d = 1
	for i :=  2 ; i < int(n); i++  {
		c = a+b
		a, b = b, c
		d+=c
	}
		fmt.Println(d)
}
