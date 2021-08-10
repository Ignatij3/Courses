package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type (
	Slice struct {
		s3d []Segment3D
		s2d []Segment2D
	}
	Segment3D struct {
		X, Y, Z float64
	}
	Segment2D struct {
		X, Y float64
	}
	Line struct {
		A, B Segment2D
	}
)
var (
	winTitle string = "Railroad"
	winWidth, winHeight float64 = 1200, 720
	window *sdl.Window
	renderer *sdl.Renderer
)

func SetPoints() Slice {
	var (
		p3d Slice
		p, k float64
	)
	
	/*for ; 1.2 + k * 0.55 < 70; k++ { Я тут попытался создать широкие шпалы
		for i := 0.01; i + 1.2 + k * 0.55 < 1.2 + k * 0.55 + 2; i += 0.01 {
			p3d.s3d = append(p3d.s3d, Segment3D{-0.95, -1.8, i + 1.2 + k * 0.55})
			p3d.s3d = append(p3d.s3d, Segment3D{0.95, -1.8, i + 1.2 + k * 0.55})
		}
	}*/
	for i := 0.6; i <= 0.75; i += 0.01 {
		p3d.s3d = append(p3d.s3d, Segment3D{i, -1.65, 1}) //Дело в том, что подряд шли 2-е начальные точки, которые рисовались как отрезок
		p3d.s3d = append(p3d.s3d, Segment3D{i, -1.65, 70})
		p3d.s3d = append(p3d.s3d, Segment3D{-i, -1.65, 1})
		p3d.s3d = append(p3d.s3d, Segment3D{-i, -1.65, 70})
	}
	for ; 1.2 + k * 0.55 < 70; k++ {
		p3d.s3d = append(p3d.s3d, Segment3D{-0.95, -1.8, 1.2 + k * 0.55})
		p3d.s3d = append(p3d.s3d, Segment3D{0.95, -1.8, 1.2 + k * 0.55})
	}
	for ; 3 + 6 * p < 70; p++ {
		p3d.s3d = append(p3d.s3d, Segment3D{1.95, -2, 3 + 6 * p})
		p3d.s3d = append(p3d.s3d, Segment3D{1.95, 2.2, 3 + 6 * p})
		p3d.s3d = append(p3d.s3d, Segment3D{-1.95, -2, 3 + 6 * p})
		p3d.s3d = append(p3d.s3d, Segment3D{-1.95, 2.2, 3 + 6 * p})
		p3d.s3d = append(p3d.s3d, Segment3D{-1.95, 2.2, 3 + 6 * p})
		p3d.s3d = append(p3d.s3d, Segment3D{-1.3, 2.0, 3 + 6 * p})
	}
	for m := 2.2; m > 2.2 - (2.2 / 4); m -= 0.15 {
		p3d.s3d = append(p3d.s3d, Segment3D{1.95, m, 3})
		p3d.s3d = append(p3d.s3d, Segment3D{1.95, m, 70})
	}
	return p3d
}

func DrawLines(s Slice) {
	var Mx, My, minX, minY, maxX, maxY, Kx, Ky, Px, Py float64
	var t Line
	
	for i := range(s.s3d) {
		Mx = -s.s3d[i].X / s.s3d[i].Z
		My = -s.s3d[i].Y / s.s3d[i].Z
		
		if i == 0 {minX, maxX, maxY, minY = Mx, Mx, My, My}
		if Mx < minX {minX = Mx}
		if Mx > maxX {maxX = Mx}
		if My < minY {minY = My}
		if My > maxY {maxY = My}
		s.s2d = append(s.s2d, Segment2D{Mx, My})
	}
	
	var A, B, C, D, a, b, c, d float64 = minX, maxX, maxY, minY, winWidth, 0, winHeight, 0
	Kx = (b - a) / (B - A)
	Px =  (a * B - b * A) / (B - A)
	Ky = (d - c) / (D - C)
	Py = (c * D - d * C) / (D - C)
	
	renderer.SetDrawColor(0, 0, 0, 255)
	for i := 1; i < len(s.s2d); i+=2 {
		t.A = Segment2D{Kx * s.s2d[i-1].X + Px, Ky * s.s2d[i-1].Y + Py}
		t.B = Segment2D{Kx * s.s2d[i].X + Px, Ky * s.s2d[i].Y + Py}
		renderer.DrawLine(int32(t.A.X), int32(t.A.Y), int32(t.B.X), int32(t.B.Y))
	}
	renderer.Present()
}

func main() {
	var event sdl.Event
	
	window, _ = sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int32(winWidth), int32(winHeight), sdl.WINDOW_SHOWN)
	renderer, _ = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	renderer.SetDrawColor(255, 255, 255, 255)
	renderer.Clear()
	
	slice := SetPoints()
	DrawLines(slice)
	
	running := true
	for running {
		event = sdl.WaitEventTimeout(1000)
		switch t := event.(type) {
		case *sdl.MouseButtonEvent:
			if t.Button == sdl.BUTTON_LEFT {running = false}
		}
	}
	sdl.Quit()
}
