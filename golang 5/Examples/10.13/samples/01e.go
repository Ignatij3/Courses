package main

import  "fmt"

// Добавляем в слайс-параметр данные группой, размер
// которой требует перераспределения памяти
func f1(c []int)  {
    cc:= make([]int, cap(c) - len(c) + 1)
    c = append(c, cc...)
    fmt.Println(c, len(c), cap(c))
}	 		 

func f2(pc *[]int)  {
    cc:= make([]int, cap(*pc) - len(*pc) + 1)
    *pc = append(*pc, cc...)
    fmt.Println(*pc, len(*pc), cap(*pc))
}	 

func f3(c []int) []int {
    cc:= make([]int, cap(c) - len(c) + 1)
    c = append(c, cc...)
    fmt.Println(c, len(c), cap(c))
    return c
}	 

func main() {  
    a:= [...]int {0,1,2,3,4,5,6,7,8,9}
    fmt.Println(a)                     // [0 1 2 3 4 5 6 7 8 9]
	
    arr:= a
    b:= arr[2:7]
    fmt.Println(b, len(b), cap(b))     // [2 3 4 5 6] 5 8
    f1(b)                              // [2 3 4 5 6 0 0 0 0] 9 16
    fmt.Println(b, len(b), cap(b))     // [2 3 4 5 6] 5 8
    fmt.Println(arr)                   // [0 1 2 3 4 5 6 7 8 9]
                                          
    arr = a                               
    b = arr[2:7]                          
    fmt.Println(b, len(b), cap(b))     // [2 3 4 5 6] 5 8
    f2(&b)                             // [2 3 4 5 6 0 0 0 0] 9 16
    fmt.Println(b, len(b), cap(b))     // [2 3 4 5 6 0 0 0 0] 9 16
    fmt.Println(arr)                   // [0 1 2 3 4 5 6 7 8 9]
                                          
    arr = a                               
    b = arr[2:7]                          
    fmt.Println(b, len(b), cap(b))     // [2 3 4 5 6] 5 8
    b = f3(b)                          // [2 3 4 5 6 0 0 0 0] 9 16
    fmt.Println(b, len(b), cap(b))     // [2 3 4 5 6 0 0 0 0] 9 16
    fmt.Println(arr)                   // [0 1 2 3 4 5 6 7 8 9]
}
