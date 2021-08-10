package main

import (
	"os"
	"math/rand"
	"github.com/nsf/termbox-go"
)

func randomColor() termbox.Attribute  {
	res:= termbox.ColorBlack + 1 + termbox.Attribute(rand.Intn(7))
	if rand.Intn(2)==0  {
		res |= termbox.AttrBold  // - добавляем цвету яркость
	}
	return res	
}	

func fillBox()  {
	width, height := termbox.Size()
	for row := 0; row < height-1; row++ {
		for col:= 0; col < width; col++ {
			termbox.SetCell(col, row, ' ', termbox.ColorDefault, randomColor())
		}	  
	}
	termbox.Flush()
}
	
func getCell (col, row int) termbox.Cell  {
	width, _ := termbox.Size()
	return termbox.CellBuffer()[row*width+col]
}
	
func drawBottomRibbon(color termbox.Attribute)  {
	width, height := termbox.Size()
	for col:= 0; col < width; col++ {
		termbox.SetCell(col, height-1, ' ', termbox.ColorDefault, color)
	}			
	termbox.Flush()
}	
	
func main() {
	if err := termbox.Init(); err != nil {
		// Ошибка инициализации termbox
		os.Exit(1)
	}
	defer termbox.Close()
	// Включаем считывание событий мышки
	termbox.SetInputMode(termbox.InputEsc + termbox.InputMouse)

	loop:  // Основной цикл
	for {
		ev := termbox.PollEvent()	// Ждём и считываем события 
		switch ev.Type  {
		case termbox.EventKey:  // - событие клавиатуры
			// Esc или CtrlC - выход не только из switch, но и из цикла 
			if ev.Key == termbox.KeyEsc || ev.Key == termbox.KeyCtrlC {
				break loop
			}
		case termbox.EventMouse:  // - событие мышки
			switch  ev.Key {
			case termbox.MouseLeft: 
				// Нажата левая кнопка - красим полоску внизу
				// в цвет текущей клетки   
				drawBottomRibbon (getCell(ev.MouseX, ev.MouseY).Bg)
			case termbox.MouseRelease:
				// Отпустили кнопку - чистим полоску внизу
				drawBottomRibbon (termbox.ColorDefault)
			}	
		case termbox.EventResize:  // - изменяем размер окна
			fillBox()
		}
	}
}
