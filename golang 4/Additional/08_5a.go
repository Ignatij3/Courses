package main

import "fmt"

func main() {
	fmt.Println("Программа заполняет слайс числами Фибоначчи,")
	fmt.Println("не превосходящими заданного числа  N.")
	fmt.Print("Введите N: ")
	var n int64
	fmt.Scan(&n)

	fibo := []int64{1,1}
	prev, current := fibo[0], fibo[1]
	var next int64
	
	for  {
		next = prev + current
		if next > n  { break }
		fibo = append(fibo, next)
		prev, current = current, next
	}
	fmt.Println(fibo)
}
