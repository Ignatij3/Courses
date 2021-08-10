package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

var winTitle string = "sdl. Simple drawing."
var winWidth, winHeight int32 = 320, 200

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
	var p sdl.Rect
	var i int32
	for i = 10; i<150; i+=15  {
		renderer.DrawLine(i, i, i+25, i)
		renderer.DrawPoint(i+33, i)
		p = sdl.Rect{i+40, i, 15, 10}
		renderer.DrawRect(&p)
		p = sdl.Rect{i+60, i, 8, 10}
		renderer.FillRect(&p)
		renderer.Present()
		sdl.Delay(100)
	}	
	sdl.Delay(3000)
}
