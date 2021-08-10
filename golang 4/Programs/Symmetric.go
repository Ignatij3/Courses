package main

import (
	"time"
	"math/rand"
	"github.com/veandco/go-sdl2/sdl"
)

var (
	winTitle string = "Triangles"
	winWidth, winHeight float64 = 1200, 720
	window *sdl.Window
	renderer *sdl.Renderer
)

type (
	Line struct {
		A, B Point
	}
	Triangle struct {
		A, B, C, H [2]Point
	}
	Point struct {
		X, Y float64
	}
)

func Lines() {
	var (
		s Line
		side string
	)
	renderer.SetDrawColor(0, 0, 0, 255)
	
	switch rand.Intn(4) {
		case 0 :
			s.A = Point{0, float64(rand.Intn(int(winHeight)))}
			side = "LEFT"
		case 1 :
			s.A = Point{float64(rand.Intn(int(winWidth))), 0}
			side = "UP"
		case 2 :
			s.A = Point{winWidth, float64(rand.Intn(int(winHeight)))}
			side = "RIGHT"
		case 3 :
			s.A = Point{float64(rand.Intn(int(winWidth))), winHeight}
			side = "DOWN"
	}
	s.B = SecondLine(side)
	renderer.DrawLine(int32(s.A.X), int32(s.A.Y), int32(s.B.X), int32(s.B.Y))
	renderer.Present()
	
	Cross(s.A, s.B)
}

func SecondLine(side string) Point {
	var (
		cSide string
		s Line
	)
	start:
	switch rand.Intn(4) {
		case 0 :
			s.B = Point{winWidth, float64(rand.Intn(int(winHeight)))}
			cSide = "RIGHT"
		case 1 :
			s.B = Point{0, float64(rand.Intn(int(winHeight)))}
			cSide = "LEFT"
		case 2 :
			s.B = Point{float64(rand.Intn(int(winWidth))), 0}
			cSide = "UP"
		case 3 :
			s.B = Point{float64(rand.Intn(int(winWidth))), winHeight}
			cSide = "DOWN"
	}
	if side == cSide {goto start}
	return s.B
}

func Cross(A, B Point) {
	var r Triangle
	
	Vx := B.X - A.X
	Vy := B.Y - A.Y
	
	for {
		event := sdl.PollEvent() 
		if _, ok := event.(*sdl.MouseButtonEvent); ok {
			break
		}
		r.A[0] = Point{float64(rand.Intn(int(winWidth))), float64(rand.Intn(int(winHeight)))}
		r.B[0] = Point{float64(rand.Intn(int(winWidth))), float64(rand.Intn(int(winHeight)))}
		r.C[0] = Point{float64(rand.Intn(int(winWidth))), float64(rand.Intn(int(winHeight)))}
		
		r.H[1] = Point{A.X + Vx * 0.5, A.Y + Vy * 0.5}
		r.A[1] = Point{r.H[1].X * 2 - r.A[0].X, r.H[1].Y * 2 - r.A[0].Y}
		r.B[1] = Point{r.H[1].X * 2 - r.B[0].X, r.H[1].Y * 2 - r.B[0].Y}
		r.C[1] = Point{r.H[1].X * 2 - r.C[0].X, r.H[1].Y * 2 - r.C[0].Y}
		
		if OverBorder(r.A[1], r.B[1], r.C[1]) {
			renderer.SetDrawColor(255, 0, 0, 255)
		} else {
			renderer.SetDrawColor(0, 0, 255, 255)
		}
		
		renderer.DrawLine(int32(r.A[0].X), int32(r.A[0].Y), int32(r.B[0].X), int32(r.B[0].Y))
		renderer.DrawLine(int32(r.B[0].X), int32(r.B[0].Y), int32(r.C[0].X), int32(r.C[0].Y))
		renderer.DrawLine(int32(r.C[0].X), int32(r.C[0].Y), int32(r.A[0].X), int32(r.A[0].Y))
		renderer.DrawLine(int32(r.A[1].X), int32(r.A[1].Y), int32(r.B[1].X), int32(r.B[1].Y))
		renderer.DrawLine(int32(r.B[1].X), int32(r.B[1].Y), int32(r.C[1].X), int32(r.C[1].Y))
		renderer.DrawLine(int32(r.C[1].X), int32(r.C[1].Y), int32(r.A[1].X), int32(r.A[1].Y))
		
		renderer.Present()
		sdl.Delay(1500)
		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.Clear()
		renderer.SetDrawColor(0, 0, 0, 255)
		renderer.DrawLine(int32(A.X), int32(A.Y), int32(B.X), int32(B.Y))
	}
}

func OverBorder(A, B, C Point) bool {
	if (A.X > winWidth || A.Y > winHeight) ||
		(B.X > winWidth || B.Y > winHeight) ||
		(C.X > winWidth || C.Y > winHeight) ||
		(A.X < 0 || A.Y < 0) ||
		(B.X < 0 || B.Y < 0) ||
		(C.X < 0 || C.Y < 0) {return true} else {return false}
}

func main() {
	
	rand.Seed(time.Now().UnixNano())
	window, _ = sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int32(winWidth), int32(winHeight), sdl.WINDOW_SHOWN)
	renderer, _ = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	defer window.Destroy()
	defer renderer.Destroy()
	
	renderer.SetDrawColor(255, 255, 255, 255)
	renderer.Clear()
	
	Lines()	
	sdl.Delay(60000)
}
