package main

import  (
	"fmt"
	"math"
)	

func main()  {
	var (
		a, b, c float64
		D float64
	)
		
	fmt.Println("Решаем уравнение  ax^2+bx+c=0\n\n")
	fmt.Print("Введите a: "); fmt.Scan(&a)
	fmt.Print("  Введите b: "); fmt.Scan(&b)
	fmt.Print("    Введите c: "); fmt.Scan(&c)
	if a != 0  {
		D = b * b - 4 * a * c
		if  D < 0  {
			fmt.Println("Нет решений")
		}  else  {
			if  D > 0  {
				D := math.Sqrt(D)
				fmt.Printf("Два решения: %7.6f \n", (-b - D)/(2 * a))
				fmt.Printf("             %7.6f \n", (-b + D)/(2 * a))
			}  else  {
				fmt.Printf("Единственное решение:  %7.6f \n",(-0.5 * b / a))
			}
		}
	}  else  {
		//  a == 0
		if b != 0  {
			fmt.Printf("Единственное решение:  %7.6f \n", -c / b)
		}  else  {
		//  a == b == 0
			if c != 0  {
				fmt.Println("Нет решений")
			}  else  {
				fmt.Println("x - любое число")
			}
		}
	}
}
