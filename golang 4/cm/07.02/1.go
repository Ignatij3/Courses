package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"os"
	"math/rand"
	"time"
)

var winTitle string = "Sierpinski triangle"
var winWidth, winHeight int32 = 800, 600

func run() int {
	var window *sdl.Window
	var renderer *sdl.Renderer
	window, err := sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		winWidth, winHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
		return 1
	}
	defer window.Destroy()
	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
		return 2
	}
	defer renderer.Destroy()
	renderer.SetDrawColor(255, 255, 255, 255) 
	renderer.Clear()
	
	var  (
		p sdl.Point
		A [3]sdl.Point
	)
	
	rand.Seed(time.Now().UnixNano())
	A[0].X = rand.Int31n(winWidth/2)		// первая вершина треугольника
	A[0].Y = rand.Int31n(winHeight/2)		// располагается в левой верхней четверти
	A[1].X = rand.Int31n(winWidth/2) + (winWidth/2)	// вторая вершина треугольника
	A[1].Y = rand.Int31n(winHeight/2)		// располагается в правой верхней четверти
	A[2].X = rand.Int31n(winWidth)		// третья вершина треугольника располагается
	A[2].Y = rand.Int31n(winHeight/2) + winHeight/2	//  в нижней половине
	p.X = (A[0].X + A[1].X + A[2].X)/3
	p.Y = (A[0].Y + A[1].Y + A[2].Y)/3
	renderer.SetDrawColor(0, 0, 0, 255)  // рисуем чёрным

	for {
		event := sdl.PollEvent() 
		if _, ok := event.(*sdl.MouseButtonEvent); ok {
			break
		}
		renderer.DrawPoint(p.X, p.Y)
		renderer.Present()
		i:= rand.Int31n(3)
		p.X = (p.X + A[i].X)/2
		p.Y = (p.Y + A[i].Y)/2
	}
	return 0
}
func main() {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		fmt.Printf("Failed to initialize SDL: %s\n", err)
		return
	}
	run()
	sdl.Quit()
}
