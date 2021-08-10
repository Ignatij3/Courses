package main

import "github.com/veandco/go-sdl2/sdl"

var (
	winTitle string = "RGB"
	winWidth, winHeight int = 1920, 1080
	window *sdl.Window
	renderer *sdl.Renderer
)

func FillScreen() {
	var color, change int = 0, 1
	
	for y := 0; y < winHeight; y++ {
		for x := 0; x < winWidth; x++ {
			switch color {
				case 0 :
					renderer.SetDrawColor(255, 0, 0, 255)
					color = 1
				case 1 :
					renderer.SetDrawColor(0, 255, 0, 255)
					color = 2
				case 2 :
					renderer.SetDrawColor(0, 0, 255, 255)
					color = 0
			}
			renderer.DrawPoint(int32(x), int32(y))
		}
		color = change
		if change == 2 {change = 0} else {change++}
	}
}

func main() {
	window, _ = sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int32(winWidth), int32(winHeight), sdl.WINDOW_SHOWN)
	renderer, _ = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	defer window.Destroy()
	defer renderer.Destroy()
	
	renderer.SetDrawColor(255, 255, 255, 255)
	renderer.Clear()
	
	FillScreen()
	renderer.Present()
	sdl.Delay(6000)
}
