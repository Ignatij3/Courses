package main

import (
	"time"
	"math"
	"math/rand"
	"github.com/veandco/go-sdl2/sdl"
)

var (
	winTitle string = "Bouncing Ball 2"
	window *sdl.Window
	renderer *sdl.Renderer
	lineS [5]Line
)

const (
	r int32 = 8
	winWidth, winHeight int32 = 1200, 720
)

type (
	Line struct {
		A, B Point
	}
	Circle struct {
		O FPoint
		radius int32
	}
	FPoint struct {
		X, Y float64
	}
	Point struct {
		X, Y int32
	}
)

func newCircle (c FPoint, r int32)  Circle  {
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
		pp[i] = sdl.Point{int32(c.O.X)+v.X, int32(c.O.Y)+v.Y}
	}	
	renderer.DrawLines(pp)
	for i, v:= range p  {
		pp[i] = sdl.Point{int32(c.O.X)-v.X, int32(c.O.Y)+v.Y}
	}	
	renderer.DrawLines(pp)
	for i, v:= range p  {
		pp[i] = sdl.Point{int32(c.O.X)-v.Y, int32(c.O.Y)+v.X}
	}	
	renderer.DrawLines(pp)
	for i, v:= range p  {
		pp[i] = sdl.Point{int32(c.O.X)-v.Y, int32(c.O.Y)-v.X}
	}	
	renderer.DrawLines(pp)
	for i, v:= range p  {
		pp[i] = sdl.Point{int32(c.O.X)-v.X, int32(c.O.Y)-v.Y}
	}	
	renderer.DrawLines(pp)
	for i, v:= range p  {
		pp[i] = sdl.Point{int32(c.O.X)+v.X, int32(c.O.Y)-v.Y}
	}	
	renderer.DrawLines(pp)
	for i, v:= range p  {
		pp[i] = sdl.Point{int32(c.O.X)+v.Y, int32(c.O.Y)-v.X}
	}	
	renderer.DrawLines(pp)
	for i, v:= range p  {
		pp[i] = sdl.Point{int32(c.O.X)+v.Y, int32(c.O.Y)+v.X}
	}	
	renderer.DrawLines(pp)
	return 
}

func DrawWalls() [9]Line {
	var (
		s Line
		lines [9]Line
	)
	
	renderer.SetDrawColor(0, 0, 0, 255)
	for i := 0; i < 9; i++ {
		if i < 5 {
			s.A = Point{rand.Int31n(winWidth), rand.Int31n(winHeight)}
			s.B = Point{rand.Int31n(winWidth), rand.Int31n(winHeight)}
			lines[i] = Line{s.A, s.B}
			renderer.DrawLine(s.A.X, s.A.Y, s.B.X, s.B.Y)
		} else if i == 5 {
			s.A = Point{0, 0}
			s.B = Point{0, winHeight - 1}
			lines[i] = Line{s.A, s.B}
			renderer.DrawLine(s.A.X, s.A.Y, s.B.X, s.B.Y)
		} else if i == 6 {
			s.A = Point{0, winHeight - 1}
			s.B = Point{winWidth - 1, winHeight - 1}
			lines[i] = Line{s.A, s.B}
			renderer.DrawLine(s.A.X, s.A.Y, s.B.X, s.B.Y)
		} else if i == 7 {
			s.A = Point{winWidth - 1, winHeight - 1}
			s.B = Point{winWidth - 1, 0}
			lines[i] = Line{s.A, s.B}
			renderer.DrawLine(s.A.X, s.A.Y, s.B.X, s.B.Y)
		} else if i == 8 {
			s.A = Point{winWidth - 1, 0}
			s.B = Point{0, 0}
			lines[i] = Line{s.A, s.B}
			renderer.DrawLine(s.A.X, s.A.Y, s.B.X, s.B.Y)
		}
	}
	return lines
}

func Move(start FPoint, lines [9]Line) {
	var (
		c Circle
		iPrev int
		Vx, Vy int32
		Cx, Cy, Hx, Hy, Dx, Dy, Kx, Ky, Ux, Uy float64
	)
	c.O = start
	Ux, Uy = 1, 1
	
	for {
		T, index := Cross(c.O.X, c.O.Y, Ux, Uy, lines)
		
		Vx = lines[index].B.X - lines[index].A.X
		Vy = lines[index].B.Y - lines[index].A.Y
		Cx = c.O.X + Ux * T
		Cy = c.O.Y + Uy * T
		
		for ; T >= 1; T-- { //Make tail
			c.O.X += Ux
			c.O.Y += Uy
			
			renderer.SetDrawColor(170, 0, 20, 255)
			newCircle(c.O, r).drawCircle()
			renderer.SetDrawColor(0, 0, 0, 255)
			renderer.DrawLine(lines[iPrev].A.X, lines[iPrev].A.Y, lines[iPrev].B.X, lines[iPrev].B.Y)
			renderer.Present()
			sdl.Delay(3)
			renderer.SetDrawColor(255, 255, 255, 255)
			newCircle(c.O, r).drawCircle()
		}
		iPrev = index
		c.O.X += Ux*T
		c.O.Y += Uy*T
		
		Dx, Dy = Cx + Ux, Cy + Uy
		Hx, Hy = Perpendicular(Dx, Dy, lines[index].A, float64(Vx), float64(Vy))
		Kx, Ky = Hx * 2 - Dx, Hy * 2 - Dy
		Ux, Uy = Kx - Cx, Ky - Cy
	}
}

func Cross(X, Y, Ux, Uy float64, lines [9]Line) (float64, int) {
	var (
		iR int
		cross bool
		tRay, T float64 
	)
	
	tRay = 1.0E30
	for i, l := range(lines) {
		if cross, T = Ray(X, Y, Ux, Uy, l); cross {
			if tRay > T {
				iR = i
				tRay = T
			}
		}
	}
	return tRay, iR
}

func Ray(X, Y, Ux, Uy float64, line Line) (bool, float64) {
	var (
		cross bool
		t, T, Vx, Vy float64
	)
	
	Vx = float64(line.B.X) - float64(line.A.X)
	Vy = float64(line.B.Y) - float64(line.A.Y)
	
	if Vx * Uy == Vy * Ux {return false, 1.0e15}
	t = ((X - float64(line.A.X)) * Uy - (Y - float64(line.A.Y)) * Ux) / (Vx * Uy - Vy * Ux) //Отрезок
	T = ((X - float64(line.A.X)) * Vy - (Y - float64(line.A.Y)) * Vx) / (Vx * Uy - Vy * Ux) //Луч
	cross = t > 0 && t < 1 && T > 1.0e-4 
	
	return cross, T
}

func Perpendicular(Px, Py float64, A Point, Vx, Vy float64) (float64, float64)  {
	var PxVx2, AxVy2, PyVy2, AyVx2, VxVy, Vx2Vy2, Hx, Hy float64
	
	PxVx2 = Px * Vx * Vx
	PyVy2 = Py * Vy * Vy
	AxVy2 = float64(A.X) * Vy * Vy
	AyVx2 = float64(A.Y) * Vx * Vx
	VxVy = Vx * Vy
	Vx2Vy2 = (Vx * Vx) + (Vy * Vy)
	
	Hx = (PxVx2 + AxVy2 + VxVy * (Py - float64(A.Y))) / (Vx2Vy2)
	Hy = (PyVy2 + AyVx2 + VxVy * (Px - float64(A.X))) / (Vx2Vy2)
	return Hx, Hy
}

func main() {
	var c Circle
	
	rand.Seed(time.Now().UnixNano())
	window, _ = sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, winWidth, winHeight, sdl.WINDOW_SHOWN)
	renderer, _ = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	
	renderer.SetDrawColor(255, 255, 255, 255)
	renderer.Clear()
	
	lines := DrawWalls()
	c.O = FPoint{float64(rand.Int31n(winWidth - r * 2) + r), float64(rand.Int31n(winHeight - r * 2) + r)}
	Move(c.O, lines)
	sdl.Quit()
}
