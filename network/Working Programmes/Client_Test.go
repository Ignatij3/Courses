package main

import (
	"fmt"
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

var (
	renderer *sdl.Renderer
	window   *sdl.Window
)

func calculate(width, height int, angle, x, y float64) {
	var (
		radians float64 = (math.Pi / 180.0) * angle
		event   sdl.Event
	)

	renderer.SetDrawColor(255, 0, 0, 255)
	for {
		event = sdl.WaitEventTimeout(1)
		switch t := event.(type) {
		case *sdl.KeyboardEvent:
			if t.State == 1 {
				switch t.Keysym.Sym {
				case 27: //Escape
					break
				}
			}
		}

		x = x + math.Cos(radians)
		y = y + math.Sin(radians)
		fmt.Println(x, y)

		if y <= 0 {
			angle = 360 - angle
			radians = (math.Pi / 180.0) * angle
			fmt.Printf("New angle - %f\nNew radians - %f\n\n", angle, radians)
		} else if y >= float64(height) {
			angle = 360 - angle
			radians = (math.Pi / 180.0) * angle
			fmt.Printf("New angle - %f\nNew radians - %f\n\n", angle, radians)
		} else if x <= 0 {
			angle = 180 - angle
			radians = (math.Pi / 180.0) * angle
			fmt.Printf("New angle - %f\nNew radians - %f\n\n", angle, radians)
		} else if x >= float64(width) {
			angle = 180 - angle
			radians = (math.Pi / 180.0) * angle
			fmt.Printf("New angle - %f\nNew radians - %f\n\n", angle, radians)
		}
		renderer.DrawPoint(int32(x), int32(y))
		renderer.Present()
	}
}

func main() {
	window, _ = sdl.CreateWindow("Bouncing dot", sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, 1200, 720, sdl.WINDOW_SHOWN)
	renderer, _ = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)

	renderer.SetDrawColor(255, 255, 255, 255)
	renderer.Clear()
	renderer.Present()
	calculate(1200, 720, 120.0, 200.0, 200.0)

	window.Destroy()
	renderer.Destroy()
	sdl.Quit()
}
