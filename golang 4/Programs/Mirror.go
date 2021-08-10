package main

import (
	"time"
	"math/rand"
	"github.com/nsf/termbox-go"
)

type (
	Runner struct {
		x, y int
		vx, vy int
		color termbox.Attribute
	}
)

func NewRunner(x, y int, color termbox.Attribute) (r Runner) {
	r.x, r.y = x, y
	r.vx, r.vy = 0, 0
	r.color = color
	return
}

func main() {
	termbox.Init()
	defer termbox.Close()
	width, _ := termbox.Size()
	
	termbox.SetInputMode(termbox.InputEsc)
	termbox.HideCursor()
	r := NewRunner(0, 0, termbox.ColorBlue)
	Line()
	termbox.SetCell(0, 0, ' ', termbox.ColorDefault, r.color)
	termbox.SetCell(width - 1, 0, ' ', termbox.ColorDefault, r.color)
	
	event := make(chan termbox.Event)
	termbox.Flush()
	
	go readkey(event)
	for {
		select {
		case key := <-event:
			switch  {
			case key.Type == termbox.EventKey:
				switch {
				case key.Key == termbox.KeyArrowLeft:
					r.vx, r.vy = -1, 0
				case key.Key == termbox.KeyArrowRight:
					r.vx, r.vy = 1, 0
				case key.Key == termbox.KeyArrowUp:
					r.vx, r.vy = 0, -1
				case key.Key == termbox.KeyArrowDown:
					r.vx, r.vy = 0, 1
				case key.Key == termbox.KeyEsc || key.Key == termbox.KeyCtrlC:
					return
				}
			}
		default:
			r.Move()
			termbox.Flush()
			time.Sleep(30 * time.Millisecond)
		}
	}
}

func (r *Runner) Move() {
	//buffer := termbox.CellBuffer()
	//if termbox.SetCursor(xnew, ynew) == termbox.ColorWhite {return} -- Как считывать текущее положение "ползунка"? Это последний вариант, который я попробовал
	if r.vy == 0 && r.vx == 0 {return}
	
	width, height := termbox.Size()
	xnew1, ynew1 := r.x + r.vx, r.y + r.vy
	xnew2, ynew2 := width - xnew1 - 1, ynew1
	bold := FColor()
	
	if xnew1 == width/2 {return}
	if xnew1 < 0 || ynew1 < 0 || xnew1 >= width/2 - 1 || ynew1 >= height {return}
	
	r.x, r.y = xnew1, ynew1
	posx2, posy2 := xnew2, ynew2
	
	termbox.SetCell(r.x, r.y, ' ', termbox.ColorDefault, bold)
	termbox.SetCell(posx2, posy2, ' ', termbox.ColorDefault, bold)
}

func Line() {
	width, height := termbox.Size()
	white := termbox.ColorWhite + termbox.AttrBold
	for i := 0; i < height; i++ {
		termbox.SetCell(width/2, i, ' ', termbox.ColorDefault, white)
		termbox.SetCell(width/2 - 1, i, ' ', termbox.ColorDefault, white)
	}
}

func readkey(event chan termbox.Event) {
	for {
		event <- termbox.PollEvent()
	}
}

func FColor() termbox.Attribute {
	bol := termbox.ColorBlack + 1 + termbox.Attribute(rand.Intn(6))
	bol |= termbox.AttrBold
	return bol
}
