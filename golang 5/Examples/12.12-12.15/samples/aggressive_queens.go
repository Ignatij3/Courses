package main

import "fmt"

const n = 8

type
	cell struct {
		row, col int
	}	
	
func (c cell) Connected(c2 cell) bool {
	// Соединены ли клетки c и c2 ходом ферзя?
	if c == c2 { return false }
	return	(c.row == c2.row) || (c.col == c2.col) ||
			(abs(c.row-c2.row) == abs(c.col-c2.col))
}		

func (c cell) Print() {
	fmt.Printf("%c%d ", c.col+'a', c.row+1)
}

func (c cell) Next() cell {
	// Возвращает клетку, следующую за клеткой c.
	// Направление движения: вдоль столбца - увеличиваем строку, 
	// в конце столбца переходим на нижнюю клетку следующего столбца
	if c.row < n-1  {
		return cell{c.row + 1, c.col}
	} else {
		return cell{0, c.col + 1}	
	}
}			

func (c cell) Terminal() bool {
	// Верно ли, что c - последняя клетка на доске?
	return c.Next().col == n
}	
	
func Success(list []cell) bool {
	// Верно ли, что все клетки доски находятся 
	// под боем какого-то ферзя из списка list 
	for row:= 0; row < n; row++ {
		for col:= 0; col < n; col++ {
			ok := false
			for _, c:= range(list) {
				if (cell{row, col}).Connected(c) {
					ok = true
					break
				}	
			}	
			if !ok {
				return false
			}	
		}
	}			
	return true
}	 	

var result []cell  // здесь храним текущее наилучшее решение

func search(list []cell) {
	if list[0].col == n-1 && list[0].row == n - len(list) {
	// терминальный случай: последняя комбинация ферзей,
	// дальше двигаться некуда
		return 
	}	
	if len(list) >= len(result)-1 {
	// добавлять ферзей бессмысленно - улучшить результат не удастся
		return
	}	
	// last - последий ферзь ф текущем списке
	last:= list[len(list) - 1]
	// добавляем ещё одного ферзя 
	for c:= last.Next(); !c.Terminal(); c = c.Next() {
		if Success( append(list, c) ) {
		// если новый ферзь делает список таким, что все
		// все клетки находятся под боем, то этот список	
		// улучшает текущий результат - запоминаем его
			result = append(list, c)
			return
		} 	
		// если новый ферзь не далает список таким,
		// что все клетки находятся под боем,
		// то пытаемся добавить ещё ферзей 
		search( append(list, c)	)
	}	
}	

func main() {
	// Начальное решение - заполняем ферзями весь нижний ряд
	for i:= 0; i< n; i++ {
		result = append(result, cell{0, i} )
	}	
	// Поиск начинается со списка из одного ферзя,
	// стоящего в первой клетке - клетке {0, 0}
	search([]cell{cell{0,0}})
	// Печать результата
	for _, c := range (result) {
		c.Print()
	}
	fmt.Println()	
}

func abs(x int) int {
	if x<0 { 
		return -x 
	} else {
		return x
	}
}			
