package main

import (
	"os"
	"time"
	"github.com/nsf/termbox-go"
)

type (
	Runner struct {
		x, y int
		vx, vy int
		color termbox.Attribute
		delay time.Duration 
	}	
)

func NewRunner (x, y int, vx, vy int, color termbox.Attribute, delay time.Duration) Runner {
	var r Runner
	r.x, r.y = x, y
	r.vx, r.vy = vx, vy
	r.color = color	
	r.delay = delay
	return r
}	

func (r *Runner) Step()  {
	width, height := termbox.Size()
	termbox.SetCell(r.x, r.y, ' ', termbox.ColorDefault, termbox.ColorDefault)
	r.x += r.vx
	r.y += r.vy
	termbox.SetCell(r.x, r.y, ' ', termbox.ColorDefault, r.color)
	time.Sleep(r.delay * time.Millisecond)
	if r.x<=0 || r.x>=width-1 {
		r.vx = -r.vx
	}	
	if r.y<=0 || r.y>=height-1 {
		r.vy = -r.vy
	}	
}	

func (r *Runner) Run () {
	for {
		r.Step()
		termbox.Flush()
	}	
}
	
func main() {
	if err := termbox.Init(); err != nil {
		os.Exit(1) // Ошибка инициализации termbox
	}
	defer termbox.Close()

	width, height := termbox.Size()
	r1:= NewRunner(width/2, height/2, 1, -1, termbox.ColorGreen, 20)
	go r1.Run()
	time.Sleep(5*time.Second)
}
