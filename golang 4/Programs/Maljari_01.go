package main

import (
	"os"
	"math/rand"
	"github.com/nsf/termbox-go"
)

func Move() {
	var (
		stat bool
		x, y int
	)
	width, height := termbox.Size()
	for {
		x = rand.Intn(width)
		y = rand.Intn(height)
		if stat {
			termbox.SetCell(x, y, ' ', termbox.ColorDefault, termbox.ColorRed)
			stat = false
		} else {
			termbox.SetCell(x, y, ' ', termbox.ColorDefault, termbox.ColorGreen)
			stat = true
		}
		termbox.Flush()
	}
}

func main() {
	if err := termbox.Init(); err != nil {os.Exit(1)}
	defer termbox.Close()
	Move()
}
