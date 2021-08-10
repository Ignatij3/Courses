package main

import (
	"fmt"
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
	Lines struct {
		A, B []Point
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
	snow []Point
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

func DrawObjects() Lines {
	var (
		R sdl.Rect
		t Triangle
		c Circle
		l Lines
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
	l.A = append(l.A, t.A)
	l.B = append(l.B, t.B)
	l.A = append(l.A, t.B)
	l.B = append(l.B, t.C)
	
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
		l.A = append(l.A, t.A)
		l.B = append(l.B, t.B)
		l.A = append(l.A, t.B)
		l.B = append(l.B, t.C)
	}
	
	t.A = Point{200, winHeight - 450} //Дом - лево
	t.B = Point{200, winHeight - 50}
	l.A = append(l.A, t.A)
	l.B = append(l.B, t.B)
	t.A = Point{800, winHeight - 450} //Дом - право
	t.B = Point{800, winHeight - 50}
	l.A = append(l.A, t.A)
	l.B = append(l.B, t.B)
	
	t.A = Point{650, winHeight - 750} //Дымоход - лево
	t.B = Point{650, winHeight - 500}
	l.A = append(l.A, t.A)
	l.B = append(l.B, t.B)
	t.A = Point{720, winHeight - 750} //Дымоход - право
	t.B = Point{720, winHeight - 500}
	l.A = append(l.A, t.A)
	l.B = append(l.B, t.B)
	
	t.A = Point{630, winHeight - 785} //Верхушка - лево
	t.B = Point{630, winHeight - 750}
	l.A = append(l.A, t.A)
	l.B = append(l.B, t.B)
	t.A = Point{720, winHeight - 785} //Верхушка - право
	t.B = Point{720, winHeight - 750}
	l.A = append(l.A, t.A)
	l.B = append(l.B, t.B)
	
	t.A = Point{winWidth - 500, winHeight - 350} //Ствол ели - лево
	t.B = Point{winWidth - 420, winHeight - 50}
	l.A = append(l.A, t.A)
	l.B = append(l.B, t.B)
	t.A = Point{winWidth - 500, winHeight - 350} //Ствол ели - право
	t.B = Point{winWidth - 420, winHeight - 50}
	l.A = append(l.A, t.A)
	l.B = append(l.B, t.B)
	
	return l
}

func FillTriangle(A, B, C Point) {
	var k Line
	for i := 0.0; k.B.X != float64(B.X); i += 0.5 {
		k.A = FPoint{float64(A.X) + i, float64(A.Y) - i}
		k.B = FPoint{float64(C.X) - i, float64(A.Y) - i}
		renderer.DrawLine(int32(k.A.X), int32(k.A.Y), int32(k.B.X), int32(k.B.Y))
	}
}

func Process(l Lines) {
	for {
		go DropSnow(l)
		time.Sleep(100 * time.Millisecond)
	}
}

func DropSnow(l Lines) {
	var (
		moveL, move, noObj, downS, leftS, rightS, leftL, rightL, vertical bool = true, true, true, false, false, false, false, false, false
		Flake Point
	)
	
	Flake = Point{rand.Int31n(winWidth), 0}
	//Flake = Point{620, 0}
	for move {
		if noObj {
			sdl.Do(func() {
				renderer.SetDrawColor(0, 0, 0, 255)
				renderer.DrawPoint(Flake.X, Flake.Y)
			})
			Flake = Point{Flake.X, Flake.Y + 1}
			sdl.Do(func() {
				renderer.SetDrawColor(125, 255, 240, 255)
				renderer.DrawPoint(Flake.X, Flake.Y)
				if rand.Intn(5) == 0 {renderer.Present()}
			})
		} else {noObj = true}
		
		if Flake.Y >= winHeight - 51 || (Flake.Y == winHeight - 786 && Flake.X >= 630 && Flake.X < 740) {move = false}
		for downS, leftS, rightS = Collision(Flake.X, Flake.Y); downS || leftS || rightS; {
			noObj = false
			for index := len(l.A) - 8; index < len(l.A); index++ {
				vertical = Compare(l.A[index], l.B[index], Flake.X, Flake.Y, Flake.X + 1, Flake.Y + 1)
				if vertical {move, moveL = false, false; break}
			}
			if vertical {break}
			
			for index := 0; index < len(l.A); index++ {
				if Compare(l.A[index], l.B[index], Flake.X, Flake.Y, Flake.X + 1, Flake.Y + 1) && !rightL {rightL = true}
				if Compare(l.A[index], l.B[index], Flake.X, Flake.Y, Flake.X - 1, Flake.Y + 1) && !rightL {leftL = true}
			}
			downS, leftS, rightS = Collision(Flake.X, Flake.Y)
			
			sdl.Do(func() {
				renderer.SetDrawColor(0, 0, 0, 255)
				renderer.DrawPoint(Flake.X, Flake.Y)
			})
			
			if move {
				switch {
					case !downS && !leftS && !rightS: //break
						break
					case (downS && leftS && rightS) || (leftL && rightS) || (rightL && leftS): //nowhere
						move, moveL = false, false
						leftS, rightS, downS = false, false, false
						break
					case !leftL && !rightL && !leftS && !rightS && downS: //choose
						switch rand.Intn(2) {
							case 0 :
								Flake = Point{Flake.X + 1, Flake.Y + 1}
							case 1 :
								Flake = Point{Flake.X - 1, Flake.Y + 1}
						}
					case !leftS && !leftL: //left
						Flake = Point{Flake.X - 1, Flake.Y + 1}
					case !rightS && !rightL: //right
						Flake = Point{Flake.X + 1, Flake.Y + 1}
				}
			}
			
			sdl.Do(func() {
				renderer.SetDrawColor(125, 255, 240, 255)
				renderer.DrawPoint(Flake.X, Flake.Y)
				renderer.Present()
			})
		}
		
		for index := 0; index < len(l.A) && move; index++ {
			if Compare(l.A[index], l.B[index], Flake.X, Flake.Y, Flake.X, Flake.Y + 1) {
				for ;moveL; {
					rightL = Compare(l.A[index], l.B[index], Flake.X, Flake.Y, Flake.X + 1, Flake.Y + 1)
					leftL = Compare(l.A[index], l.B[index], Flake.X, Flake.Y, Flake.X - 1, Flake.Y + 1)
					for index := len(l.A) - 8; index < len(l.A); index++ {
						vertical = Compare(l.A[index], l.B[index], Flake.X, Flake.Y, Flake.X + 1, Flake.Y + 1)
						if vertical {move, moveL = false, false; break}
					}
					
					for downS, leftS, rightS = Collision(Flake.X, Flake.Y); downS || leftS || rightS; {
						noObj = false
						rightL = Compare(l.A[index], l.B[index], Flake.X, Flake.Y, Flake.X + 1, Flake.Y + 1)
						leftL = Compare(l.A[index], l.B[index], Flake.X, Flake.Y, Flake.X - 1, Flake.Y + 1)
						for index := len(l.A) - 8; index < len(l.A); index++ {
							vertical = Compare(l.A[index], l.B[index], Flake.X, Flake.Y, Flake.X + 1, Flake.Y + 1)
							if vertical {move, moveL = false, false; break}
						}
						
						sdl.Do(func() {
							renderer.SetDrawColor(0, 0, 0, 255)
							renderer.DrawPoint(Flake.X, Flake.Y)
						})
						
						if move {
							switch {
								case !downS && !leftS && !rightS: //break
									break
								case (downS && leftS && rightS) || (leftL && rightS) || (rightL && leftS): //nowhere
									move, moveL = false, false
									leftS, rightS, downS = false, false, false
									break
								case !leftL && !rightL && !leftS && !rightS && downS: //choose
									switch rand.Intn(2) {
										case 0 :
											Flake = Point{Flake.X + 1, Flake.Y + 1}
										case 1 :
											Flake = Point{Flake.X - 1, Flake.Y + 1}
									}
								case !leftS && !leftL: //left
									Flake = Point{Flake.X - 1, Flake.Y + 1}
								case !rightS && !rightL: //right
									Flake = Point{Flake.X + 1, Flake.Y + 1}
							}
						}
						
						sdl.Do(func() {
							renderer.SetDrawColor(125, 255, 240, 255)
							renderer.DrawPoint(Flake.X, Flake.Y)
							renderer.Present()
						})
					}
					
					sdl.Do(func() {
						renderer.SetDrawColor(0, 0, 0, 255)
						renderer.DrawPoint(Flake.X, Flake.Y)
					})
					
					if move {
						switch {
							case !leftL && !rightL: //break
								moveL = false
								break
							case (downS && leftS && rightS) || (leftL && rightS) || (rightL && leftS): //nowhere
								move, moveL = false, false
								leftS, rightS, downS = false, false, false
								break
							case !leftL && !rightL && !leftS && !rightS && downS: //choose
								switch rand.Intn(2) {
									case 0 :
										Flake = Point{Flake.X + 1, Flake.Y + 1}
									case 1 :
										Flake = Point{Flake.X - 1, Flake.Y + 1}
								}
							case !leftS && !leftL: //left
								Flake = Point{Flake.X - 1, Flake.Y + 1}
							case !rightS && !rightL: //right
								Flake = Point{Flake.X + 1, Flake.Y + 1}
						}
					}
					
					sdl.Do(func() {
						renderer.SetDrawColor(125, 255, 240, 255)
						renderer.DrawPoint(Flake.X, Flake.Y)
						renderer.Present()
					})
				}
			}
		}
	}
	
	sdl.Do(func() {
		renderer.SetDrawColor(125, 255, 240, 255)
		renderer.DrawPoint(Flake.X, Flake.Y)
		renderer.Present()
	})
	snow = append(snow, Flake)
}

func Compare(A, B Point, X, Y, X2, Y2 int32) bool {
	var (
		Vx, Vy, Ux, Uy int32
		t, T float64
	)
	Vx = X2 - X
	Vy = Y2 - Y
	Ux = B.X - A.X
	Uy = B.Y - A.Y
	t = float64(((A.X - X) * Uy - (A.Y - Y) * Ux)) / float64((Vx * Uy - Vy * Ux))
	T = float64(((A.X - X) * Vy - (A.Y - Y) * Vx)) / float64((Vx * Uy - Vy * Ux))
	if t < 0 || t > 1 || T < 0 || T > 1 {return false}
	return true
}

func Collision(X, Y int32) (bool, bool, bool) {
	var down, left, right bool = false, false, false
	for i := 0; i < len(snow); i++ {
		if X == 0 || X - 1 == snow[i].X && Y + 1 == snow[i].Y {left = true}
		if X == winWidth - 1 || X + 1 == snow[i].X && Y + 1 == snow[i].Y {right = true}
		if X == snow[i].X && Y + 1 == snow[i].Y {down = true}
	}
	return down, left, right
}

func main() {
	window, _ = sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, winWidth, winHeight, sdl.WINDOW_SHOWN)
	renderer, _ = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	surface, _ = window.GetSurface()
	rand.Seed(time.Now().UnixNano())
	fmt.Print()
	l := DrawObjects()
	renderer.Present()
	sdl.Main(func() {Process(l)})
}
