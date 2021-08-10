package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type tResult [13]float64

const  (
	step = 1000  // проверяем концовку каждые step раз
	epsilon = 1.0e-6
)	
	
func finish (old tResult, new tResult) bool  {
	for i:= 2; i <= 12; i++  {
		tmp := math.Abs(old[i]-new[i])/new[i]
		if tmp > epsilon { return false }
	}	
	return true
}	 	
	
func main() {
	var	(
		counter [13]int
		freq0, freq tResult
		n, a int
	)	
	
	rand.Seed(time.Now().UnixNano())
	//  вариант:
	//  rand.Seed(int64(time.Now().Nanosecond()))
	
	for  {
		a = 1 + rand.Intn(6) + 1 + rand.Intn(6)
		counter[a]++
		n++
		if  n % step == 0  {
			for i:= 2; i <= 12; i++  {	
				freq0[i] = freq[i]
				freq[i] = float64(counter[i])/float64(n)
			}	
			if  finish(freq0, freq)  { break }
		}	
	}
	
	for i:= 2; i <= 12; i++  {
		fmt.Printf("%2d. %8.6f\n", i, freq[i])	
	}	
}
