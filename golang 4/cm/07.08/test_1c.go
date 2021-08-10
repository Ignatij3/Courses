package main

import "fmt"

func plusTwo() (func(v int) (int)) {
	return func(v int) (int) {
		return v+2
	}
}

func plusX(x int) (func(v int) (int)) {
	return func(v int) (int) {
		return v+x
	}
}

func main() {
	p := plusTwo()
	fmt.Printf("3+2 = %d\n", p(3))		//	3+2 = 5
	px := plusX(7)
	fmt.Printf("11+7 = %d\n", px(11))	//	11+7 = 18
	fmt.Printf("5+7 = %d\n", px(5))		//	5+7 = 12
}
