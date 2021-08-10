package main

import (
	"os"
	"time"
	"github.com/nsf/termbox-go"
)

func main() {
	if err := termbox.Init(); err != nil {
		os.Exit(1) // Ошибка инициализации termbox
	}
	defer termbox.Close()

	width, height := termbox.Size()
	
	x, y, vx, vy := width/2, height/2, 1, -2
	
	for t:= time.Now(); time.Since(t) < 5*time.Second;	{
		termbox.SetCell(x, y, ' ', termbox.ColorDefault, termbox.ColorDefault)
		x += vx
		y += vy
		termbox.SetCell(x, y, ' ', termbox.ColorDefault, termbox.ColorGreen)
		termbox.Flush()
		time.Sleep(30 * time.Millisecond)
		if x<=0 || x>=width-1 {
			vx = - vx
		}	
		if y<=0 || y>=height-1 {
			vy = - vy
		}	
	}
}

