package main

import  (
	"fmt"
	"os"
	"math/rand"
	"time"
)
												//Заполняет файл неубывающими числами, кол-во которых генерируется относительно рандомно.
func main() {
	var (
		random, k, i, j int
	)
	rand.Seed(time.Now().UnixNano())
	File, err := os.Create("../Files/Numbers.dat")
	if err != nil {fmt.Println(err)}
	defer File.Close()
	for i = 0; i <= 100000; i++ {
		k = 0
		random = i
		for j = rand.Intn(100); k < j; k++ {
			fmt.Println(random)
			_, err =File.WriteString(fmt.Sprintf("%d\n", random))
        if err != nil {fmt.Println(err)}
		}
	}
}
