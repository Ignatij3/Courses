package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

var winTitle string = "package sdl. Test Window."
var winWidth, winHeight int32 = 400, 240

func CorrectClear(renderer *sdl.Renderer) error {
	r, g, b, a, _ := renderer.GetDrawColor()
	renderer.SetDrawColor(255, 0, 220, 255)
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
	window, err = sdl.CreateWindow(winTitle, sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_UNDEFINED, winWidth, winHeight, sdl.WINDOW_SHOWN)
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
	sdl.Delay(3456)
}
