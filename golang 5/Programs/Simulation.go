package main

import (
	"fmt"
	"sync"
	"time"
	"math"
	"math/rand"
	"github.com/veandco/go-sdl2/sdl"
)

type (
	numSpec [8]Entity
	Circle struct {
		O Point
		radius int32
	}
	Point struct {
		X, Y int32
	}
	Entity struct {
		num, L, R, U, D int
		r, g, b uint8
		susp bool
		energy []int
		pos []Point
	}
)

var (
	window *sdl.Window
	renderer *sdl.Renderer
	winTitle string = "Simulation"
	winWidth, winHeight int32 = 1920, 1080
	takenFood []Point
	nS numSpec
	lives int
	wgmain, wg, wgthreads sync.WaitGroup
)

func SafeDelete(id, w int) {
	copy(nS[id].pos[w:], nS[id].pos[w + 1:])
	nS[id].pos = nS[id].pos[:len(nS[id].pos) - 1]
	
	copy(nS[id].energy[w:], nS[id].energy[w + 1:])
	nS[id].energy = nS[id].energy[:len(nS[id].energy) - 1]
}

func SafeDeleteFood(xF, yF int32) {
	for n, k := range(takenFood) {
		if k.X == xF && k.Y == yF {
			copy(takenFood[n:], takenFood[n + 1:])
			takenFood = takenFood[:len(takenFood) - 1]
		}
	}
}

func DrawGrid() {
	var grid sdl.Rect
	
	for x, y := 15, 0; x <= 1905 && y <= 1080; {
		if x == 1905 {
			x = 15; y += 270
		} else {
			for xs, ys := 15, 0; xs <= x + 270 && ys <= y + 270; {
				if xs == x + 270 {xs = 15; ys += 30} else {
					grid = sdl.Rect{int32(xs), int32(ys), 30, 30}
					renderer.DrawRect(&grid)
					xs += 30
				}
			}
			x += 270
		}
	}
	renderer.Present()
}

func DrawCell(x, y int32, r, g, b uint8) {
	var rect sdl.Rect
	x *= 30
	x += 20
	y *= 30
	y += 5
	sdl.Do(func() {
		renderer.SetDrawColor(r, g, b, 255)
		rect = sdl.Rect{x, y, 20, 20}
		renderer.FillRect(&rect)
		renderer.Present()
	})
}

func (c Circle) DrawCircles() {
	for rad := 1; rad < 8; rad++ {
		newCircle(c.O, int32(rad)).DrawFood()
	}
}

func newCircle(c Point, r int32) Circle {
	return Circle{c, r}
}

func (c Circle) DrawFood() {
	c.O.X += 60
	c.O.Y += 45
	sdl.Do(func() {renderer.SetDrawColor(225, 130, 20, 255)})
	
	p := make([]Point, 0)
	x, y := float64(c.radius), 0.0
	fi := 1.0/float64(c.radius)
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
	sdl.Do(func() {renderer.Present()})
	return 
}

func AddFood() {
	var (
		xF, yF int32
		c Circle
	)
	for n := 0; n < 10; n++ {
		xF = int32(rand.Int31n(61))
		yF = int32(rand.Int31n(34))
		for CollisionFood(xF, yF) {
			xF = int32(rand.Int31n(61))
			yF = int32(rand.Int31n(34))
		}
		takenFood = append(takenFood, Point{xF, yF})
		c.O = Point{xF * 30, yF * 30}
		c.DrawCircles()
	}
}

func CheckForFood(xF, yF int32) (bool) {
	for _, k := range(takenFood) {
		if k.X == xF && k.Y == yF {
			return true
		}
	}
	return false
}

func StartSim() { //xMax 30 * 60 yMax 30 * 33
	var (
		r, g, b, r2, g2, b2 uint8
		add, genL, genR, genU, genD, total int
		x, y, x2, y2 int32
	)
	
	wgmain.Add(1)
	r2, g2, b2, x2, y2 = 255, 255, 255, -1, -1
	for i := 0; i < len(nS); i++ {
		total, genL ,genR, genU, genD = 100, 0, 0, 0, 0
		r = uint8(rand.Intn(256))
		g = uint8(rand.Intn(256))
		b = uint8(rand.Intn(256))
		x = int32(rand.Int31n(61))
		y = int32(rand.Int31n(34))
		
		for (r == r2 && g == g2 && b == b2) || (r == 255 && g == 255 && b == 255) || (x == x2 && y == y2) {
			r = uint8(rand.Intn(256))
			g = uint8(rand.Intn(256))
			b = uint8(rand.Intn(256))
			x = int32(rand.Int31n(61))
			y = int32(rand.Int31n(34))
		}
		r2, g2, b2, x2, y2 = r, g, b, x, y
		
		for total > 0 {
			add = rand.Intn(4) + 1
			if add > total {add = total}
			switch rand.Intn(4) {
				case 0:
					genL += add
				case 1:
					genR += add
				case 2:
					genU += add
				case 3:
					genD += add
			}
			total -= add
		}
		
		nS[i] = Entity{1, genL ,genR, genU, genD, r, g, b, false, nil, nil}
		nS[i].pos = append(nS[i].pos, Point{x, y})
		nS[i].energy = append(nS[i].energy, 50)
		lives++
		wg.Add(1)
		go Organism(i)
		wgthreads.Add(1)
	}
	wgmain.Done()
	
	for {
		wg.Wait()
		if lives == 0 {
			break
		} else {
			if len(takenFood) < 200 {AddFood()}
			for i := 0; i < len(nS); i++ {
				if nS[i].num > 0 {
					wg.Add(1)
					wgmain.Done()
				}
			}
		}
	}
}

func Organism(id int) {
	var (
		res int
		xN, yN int32 //x/y New | x/y Food
	)
	fmt.Printf("START - Entity - {%d, %d}\nchance going left - %d%%\nchance going right - %d%%\nchance going up - %d%%\nchance going down - %d%%\nenergy - %d\nr, g, b - %d, %d, %d\nposition - %v\n\n", id, nS[id].num, nS[id].L, nS[id].R, nS[id].U, nS[id].D, nS[id].energy[0], nS[id].r, nS[id].g, nS[id].b, nS[id].pos)
	
	for {
		wgthreads.Done()
		wgmain.Wait()
		wgthreads.Wait()
		if rand.Intn(10) == 0 {
			switch rand.Intn(4) {
				case 0:
					nS[id].L += 5
				case 1:
					nS[id].R += 5
				case 2:
					nS[id].U += 5
				case 3:
					nS[id].D += 5
			}
			switch rand.Intn(4) {
				case 0:
					nS[id].L -= 5
				case 1:
					nS[id].R -= 5
				case 2:
					nS[id].U -= 5
				case 3:
					nS[id].D -= 5
			}
			fmt.Printf("\nEntity - %d GENES UPDATED\nchance going left - %d%%\nchance going right - %d%%\nchance going up - %d%%\nchance going down - %d%%\n\n", id, nS[id].L, nS[id].R, nS[id].U, nS[id].D)
		}
		
		for w := 0; w < nS[id].num; w++ {
			//Добавить запись
			if nS[id].energy[w] == 0 {
				DrawCell(nS[id].pos[w].X, nS[id].pos[w].Y, 255, 255, 255)
				SafeDelete(id, w)
				lives--
				nS[id].num--
				fmt.Printf("Entity - {%d, %d} DIED (%d left; %d total lives)\n", id, w, nS[id].num, lives)
			} else {
				res = rand.Intn(100) + 1
				if res <= nS[id].L && nS[id].pos[w].X > 0 && !Collision(nS[id].pos[w].X - 1, nS[id].pos[w].Y) { //Left
					xN = nS[id].pos[w].X - 1
					yN = nS[id].pos[w].Y
				} else if res > nS[id].L && res <= (nS[id].L + nS[id].R) && nS[id].pos[w].X < 60 && !Collision(nS[id].pos[w].X + 1, nS[id].pos[w].Y) { //Right
					xN = nS[id].pos[w].X + 1
					yN = nS[id].pos[w].Y
				} else if res > (nS[id].L + nS[id].R) && res <= (nS[id].L + nS[id].R + nS[id].U) && nS[id].pos[w].Y > 0 && !Collision(nS[id].pos[w].X, nS[id].pos[w].Y - 1) { //Up
					xN = nS[id].pos[w].X
					yN = nS[id].pos[w].Y - 1
				} else if res > (nS[id].L + nS[id].R + nS[id].U) && res <= (nS[id].L + nS[id].R + nS[id].U + nS[id].D) && nS[id].pos[w].Y < 33 && !Collision(nS[id].pos[w].X, nS[id].pos[w].Y + 1) { //Down
					xN = nS[id].pos[w].X
					yN = nS[id].pos[w].Y + 1
				}
				
				if CheckForFood(xN, yN) {
					SafeDeleteFood(xN, yN)
					nS[id].energy[w] += 10
					fmt.Printf("Entity - {%d, %d} FOOD CONSUMED (total energy - %d)\n", id, w, nS[id].energy[w])
				}
				
				DrawCell(nS[id].pos[w].X, nS[id].pos[w].Y, 255, 255, 255)
				nS[id].pos[w] = Point{xN, yN}
				DrawCell(nS[id].pos[w].X, nS[id].pos[w].Y, nS[id].r, nS[id].g, nS[id].b)
				nS[id].energy[w]--
			}
		}
		
		if nS[id].num > 0 {
			if nS[id].energy[0] > 8 && nS[id].num < 20 && rand.Intn(4) == 0 {
				if nS[id].pos[0].X > 0 && !Collision(nS[id].pos[0].X - 1, nS[id].pos[0].Y) {
					nS[id].pos = append(nS[id].pos, Point{nS[id].pos[0].X - 1, nS[id].pos[0].Y})
					nS[id].energy = append(nS[id].energy, 50)
					fmt.Printf("Entity - %d ADDED (%d total; %d total lives)\n", id, nS[id].num + 1, lives + 1)
					lives++
				} else if nS[id].pos[0].X < 61 && !Collision(nS[id].pos[0].X + 1, nS[id].pos[0].Y) {
					nS[id].pos = append(nS[id].pos, Point{nS[id].pos[0].X + 1, nS[id].pos[0].Y})
					nS[id].energy = append(nS[id].energy, 50)
					fmt.Printf("Entity - %d ADDED (%d total; %d total lives)\n", id, nS[id].num + 1, lives + 1)
					lives++
				} else if nS[id].pos[0].Y > 0 && !Collision(nS[id].pos[0].X, nS[id].pos[0].Y - 1) {
					nS[id].pos = append(nS[id].pos, Point{nS[id].pos[0].X, nS[id].pos[0].Y - 1})
					nS[id].energy = append(nS[id].energy, 50)
					fmt.Printf("Entity - %d ADDED (%d total; %d total lives)\n", id, nS[id].num + 1, lives + 1)
					lives++
				} else if nS[id].pos[0].Y < 34 && !Collision(nS[id].pos[0].X, nS[id].pos[0].Y + 1) {
					nS[id].pos = append(nS[id].pos, Point{nS[id].pos[0].X, nS[id].pos[0].Y + 1})
					nS[id].energy = append(nS[id].energy, 50)
					fmt.Printf("Entity - %d ADDED (%d total; %d total lives)\n", id, nS[id].num + 1, lives + 1)
					lives++
				} else {
					nS[id].energy[0] += 8
					nS[id].num--
				}
				nS[id].energy[0] -= 8
				nS[id].num++
			}
		}
		
		if nS[id].num != 0 {wgmain.Add(1); wgthreads.Add(1)}
		wg.Done()
		if nS[id].num == 0 {break}
	}
	fmt.Printf("%d Entity END %t - suspend\n", id, nS[id].susp)
}

func Collision(x, y int32) bool {
	for m, _ := range(nS) {
		if nS[m].num != 0 {
			for _, p := range(nS[m].pos) {
				if p.X == x && p.Y == y {return true}
			}
		}
	}
	return false
}

func CollisionFood(xF, yF int32) bool {
	for _, p := range(takenFood) {
		if p.X == xF && p.Y == yF {return true}
	}
	return false
}

func AllTrue() bool {
	for i := 0; i < len(nS); i++ {
		if !nS[i].susp {return false}
	}
	return true
}

func main() {
	window, _ = sdl.CreateWindow(winTitle, sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, winWidth, winHeight, sdl.WINDOW_SHOWN)
	renderer, _ = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	defer window.Destroy()
	defer renderer.Destroy()
	
	rand.Seed(time.Now().UnixNano())
	renderer.SetDrawColor(255, 255, 255, 255)
	renderer.Clear()
	renderer.SetDrawColor(0, 0, 100, 255)
	DrawGrid()
	sdl.Main(func() {StartSim()})
	
	renderer.SetDrawColor(160, 0, 0, 255)
	DrawGrid()
	time.Sleep(10 * time.Second)
}
