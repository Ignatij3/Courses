package main

import (
	"os"
	"time"
	"math/rand"
	"github.com/nsf/termbox-go"
	"github.com/veandco/go-sdl2/sdl"
)

type (
	Runner struct {
		x, y int
		vx, vy int
		char rune
		delay int    // ms
		lifeTime int // ms
	}	
)

func NewRunner (x, y int, vx, vy int, ch rune, delay int, lifeTime int) Runner {
	return Runner{x, y, vx, vy, ch, delay, lifeTime}
}	

func (r *Runner) Step()  {
	width, height := termbox.Size()
	sdl.Do(func() {
		termbox.SetCell(r.x, r.y, ' ', termbox.ColorDefault, termbox.ColorDefault)
	}) 
	r.x += r.vx
	r.y += r.vy
	sdl.Do(func() {
		termbox.SetCell(r.x, r.y, r.char, termbox.ColorDefault, termbox.ColorDefault)
	})	
	time.Sleep(time.Duration(r.delay) * time.Millisecond)
	if r.x<=0 || r.x>=width-1 {
		r.vx = -r.vx
	}	
	if r.y<=0 || r.y>=height-1 {
		r.vy = -r.vy
	}	
	sdl.Do(func() {
		termbox.Flush()
	})
}	

func (r Runner) Run (quit chan int) {
	start:= time.Now()
	for t:= start; t.Sub(start) < time.Duration(r.lifeTime)*time.Millisecond; t = time.Now()  {
		r.Step()
	}	
	sdl.Do(func() {
		termbox.SetCell(r.x, r.y, ' ', termbox.ColorDefault, termbox.ColorDefault)
		termbox.Flush()
	}) 
	quit <- 1
}

func traffic() {
	width, height := termbox.Size()
	var rs []Runner
	var waiting []chan int
	for c:= 'A'; c <='Z'; c++  {
		rs = append(rs, NewRunner(rand.Intn(width), rand.Intn(height), 1, 1, c, 20 + 5*int(c-'A'), 3000 + rand.Intn(int(c-'A')+10)*500))
		waiting = append(waiting, make(chan int, 1))
	}
	for i, r:= range rs  {
		go r.Run(waiting[i])
	}
	for i, _:= range rs  {
		<-waiting[i]
	}
}	

func main() {
	if err := termbox.Init(); err != nil {
		os.Exit(1) // Ошибка инициализации termbox
	}
	defer termbox.Close()
	termbox.HideCursor()
    rand.Seed(int64(time.Now().Nanosecond()))
    
	sdl.Main(func() {
		traffic()
	})
}
