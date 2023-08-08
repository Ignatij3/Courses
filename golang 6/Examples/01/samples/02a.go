package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(int64(time.Now().Nanosecond()))
	fmt.Printf("Int100: %d %d %d\n", rand.Intn(100), rand.Intn(100), rand.Intn(100))
}
