package main

import (
	"fmt"
	"time"
)

func write(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("send", i, "to channel")
		time.Sleep(700 * time.Millisecond)
	}
}

func main() {
	ch := make(chan int, 2)
	go write(ch)
	time.Sleep(500 * time.Millisecond)
	for  {
		fmt.Println(<-ch, "received")
		time.Sleep(500 * time.Millisecond)
		if len(ch)==0 {break}
	}	
}
/* OUTPUT
send 0 to channel
0 received
send 1 to channel
1 received
send 2 to channel
2 received
*/