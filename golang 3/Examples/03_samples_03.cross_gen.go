package main

import (
	"fmt"
	"time"
	"math/rand"
	"os"
)

func main() {
	rand.Seed(int64(time.Now().Nanosecond()))
	f, _ := os.Create("cross.dat")
	defer f.Close()

	n:= rand.Intn(25)
	x:= rand. Intn(20)
	for i := 0; i < n; i++ {
		fmt.Fprint(f, x, " ")
		x += rand.Intn(5) + 1
	}
	fmt.Fprintln(f)

	n = rand.Intn(25)
	x = rand.Intn(20)
	for i := 0; i < n; i++ {
		fmt.Fprint(f, x, " ")
		x += rand.Intn(5) + 1
	}
	fmt.Fprintln(f)
}
