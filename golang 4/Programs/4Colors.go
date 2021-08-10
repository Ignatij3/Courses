package main

import (
	"time"
	"math/rand"
	"github.com/veandco/go-sdl2/sdl"
)

var (
	window *sdl.Window
	renderer *sdl.Renderer
	winTitle string = "4 Colors"
	winWidth, winHeight int32 = 1200, 720
)

type (
	Segment struct {
		A, B sdl.Point
	}
)

func Dots() {
	var (
		s Segment
		sq sdl.Rect
		side string
	)
	mLines := make([]Segment, 2)
	renderer.SetDrawColor(0, 0, 0, 255)
	
	for i := 0; i < 2; i++ {
		switch rand.Intn(4) {
			case 0 :
				s.A = sdl.Point{0, int32(rand.Intn(int(winHeight)))}
				side = "LEFT"
			case 1 :
				s.A = sdl.Point{int32(rand.Intn(int(winWidth))), 0}
				side = "UP"
			case 2 :
				s.A = sdl.Point{winWidth, int32(rand.Intn(int(winHeight)))}
				side = "RIGHT"
			case 3 :
				s.A = sdl.Point{int32(rand.Intn(int(winWidth))), winHeight}
				side = "DOWN"
		}
		
		s.B = Lines(side)
		mLines[i] = s
		renderer.DrawLine(s.A.X, s.A.Y, s.B.X, s.B.Y)
		sdl.Delay(700)
		renderer.Present()
	}
	
	for {
		p1 := sdl.Point{0, 0}
		p2 := sdl.Point{int32(rand.Intn(int(winWidth))), int32(rand.Intn(int(winHeight)))}
		b1, b2:= Oneside(mLines[0], p1, p2), Oneside(mLines[1], p1, p2)
		
		switch {
		  case b1 && b2:
			renderer.SetDrawColor(255, 0, 0, 255)
		  case b1 && !b2:
			renderer.SetDrawColor(215, 120, 0, 255)
		  case !b1 && b2:
			renderer.SetDrawColor(0, 255, 0, 255)
		  case !b1 && !b2:
			renderer.SetDrawColor(0, 0, 255, 255)
		}
		
		sq.X, sq.Y, sq.W, sq.H  = p2.X, p2.Y, 1, 1
		renderer.FillRect(&sq)
		renderer.Present()
	}
}

func Lines(side string) sdl.Point {
	var (
		cSide string
		s Segment
	)
	start:
	switch rand.Intn(4) {
		case 0 :
			s.B = sdl.Point{winWidth, int32(rand.Intn(int(winHeight)))}
			cSide = "RIGHT"
		case 1 :
			s.B = sdl.Point{0, int32(rand.Intn(int(winHeight)))}
			cSide = "LEFT"
		case 2 :
			s.B = sdl.Point{int32(rand.Intn(int(winWidth))), 0}
			cSide = "UP"
		case 3 :
			s.B = sdl.Point{int32(rand.Intn(int(winWidth))), winHeight}
			cSide = "DOWN"
	}
	if side == cSide {goto start}
	return s.B
}

func Oneside(s Segment, A sdl.Point, B sdl.Point) bool {
	var (
		Vx, Vy, Ux, Uy int32
		t, T float64
	)
	Vx = B.X - A.X
	Vy = B.Y - A.Y
	Ux = s.B.X - s.A.X
	Uy = s.B.Y - s.A.Y
	if Vx * Uy == Vy * Ux {return false} else {
		t = float64(((s.A.X - A.X) * Uy - (s.A.Y - A.Y) * Ux)) / float64((Vx * Uy - Vy * Ux))
		T = float64(((s.A.X - A.X) * Vy - (s.A.Y - A.Y) * Vx)) / float64((Vx * Uy - Vy * Ux))
		if t < 0 || t > 1 || T < 0 || T > 1 {return false} else {return true}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	window, _ = sdl.CreateWindow(winTitle, sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, winWidth, winHeight, sdl.WINDOW_SHOWN)
	renderer, _ = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	defer window.Destroy()
	defer renderer.Destroy()
	
	renderer.SetDrawColor(255, 255, 255, 255)
	renderer.Clear()
	
	Dots()
}
