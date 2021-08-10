package main

import (
	"fmt"
	"time"
	"math/rand"
	"github.com/veandco/go-sdl2/sdl"
)

type
	Segment  struct {
		A, B sdl.Point
	}
var (
	winTitle string = "Crossing Lines"
	winWidth, winHeight int32 = 1200, 720
	window *sdl.Window
	renderer *sdl.Renderer
)

func Lines(numL int) []Segment {
	renderer.SetDrawColor(10, 85, 15, 255)
	var s Segment
	mLines := make([]Segment, numL)
	for i := 0; i < numL; i++ {
		s.A = sdl.Point{int32(rand.Intn(int(winWidth))),  int32(rand.Intn(int(winHeight)))}
		s.B = sdl.Point{int32(rand.Intn(int(winWidth))),  int32(rand.Intn(int(winHeight)))}
		mLines = append(mLines, s)
		renderer.DrawLine(s.A.X, s.A.Y, s.B.X, s.B.Y)
		renderer.Present()
		if numL > 100 && numL <= 300 {
			sdl.Delay(20)
		} else if numL > 300 && numL <= 600 {
			sdl.Delay(10)
		} else if numL > 600 && numL <= 1000 {
			sdl.Delay(5)
		} else if numL > 1000 {
			sdl.Delay(2)
		} else {
			sdl.Delay(50)
		}
	}
	return mLines
}

func Parallel(s []Segment) {
	renderer.SetDrawColor(255, 0, 0, 255)
	var delay uint32
	switch {
	case len(s) > 100 && len(s) <= 300 :
		delay = 3
	case len(s) > 300 && len(s) <= 600 :
		delay = 2
	case len(s) > 600 && len(s) <= 1000:
		delay = 1
	case len(s) > 1000:
		delay = 0
	default:
		delay = 20
	}
	for i1 := 0; i1 < len(s); i1++ {
		for  i2 := i1+1; i2 < len(s); i2++ {
			Squares(s[i1], s[i2], delay)
		}
	}
}
	
func Squares(s1, s2 Segment, delayTime uint32) {
	var (
		Vx, Vy, Ux, Uy int32
		t, T, xN, yN float64
		b sdl.Rect 
	)
	Vx = s1.B.X - s1.A.X
	Vy = s1.B.Y - s1.A.Y
	Ux = s2.B.X - s2.A.X
	Uy = s2.B.Y - s2.A.Y
	if Vx * Uy == Vy * Ux {return} else {
		t = float64(((s2.A.X - s1.A.X) * Uy - (s2.A.Y - s1.A.Y) * Ux)) / float64((Vx * Uy - Vy * Ux))
		T = float64(((s2.A.X - s1.A.X) * Vy - (s2.A.Y - s1.A.Y) * Vx)) / float64((Vx * Uy - Vy * Ux))
		if t < 0 || t > 1 || T < 0 || T > 1 {return} else {
			xN = float64(s1.A.X) + float64(Vx) * t
			yN = float64(s1.A.Y) + float64(Vy) * t
			b.X, b.Y, b.W, b.H  = int32(xN) - 5, int32(yN) - 5, 10, 10
			renderer.DrawRect(&b)
			sdl.Delay(delayTime)
		}
	}
	renderer.Present()
}

func main() {
	var (
		event sdl.Event
		numL int
	)
	fmt.Print("Enter amount of lines:")
	fmt.Scan(&numL)
	for numL <= 0 {
		fmt.Print("Error, try again:")
		fmt.Scan(&numL)
	}

	rand.Seed(time.Now().UnixNano())
	window, _ = sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, winWidth, winHeight, sdl.WINDOW_SHOWN)
	renderer, _ = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	defer window.Destroy()
	defer renderer.Destroy()
	
	renderer.SetDrawColor(255, 255, 255, 255)
	renderer.Clear()
	
	lines:= Lines(numL)
	Parallel(lines)
	running := true
	for running {
		event = sdl.WaitEventTimeout(1000) 
		switch t := event.(type) {
		case *sdl.MouseButtonEvent:
			if t.Button == sdl.BUTTON_LEFT {running = false}
		}
	}
}
