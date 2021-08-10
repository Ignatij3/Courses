package main

import  (
	"fmt"
	"os"
)	

func main()  {
	var  (
		a []int
		x int
		left, center, right int
	)
		
	fin, _ := os.Open("inc.dat")
	defer fin.Close()

 	for {
		if _, err := fmt.Fscanln(fin, &x); err == nil {
			a = append(a, x)
		}  else  {
			break
		}	
	}
	fmt.Println(a)
	
	fmt.Print("Что ищем? ")
	fmt.Scanln(&x)
	
	left, right = 0, len(a)-1
	// Проверяем выполнение условий продолжения поиска
	if a[left] > x  ||  a[right] < x  {
		fmt.Println("Нет такого числа")
		return
	}	 
	if a[left] == x  {
		fmt.Println("Число", x, "находится на", left, "-м месте.")
		return
	}	 
	if a[right] == x  {
		fmt.Println("Число", x, "находится на", right, "-м месте.")
		return
	}	
	// Входим в цикл поиска.
	// Выполняется инвариант: a[left] < x < a[right]
	for  left+1 < right  {
		center = (left + right) / 2
		if a[center] < x  {
			left = center
		}  else  
		if a[center] > x  {
			right = center
		}  else {
		// a[center] == x
			fmt.Println("Число", x, "находится на", center, "-м месте.")	
			return
		}
	}		
	fmt.Println("Нет такого числа")
}	
