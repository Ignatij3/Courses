package main

import (
	"os"
	"github.com/nsf/termbox-go"
)

type
	board struct {
		field [9]rune	// состяние игрового поля
						// каждая клетка содердит X, 0 или точку
						// точка соответствует пустой клетке
		xc, yc int		// координаты центральной клетки в окне
	}
		
func InitBoard(x, y int) board {
	return board{ field: [9]rune{'.','.','.','.','.','.','.','.','.'}, xc: x, yc: y}
	//  0 1 2
	//  3 4 5
	//  6 7 8
}	

func (a board) Output() {
	termbox.SetCell(a.xc-3, a.yc-3, a.field[0], termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(a.xc  , a.yc-3, a.field[1], termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(a.xc+3, a.yc-3, a.field[2], termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(a.xc-3, a.yc  , a.field[3], termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(a.xc  , a.yc  , a.field[4], termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(a.xc+3, a.yc  , a.field[5], termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(a.xc-3, a.yc+3, a.field[6], termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(a.xc  , a.yc+3, a.field[7], termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(a.xc+3, a.yc+3, a.field[8], termbox.ColorDefault, termbox.ColorDefault)
	termbox.Flush()
}	

func (a board) Result() rune {
	// Возвращает:
	// X, если выиграли X
	// 0, если выиграли 0
	// =, если победителя нет, а доска заполнена  - ничья
	// C, если победителя нет, а на доске есть пустые клетки - игра продолжается
	switch {
		case a.field[0]==a.field[1] && a.field[1]==a.field[2] && a.field[0]!='.': return a.field[0]
		case a.field[3]==a.field[4] && a.field[4]==a.field[5] && a.field[3]!='.': return a.field[3]
		case a.field[6]==a.field[7] && a.field[7]==a.field[8] && a.field[6]!='.': return a.field[6]
		case a.field[0]==a.field[3] && a.field[3]==a.field[6] && a.field[0]!='.': return a.field[0]
		case a.field[1]==a.field[4] && a.field[4]==a.field[7] && a.field[1]!='.': return a.field[1]
		case a.field[2]==a.field[5] && a.field[5]==a.field[8] && a.field[2]!='.': return a.field[2]
		case a.field[0]==a.field[4] && a.field[4]==a.field[8] && a.field[0]!='.': return a.field[0]
		case a.field[2]==a.field[4] && a.field[4]==a.field[6] && a.field[2]!='.': return a.field[2]
	}	
	for _, c:= range(a.field) {
		if c == '.' { return 'C'}  // game continues
	}	
	return '='  // draw
}

func (a board) Negative() board {
	// Возвращает "негатив" доски - крестики заменяются на нолики и наоборот
	res:= a
	for i, c := range(res.field) {
		if c == 'X' {
			res.field[i] = '0'
		} else 		
		if c == '0' {
			res.field[i] = 'X'
		}
	}
	return res	
}	

func (a *board) GetMove (sign rune) {  // sign = 'X' or '0'
	// Ожидает хода игрока - нажатия мышкой на пустую клетку игрового поля  
	// Ставит в выбранную клетку соответствующий знак - крестик или нолик
	var (
		exit bool = false 
		pos int
	)	
	for !exit {
		ev := termbox.PollEvent()
		if ev.Type == termbox.EventMouse && ev.Key == termbox.MouseRelease {
			exit = true
			dx, dy:= ev.MouseX - (*a).xc, ev.MouseY - (*a).yc
			switch {
				case dx>=-4 && dx<=-2 && dy>=-4 && dy<=-2: pos = 0
				case dx>=-1 && dx<= 1 && dy>=-4 && dy<=-2: pos = 1
				case dx>= 2 && dx<= 4 && dy>=-4 && dy<=-2: pos = 2
				case dx>=-4 && dx<=-2 && dy>=-1 && dy<= 1: pos = 3
				case dx>=-1 && dx<= 1 && dy>=-1 && dy<= 1: pos = 4
				case dx>= 2 && dx<= 4 && dy>=-1 && dy<= 1: pos = 5
				case dx>=-4 && dx<=-2 && dy>= 2 && dy<= 4: pos = 6
				case dx>=-1 && dx<= 1 && dy>= 2 && dy<= 4: pos = 7
				case dx>= 2 && dx<= 4 && dy>= 2 && dy<= 4: pos = 8
				default: exit = false
			}	
		}
		if exit && (*a).field[pos] != '.' {
			exit = false
		}
	}
	(*a).field[pos] = sign
}	

func (a board) Estimate() (value rune, bestMove int) {
	// Оценочная функция
	value = a.Result()
	if value=='X' || value=='0' || value=='=' { return value, -1 }
	value = '0'
	bestMove = 0
	for i, c:= range(a.field) {
		if c == '.' {
			a.field[i] = 'X'	
			e, _:= a.Negative().Estimate()
			a.field[i] = '.'	
			switch e {
			case '=': 
				value = '='
				bestMove = i
			case '0': 
				return 'X', i
			}	
		}		
	}	 
	return value, bestMove
}	

func GetFirst() (First, Second string) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	w, h := termbox.Size()
	OutputText("Who starts the game?", w/2 - 10, h/2 - 2)
	OutputText("Human", w/2 - 10, h/2 + 2)
	OutputText("Computer", w/2 + 5, h/2 + 2)
	for {
		ev := termbox.PollEvent()
		if ev.Type == termbox.EventMouse && ev.Key == termbox.MouseLeft {
			if ev.MouseY >= h/2 {
				if ev.MouseX > w/2 {
					termbox.Clear( termbox.ColorDefault, termbox.ColorDefault)
					return "Computer", "Human"
				}	
				if ev.MouseX < w/2 {
					termbox.Clear( termbox.ColorDefault, termbox.ColorDefault)
					return "Human", "Computer"
				}	
			}	
		}
	}	
}	

func OutputText(Message string, x0, y0 int) {
	x, y:= x0, y0
	for _, c:= range([]rune(Message)) {
		termbox.SetCell(x, y, c, termbox.ColorDefault, termbox.ColorDefault)
		x++
	}
	termbox.Flush()		
}
	
func main() {
	err := termbox.Init()
	if err != nil { // Ошибка инициализации termbox
		os.Exit(1)
	}
	defer termbox.Close()
	termbox.HideCursor()
	termbox.SetInputMode(termbox.InputEsc + termbox.InputMouse)

	First, Second := GetFirst()
	currentSign:= 'X'
	activePlayer:= First

	w, h := termbox.Size()
	f:= InitBoard(w/2, h/2)	// состояние игры: положение на 
							// поле и координаты центра доски 
	f.Output()

	for f.Result()=='C' {
		if activePlayer == "Human" {
			f.GetMove(currentSign)
			f.Output()
		} else {
		// activePlayer	 == "Computer"
			var bestMove int
			if currentSign == 'X' {
				_, bestMove = f.Estimate()
			} else {
				_, bestMove = f.Negative().Estimate()
			}				
			f.field[bestMove] = currentSign
		}		
		f.Output()
		if currentSign == 'X' {
			currentSign = '0'
		} else {
			currentSign = 'X'
		}	 	 	
		if activePlayer == "Human" {
			activePlayer = "Computer"
		} else {	
			activePlayer = "Human"
		}
	}
	var WinMessage string
	switch f.Result() {
	case 'X': 
		WinMessage = First + " won"
	case '0': 
		WinMessage = Second + " won"
	case '=': 
		WinMessage = "Draw"
	}
	OutputText(WinMessage, f.xc - len([]rune(WinMessage))/2, f.yc + 8)
	termbox.PollEvent()
}
