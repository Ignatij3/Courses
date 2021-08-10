package main

import (  
    "fmt"
)

func minus(x *int, quit chan bool) {  
	for  *x < 20 {
		(*x)++
		fmt.Print(-(*x), " ")
    }	
    quit <- true
}

func plus(x *int, quit chan bool) {  
	for  *x < 15 {
		(*x)++
		fmt.Print(*x, " ")
	}	
    quit <- true
}

func main() {  
	var c int
	ch1, ch2 := make(chan bool), make (chan bool) 
    go minus(&c, ch1)
    go plus(&c, ch2)
    <-ch1 //Ждём, пока эти 2 канала станут true
    <-ch2
    fmt.Println("main terminated at value", c)
}
/*
Четыре варианта вывода:
1 3 4 5 6 7 8 9 10 11 12 13 14 15 -2 -16 -17 -18 -19 -20 main terminated at value 20 --- минус остановился на 2-е, но не отпечатал (не успел), далее плюс продолжил
действовать до 15-и, -2 отпечаталась и продолжила с -16-и.
1 3 -2 -5 -6 -7 -8 -9 -10 -11 -12 -13 -14 -15 -16 -17 -18 -19 -20 4 main terminated at value 20
-2 -3 -4 1 6 7 8 9 -5 -11 -12 -13 -14 -15 -16 -17 -18 -19 -20 10 main terminated at value 20
1 3 4 5 6 7 8 9 10 11 12 13 -2 -15 -16 -17 -18 -19 -20 14 main terminated at value 20
*/
