package main

import (
	"fmt"
	"time"
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
		time.Sleep(500 * time.Millisecond)
		fmt.Println(<-ch, "received")
	}
}

/* OUTPUT
send 0 to channel
send 1 to channel
0 received
send 2 to channel
send 3 to channel
1 received        
2 received
send 4 to channel
3 received
4 received
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan receive]:
main.main()
	G:/03/test_samples/test_1b.go:20 +0x91

*/ 
