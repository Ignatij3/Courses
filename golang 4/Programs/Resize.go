package main

import (
	"time"
	"math/rand"
	"github.com/veandco/go-sdl2/sdl"
)

var (
	winTitle string = "Little Projection"
	winWidth, winHeight float64 = 1200, 720
	window *sdl.Window
	renderer *sdl.Renderer
	Ra, Rb sdl.Rect
)

type (
	Line struct {
		A, B [2]Point
	}
	Point struct {
		X, Y float64
	}
)

func DrawLines() {
	var k Line
	
	for {
		event := sdl.PollEvent() 
		if _, ok := event.(*sdl.MouseButtonEvent); ok {
			break
		}
		renderer.SetDrawColor(uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), 255)
		k.A[0] = Point{float64(rand.Int31n(int32(Ra.X) + 260)) + 100, float64(rand.Int31n(int32(Ra.Y) + 460)) + 100}
		k.B[0] = Point{float64(rand.Int31n(int32(Ra.X) + 260)) + 100, float64(rand.Int31n(int32(Ra.Y) + 460)) + 100}
		renderer.DrawLine(int32(k.A[0].X), int32(k.A[0].Y), int32(k.B[0].X), int32(k.B[0].Y))
		SecondLine(k.A, k.B)
		renderer.Present()
		sdl.Delay(100)
	}
}

func DrawRectangles() {
	renderer.SetDrawColor(0, 0, 0, 255)
	Ra.X, Ra.Y, Ra.W, Ra.H = 100, 100, 360, 560
	Rb.X, Rb.Y, Rb.W, Rb.H = int32(winWidth) - 380, int32(winHeight) - 280, 280, 180
	
	renderer.DrawRect(&Ra)
	renderer.DrawRect(&Rb)
	renderer.Present()
}

func SecondLine(Ap, Bp [2]Point) {
	var (
		k Line
		Kx, Ky, Px, Py float64
	)
	
	var A, B, C, D, a, b, c, d float64 = 100, 460, 660, 100, winWidth - 380, winWidth - 100, winHeight - 100, winHeight - 280
	Kx = (b - a) / (B - A)
	Px =  (a * B - b * A) / (B - A)
	Ky = (d - c) / (D - C)
	Py = (c * D - d * C) / (D - C)
	
	k.A[1] = Point{Kx * Ap[0].X + Px, Ky * Ap[0].Y + Py}
	k.B[1] = Point{Kx * Bp[0].X + Px, Ky * Bp[0].Y + Py}
	
	renderer.DrawLine(int32(k.A[1].X), int32(k.A[1].Y), int32(k.B[1].X), int32(k.B[1].Y))
}

func main() {
	window, _ = sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int32(winWidth), int32(winHeight), sdl.WINDOW_SHOWN)
	renderer, _ = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	
	renderer.SetDrawColor(255, 255, 255, 255)
	renderer.Clear()
	
	DrawRectangles()
	DrawLines()
	time.Sleep(5 * time.Second)
	sdl.Quit()
}
