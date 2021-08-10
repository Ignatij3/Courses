package main

import (
	"strings"
	"strconv"
	"fmt"
)

const (
	dim = 4
	N = dim * dim
)

func MakeCells(cells *[N]bool) {
	for k := 0; k < N; k++ {
		(*cells)[k] = true
	}
}

func PrintCells(a [N]bool) {
	for n, _ := range a {
		if a[n] {
			if n % dim != dim - 1 {
				fmt.Printf("%2d ", n)
			} else {
				fmt.Printf("%2d\n", n)
			}
		} else if !a[n] {
			if n % dim != dim - 1 {
				fmt.Print("-- ")
			} else {
				fmt.Println("--")
			}
		}
	}
	fmt.Println()
}

func InSlice(n int, a []int) bool {
	for _, c := range a  {
		if c == n {return true}
	}
	return false
}

func InStraightLine(cells [N]bool, checked []int) bool {
	var firstRed, lastRed int = -1, -1
	for n, c := range cells {
		if c {
			if (InSlice(n, checked) && lastRed != -1) && ((n - lastRed != 1 && n - lastRed != dim) || (n % dim == 0 && n - lastRed == 1)) {return false}
			if InSlice(n, checked) {
				if firstRed == -1 {firstRed = n}
				lastRed = n
			}
		}
	}
	return firstRed >= 0 && (firstRed == lastRed || (lastRed - firstRed <= dim || lastRed % dim == firstRed % dim))
}

func ConvertIntoSlice(delete string) []int {
	var (
		num int
		res []int
	)
	a := strings.Split(delete, ".")
	for n, _ := range a {
		num, _ = strconv.Atoi(a[n])
		res = append(res, num)
	}
	return res
}

func DeleteSelected(cells *[N]bool, del []int) {
	for _, c := range del {
		for n, _ := range *cells {
			if n == c {(*cells)[n] = false}
		}
	}
}

func MakeCopy(a [N]bool, b []int) [N]bool {
	new := a
	for n, _ := range new {
		if InSlice(n, b) {new[n] = false}
	}
	return new
}

func Estimate(cells [N]bool, comp bool) (bool, []int) { //Если comp - true, значит ход компьютера и наборот
	var (
		move, chMove []int
		allNull, winP, winC bool = true, false, false
	)
	for _, c := range cells {if c {allNull = false}}
	if allNull {return true, nil}
	
	for n, _ := range cells {
		if cells[n] {
			for horiz := n; (horiz % dim != 0 && horiz > 0) || horiz == 0; horiz++ { //Проходится по горизонтали от n-ной клетки
				for h1 := horiz; h1 >= n; h1-- { //Идёт от horiz, который постоянно увеличивается до n
					move = append(move, h1)
					if !cells[h1] {move = nil; break} //Если встречает мёртвую клетку, то break'ается, так как дальше идти нет смысла
				}
				if move == nil {break}
				if comp {
					if winP, chMove = Estimate(MakeCopy(cells, move), false); winP {if chMove == nil {return false, move} else {return true, move}}
				} else {
					if winC, chMove = Estimate(MakeCopy(cells, move), true); winC {return true, move}
				}
				move = nil
			}
			for vert := n; vert < N; vert += dim { //Проходится по вертикали от n-ной клетки
				for v1 := vert; v1 >= n; v1 -= dim { //Идёт от vert, который постоянно увеличивается до n
					move = append(move, v1)
					if !cells[v1] {move = nil; break} //Также, как и в 102
				}
				if move == nil {break}
				if comp {
					if winP, chMove = Estimate(MakeCopy(cells, move), false); winP {if chMove == nil {return false, move} else {return true, move}}
				} else {
					if winC, chMove = Estimate(MakeCopy(cells, move), true); winC {return true, move}
				}
				move = nil
			}
		}
	}
	return true, nil //До сюда не доходит, по крайней мере не должно
}

func FindBestMove(cells *[N]bool) {
	win, move := Estimate(*cells, true)
	fmt.Printf("win - %t, move - %v\n", win, move)
	DeleteSelected(cells, move)
}

func WaitForMove(cells [N]bool) {
	var (
		delete string
		delSlice []int
	)
	for {
		PrintCells(cells)
		fmt.Print("Enter numbers of cell you want to delete (with \".\" as a divisor): ")
		fmt.Scan(&delete)
		delSlice = ConvertIntoSlice(delete)
		if InStraightLine(cells, delSlice) {
			DeleteSelected(&cells, delSlice)
			FindBestMove(&cells)
		}
		delSlice = delSlice[:0]
	}
}

func main() {
	var (
		fMove string
		cells [N]bool
	)
	
	fmt.Print("Who moves first? (p/c): ")
	fmt.Scan(&fMove)
	for fMove != "p" && fMove != "c" {
		fmt.Print("Incorrect input, try again: ")
		fmt.Scan(&fMove)
	}
	
	MakeCells(&cells)
	if fMove == "c" {FindBestMove(&cells)}
	WaitForMove(cells)
}
