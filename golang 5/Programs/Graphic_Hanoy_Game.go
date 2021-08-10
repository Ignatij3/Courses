package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"rational/change"
)

type (
	Ring struct {
		n, width int32
	}
	Triangle struct {
		A, B, C Point
	}
	Point struct {
		X, Y int32
	}
)

var (
	winTitle string = "Hanoy"
	//winWidth, winHeight int32 = 1200, 720 //5:3
	winWidth, winHeight int32 = 1920, 1080 //16:9
	//winWidth, winHeight int32 = 1920, 1200 //16:10
	window *sdl.Window
	renderer *sdl.Renderer
)

func CreateRings(first *[]Ring) {
	for n := int32(8); n >= 0; n-- {
		*first = append(*first, Ring{n, ((winWidth * 3) / 32) + n * winWidth / 48})
	}
}

func FillTriangle(A, B, C Point) {
	D := Point{A.X, A.Y}
	E := Point{C.X, A.Y}
	for E.X != B.X {
		D.X++
		D.Y++
		E.X--
		E.Y++
		renderer.DrawLine(D.X, D.Y, E.X, E.Y)
	}
}

func DrawArrow(pos int32, r, g, b uint8) {
	var (
		arrow sdl.Rect
		t Triangle
		frac *change.Fraction
	)
	frac = &change.Fraction{int(winWidth), int(winHeight)}
	frac.Reduce()
	
	if r == 255 {renderer.SetDrawColor(r, g, b, 255)} else {renderer.SetDrawColor(0, 0, 0, 255)}
	t.A = Point{(((winWidth * 7) / 24) * pos) + (winWidth * 67) / 384, (winHeight * 5) / 72}
	t.B = Point{(((winWidth * 7) / 24) * pos) + (winWidth * 5) / 24, (winHeight * 7) / 54}
	t.C = Point{(((winWidth * 7) / 24) * pos) + (winWidth * 31) / 128, (winHeight * 5) / 72}
	renderer.DrawLine(t.C.X, t.C.Y, t.A.X, t.A.Y)
	
	renderer.SetDrawColor(r, g, b, 255)
	if (*frac).A == 16 && (*frac).B == 9 {arrow = sdl.Rect{(((winWidth * 7) / 24) * pos) + (winWidth * 77) / 384, 1, winWidth / 64, (winHeight * 5) / 54}}
	renderer.FillRect(&arrow)
	
	if r == 255 {renderer.SetDrawColor(r, g, b, 255)} else {renderer.SetDrawColor(0, 0, 0, 255)}
	renderer.DrawRect(&arrow)
	renderer.SetDrawColor(r, g, b, 255)
	if (*frac).A == 16 && (*frac).B == 9 {FillTriangle(t.A, t.B, t.C)}
	
	if r == 255 {renderer.SetDrawColor(r, g, b, 255)} else {renderer.SetDrawColor(0, 0, 0, 255)}
	renderer.DrawLine(t.A.X, t.A.Y, t.B.X, t.B.Y)
	renderer.DrawLine(t.B.X, t.B.Y, t.C.X, t.C.Y)
	renderer.Present()
}

func DrawObj() {
	var base sdl.Rect
	renderer.SetDrawColor(255, 255, 255, 255)
	renderer.Clear()
	
	renderer.SetDrawColor(0, 0, 0, 255)
	base = sdl.Rect{(winWidth * 149) / 1920, (winHeight * 799) / 1080, (winWidth * 811) / 960, (winHeight * 17) / 180}
	renderer.FillRect(&base)
	base = sdl.Rect{(winWidth * 49) / 1920, (winHeight * 899) / 1080, (winWidth * 911) / 960, (winHeight * 17) / 180}
	renderer.FillRect(&base)
	
	for k := int32(0); k < 3; k++ {
		base = sdl.Rect{(((winWidth * 7) / 24) * k) + (winWidth * 349) / 1920, (winHeight * 53) / 360, (winWidth * 17) / 320, (winHeight * 401) / 540}
		renderer.DrawRect(&base)
	}
	
	renderer.SetDrawColor(170, 255, 255, 255)
	for k := int32(0); k < 3; k++ {
		base = sdl.Rect{(((winWidth * 7) / 24) * k) + (winWidth * 35) / 192, (winHeight * 4) / 27, (winWidth * 5) / 96, (winHeight * 20) / 27}
		renderer.FillRect(&base)
	}
	
	base = sdl.Rect{(winWidth * 5) / 64, (winHeight * 20) / 27, (winWidth * 27) / 32, (winHeight * 5) / 54}
	renderer.FillRect(&base)
	base = sdl.Rect{(winWidth * 5) / 192, (winHeight * 5) / 6, (winWidth * 91) / 96, (winHeight * 5) / 54}
	renderer.FillRect(&base)
	renderer.Present()
}

func (a Ring) DrawLiftedRing(place int32) {
	if a.width > 0 {
		ring := sdl.Rect{(winWidth * 31) / 192 - a.n * (winWidth / 96) + (((winWidth * 7) / 24) * place), (winHeight * 4) / 27, a.width, (winHeight * 5) / 108}
		renderer.SetDrawColor(0, 0, 0, 255)
		ring = sdl.Rect{ring.X-1, ring.Y-1, ring.W+2, ring.H+2}
		renderer.DrawRect(&ring)
		renderer.SetDrawColor(225, 200, 0, 255)
		ring = sdl.Rect{ring.X+1, ring.Y+1, ring.W-2, ring.H-2}
		renderer.FillRect(&ring)
		renderer.Present()
	}
}

func (lifted *Ring) Pop(first, second, third *[]Ring, current int32) {
	switch current {
		case 0:
			*lifted = (*first)[len(*first) - 1]
			*first = (*first)[:len(*first) - 1]
		case 1:
			*lifted = (*second)[len(*second) - 1]
			*second = (*second)[:len(*second) - 1]
		case 2:
			*lifted = (*third)[len(*third) - 1]
			*third = (*third)[:len(*third) - 1]
	}
}

func (lifted *Ring) Push(first, second, third *[]Ring, current int32) {
	switch current {
		case 0:
			if len(*first) == 0 || lifted.n < (*first)[len(*first) - 1].n {*first = append(*first, *lifted); *lifted = Ring{0, 0}}
		case 1:
			if len(*second) == 0 || lifted.n < (*second)[len(*second) - 1].n {*second = append(*second, *lifted); *lifted = Ring{0, 0}}
		case 2:
			if len(*third) == 0 || lifted.n < (*third)[len(*third) - 1].n {*third = append(*third, *lifted); *lifted = Ring{0, 0}}
	}
	
}

func DrawRings(first, second, third []Ring) {
	DrawObj()
	var (
		ring sdl.Rect
		count1, count2, count3 int32
	)
	for _, p := range first {
		ring = sdl.Rect{(winWidth * 31) / 192 - p.n * (winWidth / 96), (winHeight * 749) / 1080 - count1 * (winHeight * 5) / 108, p.width, (winHeight * 5) / 108}
		renderer.SetDrawColor(0, 0, 0, 255)
		ring = sdl.Rect{ring.X-1, ring.Y-1, ring.W+2, ring.H+2}
		renderer.DrawRect(&ring)
		renderer.SetDrawColor(225, 200, 0, 255)
		ring = sdl.Rect{ring.X+1, ring.Y+1, ring.W-2, ring.H-2}
		renderer.FillRect(&ring)
		count1++
	}
	for _, p := range second {
		ring = sdl.Rect{(winWidth * 29) / 64 - p.n * (winWidth / 96), (winHeight * 749) / 1080 - count2 * (winHeight * 5) / 108, p.width, (winHeight * 5) / 108}
		renderer.SetDrawColor(0, 0, 0, 255)
		ring = sdl.Rect{ring.X-1, ring.Y-1, ring.W+2, ring.H+2}
		renderer.DrawRect(&ring)
		renderer.SetDrawColor(225, 200, 0, 255)
		ring = sdl.Rect{ring.X+1, ring.Y+1, ring.W-2, ring.H-2}
		renderer.FillRect(&ring)
		count2++
	}
	for _, p := range third {
		ring = sdl.Rect{(winWidth * 143) / 192 - p.n * (winWidth / 96), (winHeight * 749) / 1080 - count3 * (winHeight * 5) / 108, p.width, (winHeight * 5) / 108}
		renderer.SetDrawColor(0, 0, 0, 255)
		ring = sdl.Rect{ring.X-1, ring.Y-1, ring.W+2, ring.H+2}
		renderer.DrawRect(&ring)
		renderer.SetDrawColor(225, 200, 0, 255)
		ring = sdl.Rect{ring.X+1, ring.Y+1, ring.W-2, ring.H-2}
		renderer.FillRect(&ring)
		count3++
	}
	renderer.Present()
}

func NotZero(first, second, third []Ring, current int32) bool {
	switch current {
		case 0:
			return len(first) > 0
		case 1:
			return len(second) > 0
		case 2:
			return len(third) > 0
	}
	return true
}

func Hanoy(first, second, third []Ring) {
	var (
		event sdl.Event
		current int32
		lifted Ring
	)
	
	for {
		event = sdl.WaitEvent()
		switch t := event.(type) {
			case *sdl.KeyboardEvent:
				DrawArrow(current, 255, 255, 255)
				if t.State == 1 {
					switch t.Keysym.Sym {
						case 1073741906: //Up
							if lifted.width == 0 && NotZero(first, second, third, current) {
								lifted.Pop(&first, &second, &third, current)
								DrawRings(first, second, third)
								lifted.DrawLiftedRing(current)
							}
						case 1073741905: //Down
							if lifted.width != 0 {
								lifted.Push(&first, &second, &third, current)
								DrawRings(first, second, third)
								lifted.DrawLiftedRing(current)
							}
						case 1073741904: //Left
							if current > 0 {
								current--
								DrawRings(first, second, third)
								if lifted.width != 0 {lifted.DrawLiftedRing(current)}
							}
						case 1073741903: //Right
							if current < 2 {
								current++
								DrawRings(first, second, third)
								if lifted.width != 0 {lifted.DrawLiftedRing(current)}
							}
						case 27: //Escape
							return
					}
				}
				DrawArrow(current, 0, 200, 100)
		}
	}
}

func main() {
	var first, second, third []Ring
	window, _ = sdl.CreateWindow(winTitle, sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, winWidth, winHeight, sdl.WINDOW_SHOWN)
	renderer, _ = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	defer window.Destroy()
	defer renderer.Destroy()
	defer sdl.Quit()
	
	CreateRings(&first)
	DrawRings(first, second, third)
	DrawArrow(0, 0, 200, 100)
	Hanoy(first, second, third)
}
