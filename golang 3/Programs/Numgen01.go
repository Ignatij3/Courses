package main

import  (
	"fmt"
	"os"
	"math/rand"
	"time"
)

func main() {
	var (
		i, j int
	)
	rand.Seed(time.Now().UnixNano())
	File, err := os.Create("../Files/NumInsane.dat")
	if err != nil {fmt.Println(err)}
	defer File.Close()
	for i = 0; i <= 1000000; i++ {
		for j = rand.Intn(200000) - 100000; ; {
			_, err =File.WriteString(fmt.Sprintf("%d\n", j))
			if err != nil {fmt.Println(err)}
			i++
			break
		}
	}
}
