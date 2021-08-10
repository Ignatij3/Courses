package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"os"
	"math/rand"
	"math"
	"time"
)

var winTitle string = "FAN"
var winWidth, winHeight int32 = 800, 600

const spokes = 25

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
	
	rand.Seed(time.Now().UnixNano())
	
	var  p, A, B, runner sdl.Point
	
	p.X = winWidth/2	
	p.Y = 0	
	
	for {
		renderer.SetDrawColor(255, 255, 255, 255) 
		renderer.Clear()
		renderer.Present()
		
		A.X = rand.Int31n(winWidth/2)
		A.Y = rand.Int31n(winHeight)
		B.Y = rand.Int31n(winHeight)
		B.X = rand.Int31n(winWidth/2) + (winWidth/2)
		deltaX:= float64(B.X-A.X)/float64(spokes-1)
		deltaY:= float64(B.Y-A.Y)/float64(spokes-1)
		
		renderer.SetDrawColor(0, 0, 0, 255)
		for i:= 0; i<spokes; i++  {
			n:= float64(i)
			runner.X = A.X + int32(math.Round(n*deltaX))
			runner.Y = A.Y + int32(math.Round(n*deltaY))
			renderer.DrawLine(p.X, p.Y, runner.X, runner.Y)
			renderer.Present()
			sdl.Delay(20)
		}	
		event := sdl.WaitEventTimeout(2000) 
		if _, ok := event.(*sdl.KeyboardEvent); ok {
			break
		}
	}
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
