package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	start1 := time.Now()
	get_list1()
	fmt.Println(time.Since(start1))

	start2 := time.Now()
	get_list2()
	fmt.Println(time.Since(start2))

	start3 := time.Now()
	get_list3()
	fmt.Println(time.Since(start3))

	start4 := time.Now()
	get_list4()
	fmt.Println(time.Since(start4))
}

func get_list1() {
	var (
		abonents int
		plans    int
	)

	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d %d\n", &abonents, &plans)
	tariffs := make([]int, abonents)

	for i := 0; i < abonents; i++ {
		fmt.Fscanf(reader, "%d ", &tariffs[i])
	}
}

func get_list2() {

}

func get_list3() {

}

func get_list4() {

}
