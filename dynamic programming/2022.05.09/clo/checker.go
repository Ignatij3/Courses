package main

import (
	"bufio"
	"fmt"
	"os"
)

func getData() []int {
	var (
		state string
		num   int
	)

	buf := bufio.NewReader(os.Stdin)

	fmt.Fscanf(buf, "%s\n", &state)
	if state == "NIE" {
		return nil
	}

	arrows := make([]int, 0, 1000)
	for {
		if _, err := fmt.Fscanf(buf, "%d\n", &num); err != nil {
			break
		}
		arrows = append(arrows, num)
	}

	return arrows
}

func main() {
	arrows := getData()
	if arrows == nil {
		fmt.Println("true")
	} else {
		fmt.Println(isValid(arrows))
	}
}

func isValid(arrows []int) bool {
	for i, arrow := range arrows {
		if i+1 == arrows[arrow-1] {
			return false
		}
	}
	return true
}
