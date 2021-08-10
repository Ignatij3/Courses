package main

import 	"fmt"

// Изменяем только данные слайса-параметра
func f1(c []int)  {
    for i, _ := range c  {
         c[i] +=100
    }			
    fmt.Println(c, len(c), cap(c))
}	 

func f2(pc *[]int)  {
    for i, _ := range *pc  {
        (*pc)[i] +=100
    }			
    fmt.Println(*pc, len(*pc), cap(*pc))
}	 

func f3(c []int) []int {
    for i, _ := range c  {
        c[i] +=100
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
    f1(b)                              // [102 103 104 105 106] 5 8
    fmt.Println(b, len(b), cap(b))     // [102 103 104 105 106] 5 8
    fmt.Println(arr)                   // [0 1 102 103 104 105 106 7 8 9]
  
    arr = a
    b = arr[2:7]
    fmt.Println(b, len(b), cap(b))     // [2 3 4 5 6] 5 8
    f2(&b)                             // [102 103 104 105 106] 5 8
    fmt.Println(b, len(b), cap(b))     // [102 103 104 105 106] 5 8
    fmt.Println(arr)                   // [0 1 102 103 104 105 106 7 8 9]
  
    arr = a
    b = arr[2:7]
    fmt.Println(b, len(b), cap(b))     // [2 3 4 5 6] 5 8
    b = f3(b)                          // [102 103 104 105 106] 5 8
    fmt.Println(b, len(b), cap(b))     // [102 103 104 105 106] 5 8
    fmt.Println(arr)                   // [0 1 102 103 104 105 106 7 8 9]
}
