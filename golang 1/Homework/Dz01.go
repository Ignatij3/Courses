package main
import (
	"fmt"
	)
	func main() {
		var a, b, c, d, e, f, g float64
		fmt.Scan (&a)
		fmt.Scan (&b)
		fmt.Scan (&c)
		g=a+b
		f=c/g
		d=f*a
		e=f*b
		fmt.Println (a,"+",b,"=",g)
		fmt.Println (c,"/",g,"=",f)
		fmt.Println (f,"*",a,"=",d)
		fmt.Println (f,"*",b,"=",e)
		fmt.Print ("Первый поезд проедет",d, "km")
		fmt.Print ("Второй поезд проедет",e, "km")
	}
/*
 * a - Скорость №1
 * b - Скорость №2
 * c - Расстояние
 * d - Расстояние №1
 * e - Расстояние №2
 * f - Время встречи
 * g - Сумма скоростей
 * Я не знаю, как поставить пробел в конце, поэтому оставил так
 */
