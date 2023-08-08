package main

import "fmt"

func quit(message string) {
	fmt.Println(message)
}

func printList(nums []int) {
	defer quit("Finished printing")
	defer fmt.Print("...")
	fmt.Print("Started printing...")
	for _, v := range nums {
		fmt.Print(v, " ")
	}
}

func main() {
	defer quit("Bye")
	nums := []int{78, 109, 2, 563, 300}
	printList(nums)
	defer quit("Bye-bye")
}
//	Started printing...78 109 2 563 300 ...Finished printing
//	Bye-bye
//	Bye
