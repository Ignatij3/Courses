package main

import "github.com/veandco/go-sdl2/sdl"

var (
	winTitle string = "Rotated Chess Board"
	winWidth, winHeight int32 = 800, 800
	window *sdl.Window
	renderer *sdl.Renderer
)

func Carcass() {
	var ctrlw, ctrlh, mirrw, mirrh, num int32
	
	renderer.SetDrawColor(194, 73, 19, 255)
	ctrlw, mirrw = winWidth / 2, winWidth
	ctrlh, mirrh = winHeight / 2, winHeight / 2
	num = winHeight / 16
	ctrlh = 0
	
	for ctrlw >= 0 {
		renderer.DrawLine(ctrlw, ctrlh, mirrw, mirrh)
		sdl.Delay(60)
		renderer.Present()
		
		ctrlw -= num
		ctrlh += num
		mirrw -= num
		mirrh += num
	}
	
	ctrlw = winWidth / 2
	ctrlh = 0
	mirrw = 0
	mirrh = winHeight / 2
	
	for ctrlw <= winWidth {
		renderer.DrawLine(ctrlw, ctrlh, mirrw, mirrh)
		sdl.Delay(60)
		renderer.Present()
		
		ctrlw += num
		ctrlh += num
		mirrw += num
		mirrh += num
	}
	FillSquares()
}

func FillSquares() {
	var (
		w, h, w2, h2, wAdd, hAdd, wAdd2, hAdd2, check, num int32
		shift bool
	)
	w, check = winWidth / 2, winWidth / 2
	w2, wAdd2 = w - 50, w - 50
	h2, hAdd2 = h + 50, h + 50
	wAdd = w
	
	for {
		renderer.DrawLine(wAdd, hAdd, wAdd2, hAdd2)
		renderer.Present()
		sdl.Delay(10)
		wAdd += 1
		hAdd += 1
		wAdd2 += 1
		hAdd2 += 1
		
		if hAdd2 == h + 100 {
			w += 100
			h += 100
			w2 += 100
			h2 += 100
			
			wAdd = w
			hAdd = h
			wAdd2 = w2
			hAdd2 = h2
		}
		
		if shift && w == check + 400 {
			shift = false
			num += 100
			w = winWidth / 2 - num
			h = num
			check = w
			w2 = w - 50
			h2 = h + 50
			wAdd = w
			hAdd = h
			wAdd2 = w2
			hAdd2 = h2
		} else if !shift && w == check + 400 {
			shift = true
			w = winWidth / 2 - num
			h = 100 + num
			check = w
			w2 = w - 50
			h2 = h + 50
			wAdd = w
			hAdd = h
			wAdd2 = w2
			hAdd2 = h2
		}
		if h == winHeight / 2 && w == 0 {break}
	}
}

func main() {
	var event sdl.Event
	
	window, _ = sdl.CreateWindow(winTitle, sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, winWidth, winHeight, sdl.WINDOW_SHOWN)
	renderer, _ = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	defer window.Destroy()
	defer renderer.Destroy()
	
	renderer.SetDrawColor(240, 230, 200, 255)
	renderer.Clear()
	
	Carcass()
	running := true
	for running {
		event = sdl.WaitEventTimeout(1000) 
		switch t := event.(type) {
		case *sdl.MouseButtonEvent:
			if t.Button == sdl.BUTTON_LEFT {running = false}
		}
	}
}
