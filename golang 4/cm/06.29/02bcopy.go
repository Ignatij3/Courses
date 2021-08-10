package main

import (
	"os"
	"time"
	"github.com/nsf/termbox-go"
)

type (
	Runner struct { //Бегунок
		x, y int
		vx, vy int
		color termbox.Attribute
	}	
	
	Game struct { //Перерисовывает
		rs []Runner //Слайс Runner
		delay time.Duration 
	}	
)

func NewRunner (x, y int, vx, vy int, color termbox.Attribute) Runner {
	var r Runner
	r.x, r.y = x, y
	r.vx, r.vy = vx, vy
	r.color = color	
	return r
}	

func (r *Runner) Steep()  {
	width, height := termbox.Size()
	termbox.SetCell(r.x, r.y, ' ', termbox.ColorDefault, termbox.ColorDefault)
	r.x += r.vx
	r.y += r.vy
	termbox.SetCell(r.x, r.y, ' ', termbox.ColorDefault, r.color)
	if r.x<=0 || r.x>=width-1 {
		r.vx = -r.vx
	}	
	if r.y<=0 || r.y>=height-1 {
		r.vy = -r.vy
	}	
}	

func (g *Game) Step () {
	for i:=0; i < len(g.rs); i++ {
		g.rs[i].Steep()
	}	
}

func (g *Game) Run () {
	for {
		g.Step()
		termbox.Flush()
		time.Sleep(g.delay * time.Millisecond)
	}	
}

func (g *Game) Add (r Runner) {
	g.rs = append(g.rs, r)
}	


func main() {
	if err := termbox.Init(); err != nil {
		os.Exit(1) // Ошибка инициализации termbox
	}
	defer termbox.Close()
	
	width, height := termbox.Size()
	termbox.HideCursor()
	var p Game 
	p.Add (NewRunner(width/2, height/2, 1, -1, termbox.ColorGreen))
	p.Add (NewRunner(width-4, height-2, -1, -1, termbox.ColorRed))
	p.delay = 30
	go p.Run()
	time.Sleep(5*time.Second)
}
