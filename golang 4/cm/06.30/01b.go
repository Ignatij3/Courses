package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func processing() {
	sumchan := make(chan int64)
	go task1(sumchan)
	maxchan := make(chan int64)
	go task2(maxchan)
	fmt.Println("Max: ", <-maxchan, "Sum: ", <-sumchan)
}

func task1(out chan <- int64)  {
	var sum int64
	data:= make(chan int64)
	go server(data)
	for x:= range data {
		sum += x
	}
	out <- sum
}

func task2(out chan <- int64)  {
	var max int64
	data:= make(chan int64)
	go server(data)
	for x:= range data {
		max = x		// first element
		break
	}
	for x:= range data {
		if x > max {
			max = x
		}
	}
	out <- max
}

func server(c chan <- int64)  {
	var	x int64
	fin, _ := os.Open("data.in")
	in := bufio.NewReader(fin)
	for {
		_, err:= fmt.Fscanf(in, "%d\n", &x)
		if err != nil  { break }
		c <- x
	}
	close(c)
	fin.Close()
}

func main() {
	t0 := time.Now()
	processing()
	t1 := time.Now()
	fmt.Printf("The call took %v to run.\n", t1.Sub(t0))
}
