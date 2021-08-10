package main

import (  
	"fmt"
	"math/rand"
	"time"
)

func toss(max int, amount int, res chan int) {  
	sum:= 0
	for i := 0; i < amount; i++ {
		 sum +=  rand.Intn(max) + 1
	}
	res <- sum 
}

func main() {  
	rand.Seed(time.Now().UnixNano())
	achan := make (chan int)
	bchan := make (chan int)
	go toss(2, 10000, achan)  // 10000 раз бросаем монету - 1 или 2
	go toss(6, 30000, bchan)  // 30000 раз бросаем кубик - от 1 до 6 
	sum2, sum6 := <-achan, <-bchan
	fmt.Println(float32(sum2)/10000.0, float32(sum6)/30000.0 )
}

// 1.5032 3.5065668
