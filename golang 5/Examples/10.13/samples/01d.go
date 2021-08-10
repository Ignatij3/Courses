package main

import 	"fmt"	

// Добавляем в слайс-параметр данные по одному, 
// пока не произойдёт перераспределения памяти
func f1(c []int)  {
    ccap:= cap(c)
    for cap(c) == ccap  {
        c = append(c, len(c)*10)
    }			
    fmt.Println(c, len(c), cap(c))
}		 

func f2(pc *[]int)  {
    ccap:= cap(*pc)
    for cap(*pc) == ccap  {
        *pc = append(*pc, len(*pc)*10)
    }			
    fmt.Println(*pc, len(*pc), cap(*pc))
}	 

func f3(c []int) []int {
    ccap:= cap(c)
    for cap(c) == ccap  {
        c = append(c, len(c)*10)
    }			
    fmt.Println(c, len(c), cap(c))
    return c
}	 

func main() {  
    a:= [...]int {0,1,2,3,4,5,6,7,8,9}
    fmt.Println(a)                     // [0 1 2 3 4 5 6 7 8 9]
	
    arr:= a
    b:= arr[2:7]
    fmt.Println(b, len(b), cap(b))     // [2 3 4 5 6] 5 8
    f1(b)                              // [2 3 4 5 6 50 60 70 80] 9 16
    fmt.Println(b, len(b), cap(b))     // [2 3 4 5 6] 5 8
    fmt.Println(arr)                   // [0 1 2 3 4 5 6 50 60 70]
                                          
    arr = a                               
    b = arr[2:7]                          
    fmt.Println(b, len(b), cap(b))     // [2 3 4 5 6] 5 8
    f2(&b)                             // [2 3 4 5 6 50 60 70 80] 9 16
    fmt.Println(b, len(b), cap(b))     // [2 3 4 5 6 50 60 70 80] 9 16
    fmt.Println(arr)                   // [0 1 2 3 4 5 6 50 60 70]
                                          
    arr = a                               
    b = arr[2:7]                          
    fmt.Println(b, len(b), cap(b))     // [2 3 4 5 6] 5 8
    b = f3(b)                          // [2 3 4 5 6 50 60 70 80] 9 16
    fmt.Println(b, len(b), cap(b))     // [2 3 4 5 6 50 60 70 80] 9 16
    fmt.Println(arr)                   // [0 1 2 3 4 5 6 50 60 70]
}
