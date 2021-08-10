package main

import (
	"math/rand"
	"time"
	"fmt"
)

func MakeSlice() []int {
	rand.Seed(time.Now().UnixNano())
	a := make([]int, 0, 0)
	n := 0
	for k := 0; k < 1000; k++ {
		n = rand.Intn(501)
		switch rand.Intn(2) {
			case 0:
				a = append(a, n)
			case 1:
				a = append(a, -n)
		}
	}
	return a
}

func QSort(a []int) []int {
	if len(a) <= 1 {return a}
	pivot, num := len(a) / 2, 0 //Пивот выбирается посередине слайса, num нам понадобится попозже
	for n, _ := range a {
		if n < pivot && a[n] >= a[pivot] { //Если число (которое мы проверяем) находится слева и оно больше, заходим сюда
			if pivot - 1 == n { //Если пивот слева от числа, просто меняю местами
				a[pivot], a[n] = a[n], a[pivot]
				pivot-- //Меняю указатель на пивот на текущую позицию
			} else {
				if pivot < len(a) - 1 {num = pivot + 1} else {num = len(a) - 1} //Это для того, что-бы num не уходил за рамки массива
				for num < len(a) - 1 && a[num] >= a[pivot] {num++} //Прохожусь по массиву в поиске чисел меньше пивота, начиная с пивота + 1, или конца
				
				if num == len(a) - 1 && a[num] >= a[pivot] { //Если не нашёл, захожу сюда
					if pivot > 0 {num = pivot - 1} else {num = 0}
					for num > 0 && a[num] >= a[pivot] {num--} //Прохожусь по массиву в поиске чисел меньше пивота, но уже в другую сторону
					
					if num == 0 && a[num] >= a[pivot] { //Если таковых не оказалось, захожу сюда (получается число больше всех)
						for num < len(a) - 1 && a[num] == a[pivot] {num++} //Иду с правого конца налево, если самых больших чисел несколько
						a[num], a[pivot] = a[pivot], a[num] //Заменяю пивот на число в конце и число на пивот
						pivot = num //Меняю указатель на пивот на текущую позицию
					} else { //Если таки нашлись
						if num < n { //Если найденное число меньше того, которое мы проверяли, то просто меняю местами пивот и проверяемое число
							a[pivot], a[n] = a[n], a[pivot]
							pivot = n
						} else { //В ином случае, перемещаю числа так, что бы шли по возрастанию
							a[num], a[pivot], a[n] = a[pivot], a[n], a[pivot]
							pivot = num
						}
					}
				} else {
					a[num], a[n] = a[n], a[num] //Если нашлось число меньше пивота (46 строчка), меняю местами с проверяемым
				}
			}
		} else if n >= pivot && a[n] < a[pivot] { //Тут всё аналогично случаю, когда число находится слева и оно больше, только наоборот
			if pivot + 1 == n {
				a[pivot], a[n] = a[n], a[pivot]
				pivot++
			} else {
				if pivot > 0 {num = pivot - 1} else {num = 0}
				for num > 0 && a[num] < a[pivot] {num--}
				
				if num == 0 && a[num] < a[pivot] {
					if pivot < len(a) - 1 {num = pivot + 1} else {num = len(a) - 1}
					for num < len(a) - 1 && a[num] < a[pivot] {num++}
					
					if num == len(a) - 1 && a[num] < a[pivot] {
						for num > 0 && a[num] == a[pivot] {num--}
						a[num], a[pivot] = a[pivot], a[num]
						pivot = num
					} else {
						if num > n {
							a[pivot], a[n] = a[n], a[pivot]
							pivot = n
						} else {
							a[num], a[pivot], a[n] = a[pivot], a[n], a[pivot]
							pivot = num
						}
					}
				} else {
					a[num], a[n] = a[n], a[num]
				}
			}
		}
	}
	
	QSort(a[pivot+1:])
	QSort(a[:pivot])
	return a
}

func main() {
	a := MakeSlice()
	fmt.Print("Slice before:\n")
	fmt.Println(a, "\n")
	a = QSort(a)
	fmt.Print("Slice after:\n")
	fmt.Println(a, "\n")
}
