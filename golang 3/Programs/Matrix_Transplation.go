package main
import (
		"os"
		"fmt"
		)

func Reverse (x []int)  {
	j := len(x) - 1
	for i := 0; i < j; i++ {
		x[i], x[j] = x[j], x[i]
		j--
	}
}

func main() {
	File, _ := os.Open("../Files/a1.res")
	defer File.Close()
	var Check int
	fmt.Fscanln(File, &Check)
	Array := make([]int, Check, Check)
	for i := 0; i < Check; i++ {
		if _, err := fmt.Fscan(File, &Array[i]); err != nil {
			break
		}
	}
	fmt.Println("Original:")
	fmt.Println(Array)
	Reverse(Array)
	fmt.Println("Reversed:")
	fmt.Println(Array)
}
