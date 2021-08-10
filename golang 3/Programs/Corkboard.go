package main

import  (
	"fmt"
)

func Diploma(wh, ht, sew, seh int) int{
	var (
		count int
		backup int
		i int
	)
	backup = sew
	i = 1
	if sew < wh || seh < ht {
		count = -1
		i = 0
	}
	if sew == wh && seh == ht {
		count = 0
		i = 0
	}
	seh -= ht
	for i > 0 {
		//fmt.Println(wh, sew, ht, seh)
		sew -= wh
		count++
		if sew - wh < 0 {
			sew = backup
			if seh - ht < 0 {
				i = 0
			}
			seh -= ht
		}
	}
	return count
}

func main() {
	var (
		width int
		height int
		side int
	)
	fmt.Print("Введите ширину дипломов: ")
	fmt.Scan(&width)
	fmt.Print("Введите высоту дипломов: ")
	fmt.Scan(&height)
	fmt.Print("Введите сторону квадратной пробковой доски: ")
	fmt.Scan(&side)
	number := Diploma(width, height, side, side)
	if number < 0 {
		fmt.Println("На этой доске не поместится ни одного диплома!")
	} else if number == 0 {
		fmt.Println("На этой доске поместится ровно 1 диплом!")
	} else if number > 0 {
		fmt.Println("На этой доске поместится", number, "дипломов!")
	}
	
}
