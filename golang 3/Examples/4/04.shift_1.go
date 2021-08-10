package main

import  "fmt"	

func swap(x []int, start1 int, start2 int, length int)  {
	i1, i2:= start1, start2
	for i:= 0; i < length; i++  {
		x[i1], x[i2] = x[i2], x[i1]
		i1++
		i2++
	}
}		

func shift(x []int, left0 int, right0 int, rightK int)  {
//	x - изменяемый массив
//	left0 - начало левого фрагмента
//	right0 - начало правого фрагмента
//	rightK - конец правого фрагмента
	for  {
		lengthL := right0 - left0  // длина левой части
		lengthR := rightK - right0 + 1  // длина правой части
		if lengthL < lengthR  {  // левая часть короче правой
			swap(x, left0, rightK - lengthL + 1, lengthL)
			rightK = rightK - lengthL
		}  else 	
		if lengthL > lengthR  {  // левая часть длиннее правой
			swap(x, left0, right0, lengthR)
			left0 = left0 + lengthR
		}  else  {
		// длины частей равны		
			swap(x, left0, right0, lengthL)
			break
		}	
	}	
}	

func main()  {
	var  (
		n, k int
	)
		
	fmt.Print("Введите длину массива: ")
	fmt.Scanln(&n)
	fmt.Print("Введите величину сдвига: ")
	fmt.Scanln(&k)
	
	x:= make([]int, n, n) 
 	for i:= 0; i < n; i++  {
		x[i] = i+1
	}
	fmt.Println(x)
	shift(x, 0,  k, n-1) 	
	fmt.Println(x)
}	
