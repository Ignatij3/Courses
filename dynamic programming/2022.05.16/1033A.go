package main

import (
	"fmt"
)

type point struct {
	x, y int
}

func getData() (point, point, point) {
	var (
		dimensions int
		queen, king, goal point
	)


	fmt.Scanf("%d\n", &dimensions)
	fmt.Scanf("%d %d\n", &queen.x, &queen.y)
	fmt.Scanf("%d %d\n", &king.x, &king.y)
	fmt.Scanf("%d %d\n", &goal.x, &goal.y)

	return queen, king, goal
}

func main() {
	queen, king, goal := getData()
	res := calculate(queen, king, goal)
	if res {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func calculate(queen, king, goal point) bool {
	if king.x < queen.x && queen.x < goal.x {
		return false
	} else if king.x > queen.x && queen.x > goal.x {
		return false
	} else if king.y < queen.y && queen.y < goal.y {
		return false
	} else if king.y > queen.y && queen.y > goal.y {
		return false
	}

	return true
}