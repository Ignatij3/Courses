package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	N    = 1000
	SIZE = 1000
)

type (
	testcase struct {
		data     []int
		requests []request
	}
	request struct {
		a, b int
	}
)

func initData() []int {
	rand.Seed(time.Now().UnixNano())
	var size int = rand.Intn(SIZE) + 1

	arr := make([]int, size)
	for i := 0; i < size; i++ {
		// GENERATE TEST DATA HERE
	}

	return arr
}

func generateTestcase() testcase {
	arr := initData()
	test := testcase{
		data:     arr,
		requests: make([]request, rand.Intn(SIZE)+1),
	}

	for i := range test.requests {
		test.requests[i] = request{ /*GENERATE RANDOM REQUESTS HERE*/ }
	}

	return test
}

func smartTest(test testcase) []int {
	res := make([]int, 0)

	// EFFICIENT FUNCTION HERE

	return res
}

func stupidTest(test testcase) []int {
	res := make([]int, 0)

	// INEFFICIENT FUNCTION HERE

	return res
}

func compare(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main() {
	var (
		correct             bool
		resSmart, resStupid []int
	)

	for i := 0; i < N; i++ {
		test := generateTestcase()
		resSmart = smartTest(test)
		resStupid = stupidTest(test)
		correct = compare(resSmart, resStupid)

		if !correct {
			fmt.Printf("Error on test %d\n", i)
			fmt.Printf("test data: %v\nanswer: %v\nresult: %v\n\n", test.data, resStupid, resSmart)
			fmt.Println("requests:")
			for _, req := range test.requests { // CUSTOM, VERBOSE TESTCASE AND ANSWER OUTPUT
				fmt.Print(req)
				fmt.Print(resStupid)
				fmt.Print(resSmart)
			}
			return
		}
	}
	fmt.Printf("All %d tests passed!", N)
}
