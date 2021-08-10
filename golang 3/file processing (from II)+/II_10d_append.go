package main

import (  
    "fmt"
    "os"
)

func main() {  
    stat, err:= os.Stat("test.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(stat.Size())
    
    f, err := os.OpenFile("test.1", os.O_APPEND, 0666)
    if err != nil {
        fmt.Println(err)
        return
    }
	
	b:= []byte(" 012345")
/*	эквивалентно  
	b:= []byte {32, 48, 49, 50, 51, 52}
*/
	n, err := f.Write(b)
	fmt.Println(n)
    if err != nil {
        fmt.Println(err)
        return
    }

    stat, err = os.Stat("test.1")
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(stat.Size())

    err = f.Close()
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("file appended successfully")
}
