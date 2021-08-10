package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"math"
)

var winTitle string = "Circle"
var winWidth, winHeight int32 = 800, 600

func circle(renderer *sdl.Renderer, c sdl.Point, r int32) {
	x, y := float64(r), 0.0
	var x1, y1 float64
	fi:= 1.0/float64(r)
	cos, sin := math.Cos(fi), math.Sin(fi)
	
	for i:= 0; float64(i)<2.0*math.Pi*float64(r); i++  {
		x1, y1 = x*cos - y*sin, x*sin + y*cos
		renderer.DrawLine(c.X+int32(math.Round(x)), c.Y+int32(math.Round(y)), 
									 c.X+int32(math.Round(x1)), c.Y+int32(math.Round(y1)))
		x, y = x1, y1 
	}	
	return 
}

func main() {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		fmt.Printf("Failed to initialize SDL: %s\n", err)
		return
	}
	var window *sdl.Window
	var renderer *sdl.Renderer
	window, err = sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		winWidth, winHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Printf("Failed to create window: %s\n", err)
		return 
	}
	defer window.Destroy()
	
	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Printf("Failed to create renderer: %s\n", err)
		return
	}
	defer renderer.Destroy()
		renderer.SetDrawColor(255, 255, 255, 255) 
	renderer.Clear()
	renderer.Present()

	renderer.SetDrawColor(0, 0, 0, 255)
	for r:= 20; r<=200; r += 20 {
		circle(renderer, sdl.Point{winWidth/2, winHeight/2}, int32(r))
	}
	renderer.Present()
	sdl.Delay(5000)
	sdl.Quit()
}
