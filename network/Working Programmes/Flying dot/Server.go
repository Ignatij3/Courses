package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/veandco/go-sdl2/sdl"
)

func makeWindow(width, height int32) (*sdl.Window, *sdl.Renderer) {
	window, _ := sdl.CreateWindow("Bouncing dot", sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, width, height, sdl.WINDOW_SHOWN)
	renderer, _ := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)

	renderer.SetDrawColor(255, 255, 255, 255)
	renderer.Clear()
	renderer.Present()
	return window, renderer
}

func getDimensions(conn *net.UDPConn) (*sdl.Window, *sdl.Renderer) {
	var (
		data          [9]byte
		coords        []string
		width, height int
	)

	_, err := conn.Read(data[:])
	fmt.Println(string(data[:]))
	if err != nil {
		fmt.Println("Error reading dimensions")
		return nil, nil
	}

	trimmed := strings.TrimRight(string(data[:]), "\x00")
	coords = strings.Split(trimmed, ".")
	width, _ = strconv.Atoi(coords[0])
	height, _ = strconv.Atoi(coords[1])
	window, renderer := makeWindow(int32(width), int32(height))

	return window, renderer
}

func convToFloat(data string) (int32, int32) {
	trimmed := strings.TrimRight(data, "\x00")
	splitData := strings.Split(trimmed, ".")
	x, _ := strconv.Atoi(splitData[0])
	y, _ := strconv.Atoi(splitData[1])
	fmt.Printf("data - %v, split1 - %v, split2 - %v, x - %v, y - %v\n", data, splitData[0], splitData[1], int32(x), int32(y))
	return int32(x), int32(y)
}

func drawDot(x, y int32, renderer *sdl.Renderer) {
	renderer.DrawPoint(x, y)
	renderer.Present()
}

func changePosition(renderer *sdl.Renderer, conn *net.UDPConn) {
	var (
		data [9]byte
	)

	_, err := conn.Read(data[:])
	if err != nil {
		fmt.Println("Error reading position")
		return
	}
	x, y := convToFloat(string(data[:]))
	drawDot(x, y, renderer)
}

//StartServer starts server
func StartServer() {
	var event sdl.Event
	port := ":8081"
	protocol := "udp"

	adr, err := net.ResolveUDPAddr(protocol, port)
	if err != nil {
		fmt.Println("Wrong addr")
		return
	}

	fmt.Println("Server working on addr " + adr.String())

	conn, err := net.ListenUDP(protocol, adr)
	if err != nil {
		fmt.Println("Error creating server")
		return
	}

	window, renderer := getDimensions(conn)
	renderer.SetDrawColor(255, 0, 0, 255)

	for {
		event = sdl.WaitEventTimeout(1)
		switch t := event.(type) {
		case *sdl.KeyboardEvent:
			if t.State == 1 {
				switch t.Keysym.Sym {
				case 27: //Escape
					break
				}
			}
		}
		changePosition(renderer, conn)
	}
	window.Destroy()
	renderer.Destroy()
	sdl.Quit()
}

func main() {
	StartServer()
}
