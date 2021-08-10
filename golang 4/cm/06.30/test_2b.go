package main

import "fmt"

func getData(send_channel chan <- int) {  
    send_channel <- 2020
}

func main() {  
    ch := make(chan int)
    go getData(ch)
    fmt.Println(<-ch)
}
