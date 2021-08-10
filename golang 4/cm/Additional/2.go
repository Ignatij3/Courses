package main

import (
	"fmt"
	"math"
	"time"
	"github.com/veandco/go-sdl2/sdl"
)

var winTitle string = "Flags"
var winWidth, winHeight int32 = 800, 600

type (
	Flag struct {
		cx, cy float64
		height float64
	}	
	Fpoint struct {
		X, Y float64
	}
)
func newFlag (x, y float64, h float64) Flag {
	return Flag{x, y, h}
}	

func RoundP(f []Fpoint) []sdl.Point {
	pp := make([]sdl.Point, 0)
	for _, p := range f  {
		pp = append(pp, sdl.Point{int32(math.Round(p.X)), int32(math.Round(p.Y))})
	}		
	return pp
}	
	
func (f Flag) Run (rend *sdl.Renderer, quit chan int) {
	p:= make([]Fpoint, 0)
	p = append(p, Fpoint{f.cx, f.cy})
	p = append(p, Fpoint{f.cx + f.height, f.cy})
	p = append(p, Fpoint{f.cx + f.height, f.cy+f.height/2})
	p = append(p, Fpoint{f.cx + f.height/2, f.cy+f.height/2})
	p = append(p, Fpoint{f.cx + f.height/2, f.cy})
	
	fi:= 2*math.Pi/150
	cos, sin := math.Cos(fi), math.Sin(fi)
	dx, dy := f.cx*(1-cos) + f.cy*sin, -f.cx*sin + f.cy*(1-cos)

	start:= time.Now()
	for t:= start; t.Sub(start) < 5000*time.Millisecond; t = time.Now()  {
		sdl.Do( 
			func() {
				rend.SetDrawColor(0,0,0,255)
				rend.DrawLines(RoundP(p))
				rend.Present()
		} )
		sdl.Delay(10)
		sdl.Do( 
			func() {
				rend.SetDrawColor(255,255,255,255)
				rend.DrawLines(RoundP(p))
				rend.Present()
		} )
		for i:= 1; i<len(p); i++  {
			p[i].X, p[i].Y = p[i].X*cos - p[i].Y*sin + dx, p[i].X*sin + p[i].Y*cos + dy
		}					 
	}	
	quit <- 0
}


func process(rend *sdl.Renderer)  {
	rend.SetDrawColor(255, 255, 255, 255)
	rend.Clear()
	rend.Present()
	waiting1 := make(chan int, 0)
	waiting2 := make(chan int, 0)
	go newFlag(300, 150, 30).Run(rend, waiting1)
	go newFlag(100, 250, 30).Run(rend, waiting2)
	<-waiting1
	<-waiting2
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

	sdl.Main(func() {process(rend)})
}
