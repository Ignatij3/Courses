package main

import 	"fmt"	

// Изменяем только длину слайса-параметра
func f1(c []int)  {
    c = c[:cap(c)]
    fmt.Println(c, len(c), cap(c))
}	 

func f2(pc *[]int)  {
    *pc = (*pc)[:cap(*pc)]
    fmt.Println(*pc, len(*pc), cap(*pc))
}	 

func f3(c []int) []int {
    c = c[:cap(c)]
    fmt.Println(c, len(c), cap(c))
    return c
}	 

func main() {  
    a:= [...]int {0,1,2,3,4,5,6,7,8,9}
    fmt.Println(a)                     // [0 1 2 3 4 5 6 7 8 9]
	
    arr:= a
    b:= arr[2:7]
    fmt.Println(b, len(b), cap(b))     // [2 3 4 5 6] 5 8
    f1(b)                              // [2 3 4 5 6 7 8 9] 8 8
    fmt.Println(b, len(b), cap(b))     // [2 3 4 5 6] 5 8
    fmt.Println(arr)                   // [0 1 2 3 4 5 6 7 8 9]
                                          
    arr = a                               
    b = arr[2:7]                          
    fmt.Println(b, len(b), cap(b))     // [2 3 4 5 6] 5 8
    f2(&b)                             // [2 3 4 5 6 7 8 9] 8 8
    fmt.Println(b, len(b), cap(b))     // [2 3 4 5 6 7 8 9] 8 8
    fmt.Println(arr)                   // [0 1 2 3 4 5 6 7 8 9]
                                          
    arr = a                               
    b = arr[2:7]                          
    fmt.Println(b, len(b), cap(b))     // [2 3 4 5 6] 5 8
    b = f3(b)                          // [2 3 4 5 6 7 8 9] 8 8
    fmt.Println(b, len(b), cap(b))     // [2 3 4 5 6 7 8 9] 8 8
    fmt.Println(arr)                   // [0 1 2 3 4 5 6 7 8 9]
}
