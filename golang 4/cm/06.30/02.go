package main

import (
	"time"
	"github.com/nsf/termbox-go"
)

const animationSpeed = 30 * time.Millisecond

type (
	Runner struct {
		x, y int
		vx, vy int
		color termbox.Attribute
	}	
)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc+termbox.InputMouse)
	
	eventQueue := make(chan termbox.Event) //Передаёт termbox.Event

	width, height := termbox.Size()
	termbox.HideCursor()
	r:= NewRunner(width/2, height/2, termbox.ColorYellow)

	r.Show()
	termbox.Flush()
	
	var vx1, vy1 int
	go readkey(eventQueue)
	for {
		select {
		case ev := <-eventQueue:
			switch  {
			case ev.Type == termbox.EventMouse:
				switch {
				case ev.Key == termbox.MouseLeft:
					vx1, vy1 = r.vx, r.vy
					r.vx, r.vy = 0, 0
				case ev.Key == termbox.MouseRelease:
					r.vx, r.vy = vx1, vy1
				}	
			case ev.Type == termbox.EventKey:
				switch {
				case ev.Key == termbox.KeyArrowLeft:
					r.vx, r.vy = -1, 0
				case ev.Key == termbox.KeyArrowRight:
					r.vx, r.vy = 1, 0
				case ev.Key == termbox.KeyArrowUp:
					r.vx, r.vy = 0, -1
				case ev.Key == termbox.KeyArrowDown:
					r.vx, r.vy = 0, 1
				case ev.Key == termbox.KeySpace:
					r.vx, r.vy = 0, 0
				case ev.Key == termbox.KeyEsc || ev.Key == termbox.KeyCtrlC:
					return
				}
			}
		default:
			r.Step()
			termbox.Flush()
			time.Sleep(animationSpeed)
		}
	}
}

func readkey(eventQueue chan termbox.Event) {
	for {
		eventQueue <- termbox.PollEvent()
	}
}

func NewRunner (x, y int, color termbox.Attribute) (r Runner) {
	r.x, r.y = x, y
	r.vx, r.vy = 0, 0
	r.color = color	
	return 
}	

func (r *Runner) Step()  {
	if r.vx==0 && r.vy==0 {
		return
	}
	width, height := termbox.Size()
	xnew, ynew := r.x + r.vx, r.y + r.vy
	if xnew<0 || xnew>=width || ynew<0 || ynew>=height {
		return
	}	
	r.Hide()
	r.x, r.y = xnew, ynew
	r.Show()
}	

func (r *Runner) Hide()  {
	termbox.SetCell(r.x, r.y, ' ', termbox.ColorDefault, termbox.ColorDefault)
}	

func (r *Runner) Show()  {
	termbox.SetCell(r.x, r.y, ' ', termbox.ColorDefault, r.color)
}	
