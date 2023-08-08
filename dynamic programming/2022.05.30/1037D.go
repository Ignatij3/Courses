package main

import (
	"bufio"
	"fmt"
	"os"
)

type graph [][]int
type queue [][]int

func insertInOrder(arr *[]int, n int) {
	pos := findBinary(*arr, n)
	if pos == -1 {
		*arr = append(*arr, n)
	} else {
		*arr = append((*arr)[:pos], append([]int{n}, (*arr)[pos:]...)...)
	}
}

func findBinary(arr []int, target int) int {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func getData() (queue, []int) {
	var (
		nodeID, id int
		nodeAmt    int
	)

	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d\n", &nodeAmt)

	mygraph := make(graph, nodeAmt)
	for i := 0; i < nodeAmt-1; i++ {
		fmt.Fscanf(reader, "%d %d\n", &nodeID, &id)
		insertInOrder(&mygraph[nodeID-1], id)
	}

	myqueue := queue{{1}}
	order := make([]int, nodeAmt)
	for i := 0; i < nodeAmt; i++ {
		fmt.Fscanf(reader, "%d", &order[i])
		if len(mygraph[order[i]-1]) != 0 {
			myqueue = append(myqueue, mygraph[order[i]-1])
		}
	}

	return myqueue, order
}

func main() {
	myqueue, order := getData()
	res := calculate(myqueue, order)

	if res {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func calculate(myqueue queue, order []int) bool {
	return isPossibleToTraverseInOrder(myqueue, order)
}

func isPossibleToTraverseInOrder(myqueue queue, order []int) bool {
	for i := range myqueue {
		for range myqueue[i] {
			if findBinary(myqueue[i], order[0]) == -1 {
				return false
			}
			order = order[1:]
		}
	}
	return true
}
