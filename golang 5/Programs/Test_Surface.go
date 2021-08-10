package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	var (
		winTitle string = "Surface At"
		winWidth, winHeight int32 = 1200, 720
	)
	
	window, _ := sdl.CreateWindow(winTitle, sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, winWidth, winHeight, sdl.WINDOW_SHOWN)
	renderer, _ := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	surface, _ := window.GetSurface()
	
	renderer.SetDrawColor(255, 255, 255, 255)
	renderer.Clear()
	surface.FillRect(&sdl.Rect{0, 0, winWidth, winHeight}, sdl.MapRGBA((*surface).Format, 255, 255, 255, 255)) //Change color for an entire window
	surface.FillRect(&sdl.Rect{50, 50, 600, 300}, sdl.MapRGBA((*surface).Format, 35, 80, 0, 180)) //Change color from x - {50; 600}, y - {50; 300}
	
	///1-й вариант, это работает как renderer.Present(), разве что для surface
	window.UpdateSurface()
	
	///Альтернативный вариант
	//texture, _ := renderer.CreateTextureFromSurface(surface) ///Создаю текстуру, точно не знаю, зачем это надо, но без этого ничего не выведется
	//renderer.Copy(texture, nil, nil) //Как я понял, я передаю текстуру рендереру
	//renderer.Present() //Вывожу на экран
	//texture.Destroy()
	
	color1 := surface.At(100, 100)
	color2 := surface.At(5, 5)
	fmt.Printf("%v color at %d x, %d y\n", color1, 100, 100)
	fmt.Printf("%v color at %d x, %d y\n", color2, 5, 5)
	
	surface.Free()
	renderer.Destroy() //Удивительный факт, если не закрыть/уничтожить/избавиться от этих всех вещей в конце, окно зафризит и вылетит
	sdl.Delay(4000)
	window.Destroy()
	sdl.Quit()
}
