package main

import (
	"fmt"
	"math/rand"
	"time"
	"os"
)


func main() {
	rand.Seed(int64(time.Now().Nanosecond()))
	fmt.Print("Enter n: ")
	var n int
	fmt.Scan(&n)
	f, _ := os.Create("a1.dat")
	fmt.Fprintln(f, n)
	defer f.Close()
	for i := 1; i < n; i++ {
		fmt.Fprint(f, rand.Intn(100) + 1, " ")
	}
	fmt.Fprintln(f, rand.Intn(100) + 1)
}
