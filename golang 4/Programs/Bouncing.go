package main

import (
	"time"
	"math"
	"math/rand"
	"github.com/veandco/go-sdl2/sdl"
)

var (
	winTitle string = "Bouncing Ball"
	winWidth, winHeight int32 = 1200, 720
	window *sdl.Window
	renderer *sdl.Renderer
	r int32 = 15
)

type (
	Clear struct {
		back Point
	}
	Circle struct {
		O Point
		radius int32
	}
	Point struct {
		X, Y int32
	}
)

func newCircle (c Point, r int32)  Circle  {
	return Circle{c, r}
}	

func (c Circle) drawCircle () {
	p:= make([]Point, 0)
	x, y := float64(c.radius), 0.0
	fi:= 1.0/float64(c.radius)
	cos, sin := math.Cos(fi), math.Sin(fi)
	for x > y  {
		p = append(p, Point{int32(math.Round(x)), int32(math.Round(y))})	
		x, y = x*cos - y*sin, x*sin + y*cos
	}	
	
	pp := make([]sdl.Point, len(p))
	for i, v:= range p  {
		pp[i] = sdl.Point{c.O.X+v.X, c.O.Y+v.Y}
	}	
	renderer.DrawLines(pp)
	for i, v:= range p  {
		pp[i] = sdl.Point{c.O.X-v.X, c.O.Y+v.Y}
	}	
	renderer.DrawLines(pp)
	for i, v:= range p  {
		pp[i] = sdl.Point{c.O.X-v.Y, c.O.Y+v.X}
	}	
	renderer.DrawLines(pp)
	for i, v:= range p  {
		pp[i] = sdl.Point{c.O.X-v.Y, c.O.Y-v.X}
	}	
	renderer.DrawLines(pp)
	for i, v:= range p  {
		pp[i] = sdl.Point{c.O.X-v.X, c.O.Y-v.Y}
	}	
	renderer.DrawLines(pp)
	for i, v:= range p  {
		pp[i] = sdl.Point{c.O.X+v.X, c.O.Y-v.Y}
	}	
	renderer.DrawLines(pp)
	for i, v:= range p  {
		pp[i] = sdl.Point{c.O.X+v.Y, c.O.Y-v.X}
	}	
	renderer.DrawLines(pp)
	for i, v:= range p  {
		pp[i] = sdl.Point{c.O.X+v.Y, c.O.Y+v.X}
	}	
	renderer.DrawLines(pp)
	return 
}

func Move(cp Point) {
	var (
		Vx, Vy int32
		dir string = "UPLEFT"
		less bool
		c Circle
		b Clear
	)
	Vx, Vy = -3, -3
	c.O = cp
	for {
		event := sdl.PollEvent() 
		if _, ok := event.(*sdl.MouseButtonEvent); ok {
			break
		}
		
		switch dir {
			case "UPLEFT" :
				if c.O.X - r <= 0 {
					dir = "UPRIGHT"
					Vx, Vy = 2, -2
				} else if c.O.Y - r <= 0 {
					dir = "DOWNLEFT"
					Vx, Vy = -2, 2
				}
			case "UPRIGHT" :			
				if c.O.X + r >= winWidth {
					dir = "UPLEFT"
					Vx, Vy = -2, -2
				} else if c.O.Y - r <= 0 {
					dir = "DOWNRIGHT"
					Vx, Vy = 2, 2
				}
			case "DOWNLEFT" :
				if c.O.X - r <= 0 {
					dir = "DOWNRIGHT"
					Vx, Vy = 2, 2
				} else if c.O.Y + r >= winHeight {
					dir = "UPLEFT"
					Vx, Vy = -2, -2
				}
			case "DOWNRIGHT" :
				if c.O.X + r >= winWidth {
					dir = "DOWNLEFT"
					Vx, Vy = -2, 2
				} else if c.O.Y + r >= winHeight {
					dir = "UPRIGHT"
					Vx, Vy = 2, -2
				}
		}
		c.O.X = c.O.X + Vx
		c.O.Y = c.O.Y + Vy
		renderer.SetDrawColor(170, 0, 20, 255)
		renderer.Present()
		newCircle(c.O, r).drawCircle()
		renderer.SetDrawColor(255, 255, 255, 255)
		switch dir {
			case "UPLEFT" :
				b.back = Point{c.O.X - (Vx * 40), c.O.Y - (Vy * 40)}
			case "UPRIGHT" :			
				b.back = Point{c.O.X - (Vx * 40), c.O.Y - (Vy * 40)}
			case "DOWNLEFT" :
				b.back = Point{c.O.X - (Vx * 40), c.O.Y - (Vy * 40)}
			case "DOWNRIGHT" :
				b.back = Point{c.O.X - (Vx * 40), c.O.Y - (Vy * 40)}
		}
		newCircle(b.back, r).drawCircle()
		if r == 15 {less = true} else if r == 3 {less = false}
		if less {r--} else {r++}
		sdl.Delay(10)
	}
	
}

func main() {
	var c Circle
	
	rand.Seed(time.Now().UnixNano())
	window, _ = sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, winWidth, winHeight, sdl.WINDOW_SHOWN)
	renderer, _ = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	
	renderer.SetDrawColor(255, 255, 255, 255)
	renderer.Clear()
	
	c.O = Point{rand.Int31n(winWidth - 15), rand.Int31n(winHeight - 15)}
	newCircle(c.O, r).drawCircle()
	renderer.Present()
	
	Move(c.O)
	sdl.Quit()
}
