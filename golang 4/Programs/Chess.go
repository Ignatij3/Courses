package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

var (
	winTitle string = "Chess Board"
	winWidth, winHeight int32 = 1200, 720
	window *sdl.Window
	renderer *sdl.Renderer
	b sdl.Rect
)

func Chess() {
	var (
		ctrlw, ctrlh int32
		R sdl.Rect
		fill bool
		num int
	)
	
	renderer.SetDrawColor(194, 73, 19, 255)
	for {
		R.X, R.Y, R.W, R.H = (winWidth - 720)/2 + ctrlw, ctrlh, 90, 90
		renderer.DrawRect(&R)
		if fill {renderer.FillRect(&R); fill = false} else {fill = true}
		renderer.Present()
		if ctrlh == winHeight && num == 7 {break} else {
			if ctrlh == winHeight {ctrlh = 0; ctrlw += 90; num++} else {ctrlh +=90}
		}
	}
}

func main() {
	window, _ = sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, winWidth, winHeight, sdl.WINDOW_SHOWN)
	renderer, _ = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	defer window.Destroy()
	defer renderer.Destroy()
	
	b.X, b.Y, b.W, b.H  = 0, 0, winWidth, winHeight
	renderer.SetDrawColor(230, 220, 167, 255)
	renderer.FillRect(&b)
	renderer.Present()
	Chess()
	sdl.Delay(5000)
}
