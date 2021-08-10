package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"os"
	"math/rand"
	"math"
	"time"
)

var winTitle string = "WRAP"
var winWidth, winHeight int32 = 800, 600
const alpha = 0.1

func run() int {
	var window *sdl.Window
	var renderer *sdl.Renderer
	window, err := sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		winWidth, winHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
		return 1
	}
	defer window.Destroy()
	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
		return 2
	}
	defer renderer.Destroy()
	
	renderer.SetDrawColor(255, 255, 255, 255) 
	renderer.Clear()
	renderer.Present()

	rand.Seed(time.Now().UnixNano())
	p := make ([]sdl.Point, 5)
	p[0].X = rand.Int31n(winWidth/3)
	p[0].Y = rand.Int31n(winHeight/3)
	p[1].X = winWidth - p[0].X
	p[1].Y = p[0].Y
	p[2].X = p[1].X
	p[2].Y = winHeight- p[1].Y
	p[3].X = p[0].X
	p[3].Y = p[2].Y
	p[4] = p[0]

	renderer.SetDrawColor(0, 0, 0, 255)
	for {
		renderer.DrawLines(p)
		renderer.Present()
		sdl.Delay(40)
		for i:= 0; i < 4; i++  {
			p[i].X = int32(math.Round(float64(p[i].X)*(1-alpha) + float64(p[i+1].X)*alpha))
			p[i].Y = int32(math.Round(float64(p[i].Y)*(1-alpha) + float64(p[i+1].Y)*alpha))	
		}
		if p[0].X == p[4].X && p[0].Y == p[4].Y  { break }
		p[4] = p[0]
	}	
	sdl.Delay(3000)
	return 0
}

func main() {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		fmt.Printf("Failed to initialize SDL: %s\n", err)
		return
	}
	run()
	sdl.Quit()
}
