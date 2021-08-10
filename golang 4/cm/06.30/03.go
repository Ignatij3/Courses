package main

import (
	"math/rand"
	"time"
	"github.com/nsf/termbox-go"
)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	firework()
}

func firework() {
	width, height := termbox.Size()
	start:= time.Now()
	startcounter:= start
	for  {
		current:= time.Now()
		switch  {
		case current.Sub(startcounter) >= 120 * time.Millisecond:
			if rand.Intn(25) == 0  {//Если 0, рисует цифру
				x, y:= rand.Intn(width), rand.Intn(height)	
				go countdown(x, y)
				startcounter = current
			}	
			termbox.Flush()
		case current.Sub(start) >= 10 * time.Second:
			return
		}
	}	
}

func countdown(x, y int) {
	for _, r := range "9876543210" {
		termbox.SetCell(x, y, r, termbox.ColorDefault, termbox.ColorDefault)
		time.Sleep(200 * time.Millisecond)
	}
	termbox.SetCell(x, y, ' ', termbox.ColorDefault, termbox.ColorDefault)
}
