package main

import (
	"os"
	"fmt"
	"time"
	"sync"
	"math/rand"
	"github.com/nsf/termbox-go"
)

type (
	Workers struct {
		width int
		next int
		height int
		prev int
	}
)

var jobs = make (chan Workers, 20)
var finish bool

func Screen() {
	width, height := termbox.Size()
	var prev int
	for i := 0; i < height; {
		prev = i
		i++
		worker := Workers{width, i, height, prev}
		jobs <- worker
	}
	//close(jobs)
}

func WorkerPool(n int) {
	var group sync.WaitGroup
	for i := 0; i < n; i++ {
		if len(jobs) == 0 {finish = true}
		group.Add(1)
		go Work(&group)
		termbox.Flush()
	}
	group.Wait()
}

func Work(group *sync.WaitGroup) {
	color := Color()
	for w := range jobs {
		for xp := 0; xp <= w.width - 1; xp++ {
			termbox.SetCell(xp, w.prev, ' ', termbox.ColorDefault, color)
			time.Sleep(30 * time.Millisecond)
		}
	}
	termbox.Flush()
	group.Done()
}

func Color() termbox.Attribute {
	color := termbox.ColorRed + termbox.Attribute(rand.Intn(7) + 1)
	if rand.Intn(1) == 0 {color |= termbox.AttrBold}
	return color
}

func main() { //no idea how to fix
	if err := termbox.Init(); err != nil {os.Exit(1)}
	defer termbox.Close()
	var end Workers
	for !finish {
		go Screen()
		WorkerPool(10)
		//fmt.Println("POINT_01")
	}
	termbox.Flush()
	time.Sleep(300 * time.Second)
	fmt.Println(end)
}
