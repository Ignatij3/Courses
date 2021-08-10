package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

var winTitle string = "package sdl. Test Window"
var winWidth, winHeight int32 = 400, 240

func main() {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		fmt.Printf("Failed to initialize SDL: %s\n", err)
		return
	}
	defer sdl.Quit()

	var window *sdl.Window
	window, err = sdl.CreateWindow(winTitle, sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_UNDEFINED, winWidth, winHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Printf("Failed to create window: %s\n", err)
		return
	}
	defer window.Destroy()
	sdl.Delay(3000)
}
