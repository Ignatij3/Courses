package main

import (  
    "fmt"
    "time"
)

func nums(s string) {  
    for i, _ := range []rune(s) {
        time.Sleep(120 * time.Millisecond)
        fmt.Print(i, " ")
    }
}
func chars(s string) {  
    for _, c := range []rune(s) {
        time.Sleep(200 * time.Millisecond)
        fmt.Printf("%c ", c) //Буква с печатается буквенным образом
    }
}
func main() {  
    go nums("Hello_Go") //Печатает номера рун
    go chars("Hello_Go") //Печатает буквы
    time.Sleep(4000 * time.Millisecond)
    fmt.Printf("\nmain terminated\n")
}

/*
0 H 1 2 e 3 l 4 5 l 6 7 o _ G o 
main terminated
* _- -- _- -- _- -- _
*/
