package main

import (
	"os"
)

func Writer(id string, quit chan int) {
	text, _ := os.Create(id + "infText.txt")
	defer text.Close()
	for {text.WriteString("NEWLINE\n")}
	quit <- 1
}

func main() {
	var chanSlice []chan int
	
	for i := 500; i < 503; i++ {
		chanSlice = append(chanSlice, make(chan int, 1))
	}
	
	for id, ch := range chanSlice {
			go Writer(string(id), ch)
	}
	for _, ch2 := range chanSlice {<-ch2}
}
