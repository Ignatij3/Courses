package main

import  (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var (
		j int
	)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i <= 30; i++ {
		j = rand.Intn(20000000000)
		fmt.Println(j)
	}
}
