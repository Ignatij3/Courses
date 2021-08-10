package main

import (
	"fmt"
	"time"
	"math"
	"math/rand"
	"github.com/veandco/go-sdl2/sdl"
)

var (
	winTitle string = "Flags"
	winWidth, winHeight float64 = 1200, 720
	window *sdl.Window
	renderer *sdl.Renderer
)
const (
	alfa = 0.7
 ) 
type (
	Line struct {
		A, B Point
	}
	Triangle struct {
		A, B, C Point
	}
	Point struct {
		X, Y float64
	}
)

func Lines(num int) {
	var (
		l Line
	)
	
	l.A = Point{winWidth / 2, winHeight / 2}
	l.B = Point{winWidth / 2, winHeight / 10}

	renderer.DrawLine(int32(l.A.X), int32(l.A.Y), int32(l.B.X), int32(l.B.Y))
	renderer.Present()
	
	sin, cos := math.Sin(2 * math.Pi / float64(num)), math.Cos(2 * math.Pi / float64(num))
	for i := 0; i < num; i++ {
		renderer.SetDrawColor(0, 0, 0, 255)
		
		l.B.X, l.B.Y = (l.B.X - l.A.X) * cos - (l.B.Y - l.A.Y ) * sin + l.A.X,  (l.B.Y - l.A.Y) * cos + (l.B.X - l.A.X) * sin + l.A.Y
		
		renderer.DrawLine(int32(l.A.X), int32(l.A.Y), int32(l.B.X), int32(l.B.Y))

		Triangles(l.A, l.B)
		renderer.Present()
	}
}

func Triangles(A, B Point) {
	var t Triangle
	
	t.A = Point{B.X, B.Y}
	t.C = Point{A.X*(1-alfa) + B.X *alfa, A.Y*(1-alfa) + B.Y *alfa}
	c, s:= 0.5, math.Sqrt(3.0)/2.0
	t.B.X, t.B.Y = (t.A.X - t.C.X) * c - (t.A.Y - t.C.Y ) * s + t.C.X,  (t.A.Y - t.C.Y) * c + (t.A.X - t.C.X) * s + t.C.Y
	
	renderer.DrawLine(int32(t.A.X), int32(t.A.Y), int32(t.B.X), int32(t.B.Y))
	renderer.DrawLine(int32(t.B.X), int32(t.B.Y), int32(t.C.X), int32(t.C.Y))
}

func main() {

	var amount int
	fmt.Print("Enter amount of flags:")
	fmt.Scan(&amount)
	for amount <= 0 {
		fmt.Print("Error, try again:")
		fmt.Scan(&amount)
	}
	
	rand.Seed(time.Now().UnixNano())
	window, _ = sdl.CreateWindow(winTitle, sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, int32(winWidth), int32(winHeight), sdl.WINDOW_SHOWN)
	renderer, _ = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	defer window.Destroy()
	defer renderer.Destroy()
	
	renderer.SetDrawColor(255, 255, 255, 255)
	renderer.Clear()
	
	renderer.SetDrawColor(0, 0, 0, 255)
	Lines(amount)
	sdl.Delay(60000)
}
