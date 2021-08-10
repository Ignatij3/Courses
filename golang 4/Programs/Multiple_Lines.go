package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

var winTitle string = "Multiple Lines"
var winWidth, winHeight int32 = 1920, 1080

func main() {
	var (
		window *sdl.Window
		renderer *sdl.Renderer
		event sdl.Event
		a, b sdl.Point
		ctrl, ctrl2 int32
		count int
	)
	window, _ = sdl.CreateWindow(winTitle, sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, winWidth, winHeight, sdl.WINDOW_SHOWN)
	renderer, _ = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	defer window.Destroy()
	defer renderer.Destroy()
	
	renderer.SetDrawColor(255, 100, 30, 255)
	for {
		a.X, a.Y = ctrl, 0
		b.X, b.Y = ctrl2, winHeight
		renderer.DrawLine(a.X, a.Y, b.X, b.Y)
		count++
		
		if ctrl == winWidth {
			ctrl = 0
			ctrl2 += winWidth/3
		} else {
			ctrl += winWidth/5
		}
		if ctrl2 > winWidth {break}
		
		renderer.Present()
		sdl.Delay(500)
	}
	
	running := true
	for running {
		event = sdl.WaitEventTimeout(1000) 
		switch t := event.(type) {
		case *sdl.MouseButtonEvent:
			if t.Button == sdl.BUTTON_LEFT {running = false; fmt.Println("Было нарисовано столько линий:", count)}
		}
	}
}
