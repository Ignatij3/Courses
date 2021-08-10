package main

import (
	"fmt"
	"time"
	"math/rand"
	"os"
)

func main() {
	rand.Seed(int64(time.Now().Nanosecond()))
	f, _ := os.Create("separate.dat")
	defer f.Close()

	n:= rand.Intn(25) + 12
	for i := 0; i < n; i++ {
		fmt.Fprint(f, rand.Intn(150), " ")
	}	
	fmt.Fprintln(f)
}
