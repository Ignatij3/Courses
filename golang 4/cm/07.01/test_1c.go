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
	close(ch)
}
func main() {
	ch := make(chan int, 2)
	go write(ch)
	time.Sleep(500 * time.Millisecond)
	for v := range ch {
		fmt.Println(v, "received")
		time.Sleep(500 * time.Millisecond)
	}
}
/*OUTPUT

send 0 to channel
send 1 to channel
0 received
send 2 to channel
1 received
send 3 to channel
2 received
send 4 to channel
3 received
4 received
*/
