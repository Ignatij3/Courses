package main
import ("fmt"
			  )
func main() {
	var y int
	fmt.Println("Високосный ли этот год?")
	fmt.Scan(&y)
	if y%4==0 && y%400==0 {
		fmt.Println("Этот год високосный")
	} else {
			fmt.Println("Этот год не високосный")
		}
}
