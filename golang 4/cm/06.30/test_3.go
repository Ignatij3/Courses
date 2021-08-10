package main

import (  
    "fmt"
    "time"
)

func server1(ch chan []int) {  
	list:= make([]int, 1)
	list[0] = 1111
	for i:= 1; i<=10; i++  {
		list = append(list, i)
		time.Sleep(80 * time.Millisecond)
	}
	ch <- list	
}

func server2(ch chan []int) {  
	list:= make([]int, 1)
	list[0] = 2222
	for i:= 1; i<=10; i++  {
		list = append(list, i)
		time.Sleep(50 * time.Millisecond)
	}
	ch <- list	
}

func main() {  
    output1 := make(chan []int)
    output2 := make(chan []int)
    go server1(output1)
    go server2(output2)
    select {
    case s1 := <-output1:
        fmt.Println(s1)
    case s2 := <-output2:
        fmt.Println(s2)
    }
}

/* Output:
[2222 1 2 3 4 5 6 7 8 9 10]
*/

