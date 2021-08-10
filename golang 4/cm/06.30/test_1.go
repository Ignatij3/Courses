package main

import (  
    "fmt"
    "math/rand"
    "time"
)

func datagen(c chan int) {  
    for i := 0; i < 6; i++ {
        c <- rand.Intn(6) + 1
    }
    close(c)
}

func main() {  
	rand.Seed(int64(time.Now().Nanosecond()))
    ch := make(chan int)
    go datagen(ch)
    i := 0
    for {
        v, prolong := <-ch //1-я - То что в канале, 2-я - Закрыт или открыт
        if !prolong { // prolong == false
            break
        }
        i++
        fmt.Printf("Dice #%d - %d point(s)\n", i, v)
    }
}
