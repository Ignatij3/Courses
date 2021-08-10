package main

import (
	"fmt"
	"time"
	"math"
	"image/color"
	_"math/rand"
	"github.com/veandco/go-sdl2/sdl"
)

type (
	Circle struct {
		O Point
		radius int32
	}
	Triangle struct {
		A, B, C Point
	}
	Line struct {
		A, B FPoint
	}
	FPoint struct {
		X, Y float64
	}
	Point struct {
		X, Y int32
	}
)

var (
	winTitle string = "Snowy Night"
	winWidth, winHeight int32 = 1920, 1200
	window *sdl.Window
	renderer *sdl.Renderer
	surface *sdl.Surface
)

func newCircle (c Point, r int32)  Circle  {
	return Circle{c, r}
}	

func (c Circle) drawCircle() {
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

func DrawLine(x1, y1, x2, y2 int, r, g, b uint8) {
	var Lcolor color.Color
	
	if x1 == x2 {
		for targY := y1; targY != y2; {
			surface.Set(x1, targY, Lcolor)
			if y1 < y2 {targY++} else {targY--}
		}
	} else if y1 == y2 {
		for targX := x1; targX < x2; targX++ {
			surface.Set(targX, y1, Lcolor)
		}
	} else {
		var diff [2]float64 = [2]float64{float64(x2 - x1), math.Abs(float64(y2) - float64(y2))}
		if y1 < y2 && diff[0] / diff[1] >= 1 {
			rate := int(math.Ceil(diff[0] / diff[1]))
			for targX, targY := x1, y1; targX < x2; targY++ {
				if y1 == y2 {for {if targX == x2 {return} else {surface.Set(targX, targY, Lcolor); targX++}}}
				for cnt := 0; cnt <= rate; cnt++ {
					targX += cnt
					surface.Set(targX, targY, Lcolor)
				}
			}
		} else if y1 < y2 && diff[0] / diff[1] < 1 {
			rate := int(math.Ceil(diff[1] / diff[0]))
			
		} else if y1 > y2 && diff[0] / diff[1] >= 1 {
			rate := int(math.Ceil(diff[0] / diff[1]))
			
		} else if y1 > y2 && diff[0] / diff[1] < 1 {
			rate := int(math.Ceil(diff[1] / diff[0]))
			
		}
	}
}

func CreateSurfaceObjects() {
	var (
		R sdl.Rect
		t Triangle
		c Circle
	)
	
	R = sdl.Rect{0, winHeight - 50, winWidth, winHeight} //Трава
	surface.FillRect(&R, sdl.MapRGBA((*surface).Format, 0, 130, 0, 255))
	
	R = sdl.Rect{200, winHeight - 450, 600, 400} //Дом
	surface.FillRect(&R, sdl.MapRGBA((*surface).Format, 60, 0, 0, 255))
	
	R = sdl.Rect{250, winHeight - 350, 200, 200} //Стеклопакет
	surface.FillRect(&R, sdl.MapRGBA((*surface).Format, 200, 200, 200, 255))
	
	R = sdl.Rect{270, winHeight - 330, 70, 70} //Окно
	surface.FillRect(&R, sdl.MapRGBA((*surface).Format, 10, 10, 120, 255))
	R = sdl.Rect{360, winHeight - 330, 70, 70} //Окно
	surface.FillRect(&R, sdl.MapRGBA((*surface).Format, 10, 10, 120, 255))
	R = sdl.Rect{270, winHeight - 240, 70, 70} //Окно
	surface.FillRect(&R, sdl.MapRGBA((*surface).Format, 10, 10, 120, 255))
	R = sdl.Rect{360, winHeight - 240, 70, 70} //Окно
	surface.FillRect(&R, sdl.MapRGBA((*surface).Format, 10, 10, 120, 255))
	
	R = sdl.Rect{575, winHeight - 350, 175, 300} //Дверь
	surface.FillRect(&R, sdl.MapRGBA((*surface).Format, 22, 0, 0, 255))
	
	R = sdl.Rect{650, winHeight - 750, 70, 250} //Дымоход
	surface.FillRect(&R, sdl.MapRGBA((*surface).Format, 70, 70, 70, 255))
	R = sdl.Rect{630, winHeight - 785, 110, 35} //Верхушка
	surface.FillRect(&R, sdl.MapRGBA((*surface).Format, 70, 70, 70, 255))
	
	renderer.SetDrawColor(0, 0, 0, 255)
	c.O = Point{700, winHeight - 190} //Дверная ручка
	for i := 15; i > 0; i-- {
		newCircle(c.O, int32(i)).drawCircle()
	}
	
	renderer.SetDrawColor(35, 0, 70, 255)
	t.A = Point{100, winHeight - 450}
	t.B = Point{500, winHeight - 850} //Крыша
	t.C = Point{900, winHeight - 450}
	renderer.DrawLine(t.A.X, t.A.Y, t.B.X, t.B.Y)
	renderer.DrawLine(t.B.X, t.B.Y, t.C.X, t.C.Y)
	renderer.DrawLine(t.A.X, t.A.Y, t.C.X, t.C.Y)
	FillTriangle(t.A, t.B, t.C)
	
	R = sdl.Rect{winWidth - 500, winHeight - 350, 80, 300} //Ствол ели
	surface.FillRect(&R, sdl.MapRGBA((*surface).Format, 18, 0, 0, 255))
	
	renderer.SetDrawColor(0, 90, 0, 255)
	for v := winHeight - 350; v > 420; v -= 150 {
		t.A = Point{winWidth - 700, v}
		t.B = Point{winWidth - 460, v - 240} //Ель
		t.C = Point{winWidth - 220, v}
		renderer.DrawLine(t.A.X, t.A.Y, t.B.X, t.B.Y)
		renderer.DrawLine(t.B.X, t.B.Y, t.C.X, t.C.Y)
		renderer.DrawLine(t.A.X, t.A.Y, t.C.X, t.C.Y)
		FillTriangle(t.A, t.B, t.C)
	}
}

func FillTriangle(A, B, C Point) {
	var k Line
	for i := 0.0; k.B.X != float64(B.X); i += 0.5 {
		k.A = FPoint{float64(A.X) + i, float64(A.Y) - i}
		k.B = FPoint{float64(C.X) - i, float64(A.Y) - i}
		renderer.DrawLine(int32(k.A.X), int32(k.A.Y), int32(k.B.X), int32(k.B.Y))
	}
}

func DrawObjects() {
	var (
		R sdl.Rect
		t Triangle
		c Circle
	)
	
	renderer.SetDrawColor(0, 130, 0, 255)
	R = sdl.Rect{0, winHeight - 50, winWidth, winHeight} //Трава
	renderer.FillRect(&R)
	
	renderer.SetDrawColor(60, 0, 0, 255)
	R = sdl.Rect{200, winHeight - 450, 600, 400} //Дом
	renderer.FillRect(&R)
	
	renderer.SetDrawColor(200, 200, 200, 255)
	R = sdl.Rect{250, winHeight - 350, 200, 200} //Стеклопакет
	renderer.FillRect(&R)
	
	renderer.SetDrawColor(10, 10, 120, 255)
	R = sdl.Rect{270, winHeight - 330, 70, 70} //Окно
	renderer.FillRect(&R)
	R = sdl.Rect{360, winHeight - 330, 70, 70} //Окно
	renderer.FillRect(&R)
	R = sdl.Rect{270, winHeight - 240, 70, 70} //Окно
	renderer.FillRect(&R)
	R = sdl.Rect{360, winHeight - 240, 70, 70} //Окно
	renderer.FillRect(&R)
	
	renderer.SetDrawColor(22, 0, 0, 255)
	R = sdl.Rect{575, winHeight - 350, 175, 300} //Дверь
	renderer.FillRect(&R)
	
	renderer.SetDrawColor(0, 0, 0, 255)
	c.O = Point{700, winHeight - 190} //Дверная ручка
	for i := 15; i > 0; i-- {
		newCircle(c.O, int32(i)).drawCircle()
	}
	
	renderer.SetDrawColor(70, 70, 70, 255)
	R = sdl.Rect{650, winHeight - 750, 70, 250} //Дымоход
	renderer.FillRect(&R)
	R = sdl.Rect{630, winHeight - 785, 110, 35} //Верхушка
	renderer.FillRect(&R)
	
	renderer.SetDrawColor(35, 0, 70, 255)
	t.A = Point{100, winHeight - 450}
	t.B = Point{500, winHeight - 850} //Крыша
	t.C = Point{900, winHeight - 450}
	renderer.DrawLine(t.A.X, t.A.Y, t.B.X, t.B.Y)
	renderer.DrawLine(t.B.X, t.B.Y, t.C.X, t.C.Y)
	renderer.DrawLine(t.C.X, t.C.Y, t.A.X, t.A.Y)
	FillTriangle(t.A, t.B, t.C)
	
	renderer.SetDrawColor(18, 0, 0, 255)
	R = sdl.Rect{winWidth - 500, winHeight - 350, 80, 300} //Ствол ели
	renderer.FillRect(&R)
	
	renderer.SetDrawColor(0, 90, 0, 255)
	for v := winHeight - 350; v > 420; v -= 150 {
		t.A = Point{winWidth - 700, v}
		t.B = Point{winWidth - 460, v - 240} //Ель
		t.C = Point{winWidth - 220, v}
		renderer.DrawLine(t.A.X, t.A.Y, t.B.X, t.B.Y)
		renderer.DrawLine(t.B.X, t.B.Y, t.C.X, t.C.Y)
		renderer.DrawLine(t.C.X, t.C.Y, t.A.X, t.A.Y)
		FillTriangle(t.A, t.B, t.C)
	}
}

func Process() {
	for {
		go DropSnow()
		time.Sleep(100 * time.Millisecond)
	}
}

func DropSnow() {
	
}

func main() {
	window, _ = sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, winWidth, winHeight, sdl.WINDOW_SHOWN)
	renderer, _ = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	surface, _ = window.GetSurface()
	surface.Free()
	
	CreateSurfaceObjects()
	//DrawObjects()
	//renderer.Present()
	window.UpdateSurface()
	
	/*rand.Seed(time.Now().UnixNano())
	sdl.Main(func() {Process()})*/
	
	color1 := surface.At(270, int(winHeight - 330))
	color2 := surface.At(500, int(winHeight - 10))
	fmt.Printf("%v color at %d x, %d y\n", color1, 100, 100)
	fmt.Printf("%v color at %d x, %d y\n", color2, 5, 5)
	sdl.Delay(5000)
	
	renderer.Destroy()
	window.Destroy()
	sdl.Quit()
}
