package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"math"
)

var winTitle string = "Circle"
var winWidth, winHeight int32 = 800, 600

type circle struct  {
	center sdl.Point
	radius int32
}	

func newCircle (c sdl.Point, r int32)  circle  {
	return circle{c, r}
}	

func (circum circle) drawCircle (rend *sdl.Renderer) {
	p:= make([]sdl.Point, 0)
	x, y := float64(circum.radius), 0.0		// starting point
	fi:= 1.0/float64(circum.radius)			// angle between adjacent vertices
	cos, sin := math.Cos(fi), math.Sin(fi)
	// build 1/8 of circum
	for x > y  {
		p = append(p, sdl.Point{int32(math.Round(x)), int32(math.Round(y))})	
		x, y = x*cos - y*sin, x*sin + y*cos
	}	
	
	pp:= make([]sdl.Point, len(p))
	for i, v:= range p  {
		pp[i] = sdl.Point{circum.center.X+v.X, circum.center.Y+v.Y}
	}	
	rend.DrawLines(pp)
	for i, v:= range p  {
		pp[i] = sdl.Point{circum.center.X-v.X, circum.center.Y+v.Y}
	}	
	rend.DrawLines(pp)
	for i, v:= range p  {
		pp[i] = sdl.Point{circum.center.X-v.Y, circum.center.Y+v.X}
	}	
	rend.DrawLines(pp)
	for i, v:= range p  {
		pp[i] = sdl.Point{circum.center.X-v.Y, circum.center.Y-v.X}
	}	
	rend.DrawLines(pp)
	for i, v:= range p  {
		pp[i] = sdl.Point{circum.center.X-v.X, circum.center.Y-v.Y}
	}	
	rend.DrawLines(pp)
	for i, v:= range p  {
		pp[i] = sdl.Point{circum.center.X+v.X, circum.center.Y-v.Y}
	}	
	rend.DrawLines(pp)
	for i, v:= range p  {
		pp[i] = sdl.Point{circum.center.X+v.Y, circum.center.Y-v.X}
	}	
	rend.DrawLines(pp)
	for i, v:= range p  {
		pp[i] = sdl.Point{circum.center.X+v.Y, circum.center.Y+v.X}
	}	
	rend.DrawLines(pp)
	
	return 
}

func main() {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		fmt.Printf("Failed to initialize SDL: %s\n", err)
		return
	}
	wind, err := sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		winWidth, winHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Printf("Failed to create window: %s\n", err)
	}
	defer wind.Destroy()
	
	rend, err := sdl.CreateRenderer(wind, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Printf("Failed to create renderer: %s\n", err)
	}
	defer rend.Destroy()
	rend.SetDrawColor(255, 255, 255, 255)
	rend.Clear()
	rend.Present()

	rend.SetDrawColor(0, 0, 0, 255)
	for r:= 20; r<=200; r += 20 {
		newCircle(sdl.Point{winWidth/2, winHeight/2}, int32(r)).drawCircle(rend)
	}
	rend.Present()
	
	sdl.Delay(5000)
	sdl.Quit()
}
