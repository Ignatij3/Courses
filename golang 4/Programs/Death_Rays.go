package main

import (
	"time"
	"math"
	"math/rand"
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
	Snowflakes struct {
		A Point
	}
	FPoint struct {
		X, Y float64
	}
	Point struct {
		X, Y int32
	}
	Job struct {
		id int
		n int64
		from, to int64
	}
)

var (
	winTitle string = "Snowy Night"
	winWidth, winHeight int32 = 1920, 1080
	window *sdl.Window
	renderer *sdl.Renderer
	jobs = make(chan Job, 100)
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
	R = sdl.Rect{200, winHeight - 400, 600, 500} //Дом
	renderer.FillRect(&R)
	
	renderer.SetDrawColor(200, 200, 200, 255)
	R = sdl.Rect{250, winHeight - 300, 200, 200} //Стеклопакет
	renderer.FillRect(&R)
	
	renderer.SetDrawColor(10, 10, 120, 255)
	R = sdl.Rect{270, winHeight - 280, 70, 70} //Окно
	renderer.FillRect(&R)
	R = sdl.Rect{360, winHeight - 280, 70, 70} //Окно
	renderer.FillRect(&R)
	R = sdl.Rect{270, winHeight - 190, 70, 70} //Окно
	renderer.FillRect(&R)
	R = sdl.Rect{360, winHeight - 190, 70, 70} //Окно
	renderer.FillRect(&R)
	
	renderer.SetDrawColor(22, 0, 0, 255)
	R = sdl.Rect{575, winHeight - 300, 175, 300} //Дверь
	renderer.FillRect(&R)
	
	renderer.SetDrawColor(0, 0, 0, 255)
	c.O = Point{700, winHeight - 140} //Дверная ручка
	for i := 15; i > 0; i-- {
		newCircle(c.O, int32(i)).drawCircle()
	}
	
	renderer.SetDrawColor(70, 70, 70, 255)
	R = sdl.Rect{650, winHeight - 700, 70, 250} //Труба
	renderer.FillRect(&R)
	R = sdl.Rect{630, winHeight - 735, 110, 35} //Дымоход
	renderer.FillRect(&R)
	
	renderer.SetDrawColor(35, 0, 70, 255)
	t.A = Point{100, winHeight - 400}
	t.B = Point{500, winHeight - 800} //Крыша
	t.C = Point{900, winHeight - 400}
	renderer.DrawLine(t.A.X, t.A.Y, t.B.X, t.B.Y)
	renderer.DrawLine(t.B.X, t.B.Y, t.C.X, t.C.Y)
	renderer.DrawLine(t.C.X, t.C.Y, t.A.X, t.A.Y)
	FillTriangle(t.A, t.B, t.C)
	
	renderer.SetDrawColor(18, 0, 0, 255)
	R = sdl.Rect{winWidth - 500, winHeight - 300, 80, 300} //Ствол ели
	renderer.FillRect(&R)
	
	renderer.SetDrawColor(0, 90, 0, 255)
	for v := winHeight - 300; v > 420; v -= 150 {
		t.A = Point{winWidth - 700, v}
		t.B = Point{winWidth - 460, v - 240} //Ель
		t.C = Point{winWidth - 220, v}
		renderer.DrawLine(t.A.X, t.A.Y, t.B.X, t.B.Y)
		renderer.DrawLine(t.B.X, t.B.Y, t.C.X, t.C.Y)
		renderer.DrawLine(t.C.X, t.C.Y, t.A.X, t.A.Y)
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

func Process() {
	var chanSlice []chan int
	
	for i := 0; i < 10; i++ {
		chanSlice = append(chanSlice, make(chan int, 1))
	}
	for {
		event := sdl.PollEvent() 
		if _, ok := event.(*sdl.MouseButtonEvent); ok {
			break
			sdl.Quit()
		}
		for _, ch := range chanSlice {
			go DropSnow(ch)
		}
		for _, ch2 := range chanSlice {
			<-ch2
		}
		renderer.Present()
		sdl.Delay(60)
	}
}

func DropSnow(quit chan int) {
	var (
		Vy, Vx int32 = 1, 2
		move bool = true
		s Snowflakes
	)
	s.A = Point{0, rand.Int31n(winHeight)}
	for move {
		s.A = Point{s.A.X + Vx, s.A.Y + Vy}
		sdl.Do(func() {
			renderer.SetDrawColor(255, 0, 0, 255)
			renderer.DrawPoint(s.A.X, s.A.Y)
		})
		if s.A.X >= winWidth - 1 {move = false}
	}
	quit <- 1
}

func main() {
	
	window, _ = sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, winWidth, winHeight, sdl.WINDOW_SHOWN)
	renderer, _ = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	
	renderer.SetDrawColor(150, 255, 255, 255)
	renderer.Clear()
	rand.Seed(time.Now().UnixNano())
	
	DrawObjects()
	renderer.Present()
	
	sdl.Main(func() {Process()})
}
