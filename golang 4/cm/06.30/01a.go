package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func processing() {
	ch:= make(chan int64)
	go server(ch)
	var  (
		x int64
		ok bool
	)	
	max := <- ch
	sum := max
	for  {
		x, ok = <- ch
		if ok  {	
			sum += x
			if x > max  { max = x }
		}  else {
			break
		}		
	}	
	fmt.Println(max, sum)
}	

func server(c chan <- int64)  {//Работает только на запись
	var x int64
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
	fmt.Printf("The call took %v to run.\n", t1.Sub(t0)) //Difference
}
