package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

var winTitle string = "sdl. Simple drawing."
var winWidth, winHeight int32 = 320, 240

func CorrectClear(renderer *sdl.Renderer) error {
	// Clear renderer in white
	r, g, b, a, _ := renderer.GetDrawColor()
	renderer.SetDrawColor(255, 255, 255, 255)
	err := renderer.Clear()
	renderer.SetDrawColor(r, g, b, a)
	renderer.Present()
	return err
}

func main() {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		fmt.Printf("Failed to initialize SDL: %s\n", err)
		return
	}
	defer sdl.Quit()

	var window *sdl.Window
	window, err = sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, winWidth, winHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Printf("Failed to create window: %s\n", err)
		return
	}
	defer window.Destroy()

	var renderer *sdl.Renderer
	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Printf("Failed to create renderer: %s\n", err)
		return
	}
	defer renderer.Destroy()

	CorrectClear(renderer)
	renderer.SetDrawColor(0, 0, 0, 255)  // black on white

	var points []sdl.Point
	var rects []sdl.Rect
	var i int32
	for i = 0; i<100; i+=12  {
		points = append (points, sdl.Point{3*i+15, i*i/42})
		rects = append (rects, sdl.Rect{3*i+15-3, i*i/42-3, 7, 7})
	}	
	points = append(points, points[0])
	renderer.DrawLines(points)
	points = nil 		// empty slice
	renderer.FillRects(rects)
	rects = nil 		// empty slice
	for i = 0; i<100; i+=12  {
		points = append (points, sdl.Point{3*i+5, i*i/42+10})
		rects = append (rects, sdl.Rect{3*i+5-3, i*i/42-3+10, 7, 7})
	}
	renderer.DrawPoints(points)
	renderer.DrawRects(rects)
	renderer.Present()
	sdl.Delay(3000)
}
