package main

import (
	"math"
	"github.com/veandco/go-sdl2/sdl"
)

var (
	winTitle string = "Graph"
	winWidth, winHeight float64 = 1200, 720
	window *sdl.Window
	renderer *sdl.Renderer
	R sdl.Rect
)

type (
	Line struct {
		A, B Point
		Dots []sdl.Point 
	}
	Point struct {
		X, Y float64
	}
)

func Calculations() (int, int, []sdl.Point) {
	var (
		y, minY, maxY int
		h2 int32 = int32(winHeight) / 2
		c Line
	)
	
	renderer.SetDrawColor(0, 0, 0, 255)
	for x := -20; x <= 20; x++ {
		y = x * x - 6 * x + 12
		if x == -20 {maxY ,minY = y, y}
		c.Dots = append(c.Dots, sdl.Point{int32(x), int32(y) - h2})
		if y > maxY {maxY = y}
		if y < minY {minY = y}
	}
	renderer.Present()
	return maxY, minY, c.Dots
}

func DrawChart(maxY, minY float64, Chart []sdl.Point) {
	var (
		R sdl.Rect
		Rx, Ry, Rw, Rh, Kx, Ky, Px, Py float64
		w2, h2 int32
		v Line
	)
	Rx, Ry, Rw, Rh = winWidth / 2 - 250, winHeight / 2 - maxY / 2, 500, maxY + math.Abs(minY)
	w2, h2 = int32(Rw) / 2, int32(Rh) / 2
	R = sdl.Rect{int32(Rx), int32(Ry), int32(Rw), int32(Rh)}
	renderer.DrawRect(&R)
	
	var A, B, C, D, a, b, c, d float64 = 0, winWidth, winHeight, 0, Rx, Rx + Rw, Ry + Rh, Ry
	Kx = (b - a) / (B - A)
	Px =  (a * B - b * A) / (B - A)
	Ky = (d - c) / (D - C)
	Py = (c * D - d * C) / (D - C)
	
	renderer.SetDrawColor(0, 0, 0, 255)
	
	for i := 1; i < len(Chart); i++ {
		v.A = Point{Kx * float64(Chart[i-1].X) + Px, Ky * float64(Chart[i-1].Y) + Py}
		v.B = Point{Kx * float64(Chart[i].X) + Px, Ky * float64(Chart[i].Y) + Py}
		renderer.DrawLine(int32(v.A.X) + w2, int32(v.A.Y) + h2, int32(v.B.X) + w2, int32(v.B.Y) + h2)
	}
	renderer.Present()
}

func main() {
	var event sdl.Event
	
	window, _ = sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int32(winWidth), int32(winHeight), sdl.WINDOW_SHOWN)
	renderer, _ = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	
	renderer.SetDrawColor(255, 255, 255, 255)
	renderer.Clear()
	
	maxY, minY, Chart := Calculations()
	DrawChart(float64(maxY), float64(minY), Chart)
	
	running := true
	for running {
		event = sdl.WaitEventTimeout(1000)
		switch t := event.(type) {
		case *sdl.MouseButtonEvent:
			if t.Button == sdl.BUTTON_LEFT {running = false}
		}
	}
	sdl.Quit()
}
