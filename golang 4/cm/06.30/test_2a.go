package main

func getData(send_channel chan <- int) {  
    send_channel <- 10
}

func main() {  
    ch := make(chan <- int)
    go getData(ch)
    x := <-ch
    // . . . 
}
