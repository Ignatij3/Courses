package main

import (
	"fmt"
)

func write(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("send", i, "to channel")
	}
}

func main() {
	ch := make(chan int, 2)
	go write(ch)
	for  {
		fmt.Println(<-ch, "received")
	}	
}

/*  OUTPUT
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
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan receive]:
main.main()
	G:/03/test_samples/test_1a.go:18 +0x84
*/
