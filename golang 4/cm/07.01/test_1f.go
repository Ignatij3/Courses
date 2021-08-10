package main

import (
	"fmt"
)

func write(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("send", i, "to channel")
	}
	close(ch)
}

func main() {
	ch := make(chan int, 2)
	go write(ch)
	for {
		a, more:= <- ch
		if !more {break}
		fmt.Println(a, "received")
	}	
}
/* OUTPUT
send 0 to channel
send 1 to channel
send 2 to channel
0 received
1 received
2 received
3 received
send 3 to channel
send 4 to channel
4 received
*/
