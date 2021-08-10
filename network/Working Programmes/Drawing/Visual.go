package main

import (
	"fmt"
	"image/color"
	"log"
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

type (
	rectangle struct {
		center        point
		width, height int32
	}
	circle struct {
		center point
		radius int32
	}
	triangle struct {
		p1, p2, p3 point
		size       int32
	}
	hexagon struct {
		center point
		size   int32
	}
	toolHitbox struct {
		upperLeftCorner point
		side            int32
	}
	point struct {
		x, y int32
	}
)

var (
	winWidth, winHeight int32 = 1280, 820
	toolsBar            int32 = 100 //100px height for tools
)

func makeWindow() (*sdl.Window, *sdl.Renderer) {
	window, err := sdl.CreateWindow("Paint", sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, winWidth, winHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		log.Fatal(err)
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		log.Fatal(err)
	}

	renderer.SetDrawColor(255, 255, 255, 255)
	renderer.Clear() //Заполняю экран белым цветом
	return window, renderer
}

func changeFrame(renderer *sdl.Renderer, hitbox [14]toolHitbox, n, lCT, lCC int) (int, int) {
	renderer.SetDrawColor(0, 0, 0, 255)
	switch n {
	case 0, 1, 2, 3: //Если это одна из фигур
		renderer.DrawRect(&sdl.Rect{hitbox[lCT].upperLeftCorner.x - 1, hitbox[lCT].upperLeftCorner.y - 1, hitbox[lCT].side + 2, hitbox[lCT].side + 2}) //Закрашиываю рамку чёрным цветом
		lCT = n                                                                                                                                        //Обновляю информацию о последней выбранной фигуре
	default: //Если это цвет
		renderer.DrawRect(&sdl.Rect{hitbox[lCC].upperLeftCorner.x - 1, hitbox[lCC].upperLeftCorner.y - 1, hitbox[lCC].side + 2, hitbox[lCC].side + 2}) //Закрашиываю рамку чёрным цветом
		lCC = n                                                                                                                                        //Обновляю информацию о последнем выбранным цвете
	}

	renderer.SetDrawColor(255, 255, 0, 255)
	renderer.DrawRect(&sdl.Rect{hitbox[n].upperLeftCorner.x - 1, hitbox[n].upperLeftCorner.y - 1, hitbox[n].side + 2, hitbox[n].side + 2}) //Перерисовываю рамку выбранного инструмента жёлтым цветом
	renderer.Present()
	renderer.SetDrawColor(cl.R, cl.G, cl.B, cl.A) //Ставлю тот цвет, что выбрал пользователь, изначально это чёрный

	return lCT, lCC
}

func initHitbox() [14]toolHitbox {
	var hitbox [14]toolHitbox

	for n := 0; n < 14; n++ {
		hitbox[n] = toolHitbox{upperLeftCorner: point{int32(32/3 + n*32/3 + n*80), 10}, side: 80} //Мои любимые магические константы
		//Они родились вследствие расчётов расположения рамок с инструментами на равном расстоянии друг от друга и от краёв экрана
		//К сожалению, времени на написание программы было мало (3-4 дня), поэтому я сделал её немаштабируемой
	}
	return hitbox //Возвращает координаты всех квадратных рамок, для взаимодействия с ними
}

//Add Alpha 0
func drawOptions(renderer *sdl.Renderer, hitbox [14]toolHitbox, lCT, lCC int) {
	renderer.SetDrawColor(255, 255, 255, 255)
	renderer.FillRect(&sdl.Rect{0, 0, winWidth, toolsBar}) //Заполняю белым прямоугольник с инструментами (чтобы избавиться от частей фигур, рисуемых выше положенного)
	renderer.SetDrawColor(0, 0, 0, 255)
	renderer.DrawLine(0, toolsBar, winWidth, toolsBar) //Рисую линию, разделяющую инструменты и область рисования
	var toolFrClr, colorFrClr sdl.Color                //toolFrameColor и colocFrameColor

	for n := 0; n < 14; n++ {
		//Есть 2 вида "инстрментов" - фигура и цвет, рамка выбранной фигуры и выбранного цвета будут жёлтыми, остальные - чёрные
		if n == lCT { //lCT - lastCheckedTool. Последняя выбранная фигура (нынешняя)
			toolFrClr = sdl.Color{255, 255, 0, 255}
		} else {
			toolFrClr = sdl.Color{0, 0, 0, 255}
		}
		if n == lCC { //lCC - lastCheckedColor. Последний выбранный цвет (нынешний)
			colorFrClr = sdl.Color{255, 255, 0, 255}
		} else {
			colorFrClr = sdl.Color{0, 0, 0, 255}
		}

		switch n { //Далее расписаны рисуемые фигуры и цвета
		case 0:
			renderer.SetDrawColor(toolFrClr.R, toolFrClr.G, toolFrClr.B, 255)
			renderer.DrawRect(&sdl.Rect{hitbox[0].upperLeftCorner.x - 1, hitbox[0].upperLeftCorner.y - 1, hitbox[0].side + 2, hitbox[0].side + 2}) //frame

			renderer.SetDrawColor(0, 0, 0, 255)
			renderer.FillRect(&sdl.Rect{hitbox[0].upperLeftCorner.x, hitbox[0].upperLeftCorner.y, hitbox[0].side, hitbox[0].side}) //rectangle
		case 1:
			renderer.SetDrawColor(toolFrClr.R, toolFrClr.G, toolFrClr.B, 255)
			renderer.DrawRect(&sdl.Rect{hitbox[1].upperLeftCorner.x - 1, hitbox[1].upperLeftCorner.y - 1, hitbox[1].side + 2, hitbox[1].side + 2}) //frame

			renderer.SetDrawColor(0, 0, 0, 255)
			for rad := int32(40); rad > 0; rad-- {
				circle{center: point{hitbox[1].upperLeftCorner.x + 40, hitbox[1].upperLeftCorner.y + 40}, radius: rad}.drawCircle(renderer) //circle
			}
		case 2:
			renderer.SetDrawColor(toolFrClr.R, toolFrClr.G, toolFrClr.B, 255)
			renderer.DrawRect(&sdl.Rect{hitbox[2].upperLeftCorner.x - 1, hitbox[2].upperLeftCorner.y - 1, hitbox[2].side + 2, hitbox[2].side + 2}) //frame

			renderer.SetDrawColor(0, 0, 0, 255)
			triangle{p1: point{hitbox[2].upperLeftCorner.x + 40, hitbox[2].upperLeftCorner.y}, p2: point{hitbox[2].upperLeftCorner.x, hitbox[2].upperLeftCorner.y + 80},
				p3: point{hitbox[2].upperLeftCorner.x + 80, hitbox[2].upperLeftCorner.y + 80}, size: 0}.drawTriangle(renderer) //triangle
		case 3:
			renderer.SetDrawColor(toolFrClr.R, toolFrClr.G, toolFrClr.B, 255)
			renderer.DrawRect(&sdl.Rect{hitbox[3].upperLeftCorner.x - 1, hitbox[3].upperLeftCorner.y - 1, hitbox[3].side + 2, hitbox[3].side + 2}) //frame

			renderer.SetDrawColor(0, 0, 0, 255)
			hexagon{center: point{hitbox[3].upperLeftCorner.x, hitbox[3].upperLeftCorner.y}, size: 40}.drawHexagon(renderer) //hexagon
		case 4:
			renderer.SetDrawColor(colorFrClr.R, colorFrClr.G, colorFrClr.B, 255)
			renderer.DrawRect(&sdl.Rect{hitbox[4].upperLeftCorner.x - 1, hitbox[4].upperLeftCorner.y - 1, hitbox[4].side + 2, hitbox[4].side + 2}) //frame

			renderer.SetDrawColor(255, 0, 0, 255) //red
			renderer.FillRect(&sdl.Rect{hitbox[4].upperLeftCorner.x, hitbox[4].upperLeftCorner.y, hitbox[4].side, hitbox[4].side})
		case 5:
			renderer.SetDrawColor(colorFrClr.R, colorFrClr.G, colorFrClr.B, 255)
			renderer.DrawRect(&sdl.Rect{hitbox[5].upperLeftCorner.x - 1, hitbox[5].upperLeftCorner.y - 1, hitbox[5].side + 2, hitbox[5].side + 2}) //frame

			renderer.SetDrawColor(255, 145, 0, 255) //orange
			renderer.FillRect(&sdl.Rect{hitbox[5].upperLeftCorner.x, hitbox[5].upperLeftCorner.y, hitbox[5].side, hitbox[5].side})
		case 6:
			renderer.SetDrawColor(colorFrClr.R, colorFrClr.G, colorFrClr.B, 255)
			renderer.DrawRect(&sdl.Rect{hitbox[6].upperLeftCorner.x - 1, hitbox[6].upperLeftCorner.y - 1, hitbox[6].side + 2, hitbox[6].side + 2}) //frame

			renderer.SetDrawColor(255, 255, 0, 255) //yellow
			renderer.FillRect(&sdl.Rect{hitbox[6].upperLeftCorner.x, hitbox[6].upperLeftCorner.y, hitbox[6].side, hitbox[6].side})
		case 7:
			renderer.SetDrawColor(colorFrClr.R, colorFrClr.G, colorFrClr.B, 255)
			renderer.DrawRect(&sdl.Rect{hitbox[7].upperLeftCorner.x - 1, hitbox[7].upperLeftCorner.y - 1, hitbox[7].side + 2, hitbox[7].side + 2}) //frame

			renderer.SetDrawColor(0, 255, 0, 255) //green
			renderer.FillRect(&sdl.Rect{hitbox[7].upperLeftCorner.x, hitbox[7].upperLeftCorner.y, hitbox[7].side, hitbox[7].side})
		case 8:
			renderer.SetDrawColor(colorFrClr.R, colorFrClr.G, colorFrClr.B, 255)
			renderer.DrawRect(&sdl.Rect{hitbox[8].upperLeftCorner.x - 1, hitbox[8].upperLeftCorner.y - 1, hitbox[8].side + 2, hitbox[8].side + 2}) //frame

			renderer.SetDrawColor(0, 255, 255, 255) //cyan
			renderer.FillRect(&sdl.Rect{hitbox[8].upperLeftCorner.x, hitbox[8].upperLeftCorner.y, hitbox[8].side, hitbox[8].side})
		case 9:
			renderer.SetDrawColor(colorFrClr.R, colorFrClr.G, colorFrClr.B, 255)
			renderer.DrawRect(&sdl.Rect{hitbox[9].upperLeftCorner.x - 1, hitbox[9].upperLeftCorner.y - 1, hitbox[9].side + 2, hitbox[9].side + 2}) //frame

			renderer.SetDrawColor(0, 0, 255, 255) //blue
			renderer.FillRect(&sdl.Rect{hitbox[9].upperLeftCorner.x, hitbox[9].upperLeftCorner.y, hitbox[9].side, hitbox[9].side})
		case 10:
			renderer.SetDrawColor(colorFrClr.R, colorFrClr.G, colorFrClr.B, 255)
			renderer.DrawRect(&sdl.Rect{hitbox[10].upperLeftCorner.x - 1, hitbox[10].upperLeftCorner.y - 1, hitbox[10].side + 2, hitbox[10].side + 2}) //frame

			renderer.SetDrawColor(150, 0, 255, 255) //purple
			renderer.FillRect(&sdl.Rect{hitbox[10].upperLeftCorner.x, hitbox[10].upperLeftCorner.y, hitbox[10].side, hitbox[10].side})
		case 11:
			renderer.SetDrawColor(colorFrClr.R, colorFrClr.G, colorFrClr.B, 255)
			renderer.DrawRect(&sdl.Rect{hitbox[11].upperLeftCorner.x - 1, hitbox[11].upperLeftCorner.y - 1, hitbox[11].side + 2, hitbox[11].side + 2}) //frame

			renderer.SetDrawColor(255, 255, 255, 255) //white
			renderer.FillRect(&sdl.Rect{hitbox[11].upperLeftCorner.x, hitbox[11].upperLeftCorner.y, hitbox[11].side, hitbox[11].side})
		case 12:
			renderer.SetDrawColor(colorFrClr.R, colorFrClr.G, colorFrClr.B, 255)
			renderer.DrawRect(&sdl.Rect{hitbox[12].upperLeftCorner.x - 1, hitbox[12].upperLeftCorner.y - 1, hitbox[12].side + 2, hitbox[12].side + 2}) //frame

			renderer.SetDrawColor(60, 30, 0, 255) //brown
			renderer.FillRect(&sdl.Rect{hitbox[12].upperLeftCorner.x, hitbox[12].upperLeftCorner.y, hitbox[12].side, hitbox[12].side})
		case 13:
			renderer.SetDrawColor(colorFrClr.R, colorFrClr.G, colorFrClr.B, 255)
			renderer.DrawRect(&sdl.Rect{hitbox[13].upperLeftCorner.x - 1, hitbox[13].upperLeftCorner.y - 1, hitbox[13].side + 2, hitbox[13].side + 2}) //frame

			renderer.SetDrawColor(0, 0, 0, 255) //black
			renderer.FillRect(&sdl.Rect{hitbox[13].upperLeftCorner.x, hitbox[13].upperLeftCorner.y, hitbox[13].side, hitbox[13].side})
		}
	}
}

//Эта функция рисует фигуру между нынешним и прошлым положением мыши, используя адаптированный алгоритм Брезенхэма, таким образом убирает разрывы
func drawBetween(object interface{}, renderer *sdl.Renderer, hitbox [14]toolHitbox, x1, y1, x2, y2 int32, lCT, lCC int) { //Функция взята с вики и изменена
	if x1 == 0 && y1 == 0 { //Если это первые координаты, то функция заканчивает своё выполнкник
		return
	}

	if math.Abs(float64(y2-y1)) < math.Abs(float64(x2-x1)) {
		if x1 > x2 {
			drawShapeLow(object, renderer, hitbox, x2, y2, x1, y1, lCT, lCC)
		} else {
			drawShapeLow(object, renderer, hitbox, x1, y1, x2, y2, lCT, lCC)
		}
	} else {
		if y1 > y2 {
			drawShapeHigh(object, renderer, hitbox, x2, y2, x1, y1, lCT, lCC)
		} else {
			drawShapeHigh(object, renderer, hitbox, x1, y1, x2, y2, lCT, lCC)
		}
	}
}

func drawShapeLow(object interface{}, renderer *sdl.Renderer, hitbox [14]toolHitbox, x1, y1, x2, y2 int32, lCT, lCC int) { //Функция взята с вики и изменена
	dx := x2 - x1
	dy := y2 - y1
	yi := int32(1)
	if dy < 0 {
		yi = -1
		dy = -dy
	}
	D := (2 * dy) - dx

	for ; x1 < x2; x1++ {
		drawShape(object, renderer, hitbox, x1, y1, lCT, lCC)
		if D > 0 {
			y1 += yi
			D += 2 * -dx
		}
		D += 2 * dy
	}
}

func drawShapeHigh(object interface{}, renderer *sdl.Renderer, hitbox [14]toolHitbox, x1, y1, x2, y2 int32, lCT, lCC int) { //Функция взята с вики и изменена
	dx := x2 - x1
	dy := y2 - y1
	xi := int32(1)
	if dx < 0 {
		xi = -1
		dx = -dx
	}
	D := (2 * dx) - dy

	for ; y1 < y2; y1++ {
		drawShape(object, renderer, hitbox, x1, y1, lCT, lCC)
		if D > 0 {
			x1 += xi
			D += 2 * -dy
		}
		D += 2 * dx
	}
}

func drawShape(object interface{}, renderer *sdl.Renderer, hitbox [14]toolHitbox, x, y int32, lCT, lCC int) (int32, int32) {
	switch sh := object.(type) { //В этом switch'e сначала обновляется информация об объекте, затем рисуется сам объект
	case rectangle:
		sh.center.x = x
		sh.center.y = y
		sh.drawRectangle(renderer)

		if sh.center.y-(sh.height/2) <= 100 { //Если низшая точка объекта находится на, или ниже 100 пикселей, обновляется панель с инструментами
			drawOptions(renderer, hitbox, lCT, lCC)
			renderer.SetDrawColor(cl.R, cl.G, cl.B, cl.A) //Ставлю обратно цвет, так как в процессе обновления тулбара он изменился
		}
	case circle:
		sh.center.x = x
		sh.center.y = y
		backupRad := sh.radius
		for rad := sh.radius; rad > 0; rad-- {
			sh.radius = rad
			sh.drawCircle(renderer)
		}

		if sh.center.y-backupRad <= 100 {
			drawOptions(renderer, hitbox, lCT, lCC)
			renderer.SetDrawColor(cl.R, cl.G, cl.B, cl.A)
		}
	case triangle:
		sh.p1.x, sh.p1.y = x, y-sh.size //size - это расстояние до угла треугольника
		sh.p2.x, sh.p2.y = x-sh.size, y+sh.size
		sh.p3.x, sh.p3.y = x+sh.size, y+sh.size
		sh.drawTriangle(renderer)

		if sh.p1.y-(sh.size/2) <= 100 {
			drawOptions(renderer, hitbox, lCT, lCC)
			renderer.SetDrawColor(cl.R, cl.G, cl.B, cl.A)
		}
	case hexagon:
		sh.center.x = x - sh.size //Перемещаю центр в то место, куда кликнула мышь
		sh.center.y = y - sh.size
		sh.drawHexagon(renderer)

		if sh.center.y-(sh.size/2) <= 100 {
			drawOptions(renderer, hitbox, lCT, lCC)
			renderer.SetDrawColor(cl.R, cl.G, cl.B, cl.A)
		}
	default: //Если ничего не выбрано (в самом начале исполнении программы)
		fmt.Println("Shape is not chosen, unable to draw")
	}
	return x, y
}

func (r rectangle) drawRectangle(renderer *sdl.Renderer) {
	renderer.FillRect(&sdl.Rect{r.center.x - (r.width / 2), r.center.y - (r.height / 2), r.width, r.height})

	var x1, x2, y1, y2 int32 = r.center.x - (r.width / 2), r.center.x + (r.width / 2), r.center.y - (r.height / 2), r.center.y + (r.height / 2)
	for y := y1; y <= y2; y++ {
		drawImageLine(x1, y, x2, y) //Рисую прямоугольник в *image.RGBA
	}
}

func (c circle) drawCircle(renderer *sdl.Renderer) { //Функция Сергея Мельника, с небольшими изменениями
	p := make([]point, 0)
	x, y := float64(c.radius), 0.0
	fi := 1.0 / float64(c.radius)
	cos, sin := math.Cos(fi), math.Sin(fi)
	for x > y {
		p = append(p, point{int32(math.Round(x)), int32(math.Round(y))})
		x, y = x*cos-y*sin, x*sin+y*cos
	}

	pp := make([]sdl.Point, len(p))
	imgp := make([]point, len(p)) //Этот слайс нужен для того, чтобы рисовать круги на *image.RGBA
	for i, v := range p {
		pp[i] = sdl.Point{c.center.x + v.x, c.center.y + v.y}
		imgp[i] = point{c.center.x + v.x, c.center.y + v.y}
	}
	renderer.DrawLines(pp)
	drawLines(imgp)

	for i, v := range p {
		pp[i] = sdl.Point{c.center.x - v.x, c.center.y + v.y}
		imgp[i] = point{c.center.x - v.x, c.center.y + v.y}
	}
	renderer.DrawLines(pp)
	drawLines(imgp)

	for i, v := range p {
		pp[i] = sdl.Point{c.center.x - v.y, c.center.y + v.x}
		imgp[i] = point{c.center.x - v.y, c.center.y + v.x}
	}
	renderer.DrawLines(pp)
	drawLines(imgp)

	for i, v := range p {
		pp[i] = sdl.Point{c.center.x - v.y, c.center.y - v.x}
		imgp[i] = point{c.center.x - v.y, c.center.y - v.x}
	}
	renderer.DrawLines(pp)
	drawLines(imgp)

	for i, v := range p {
		pp[i] = sdl.Point{c.center.x - v.x, c.center.y - v.y}
		imgp[i] = point{c.center.x - v.x, c.center.y - v.y}
	}
	renderer.DrawLines(pp)
	drawLines(imgp)

	for i, v := range p {
		pp[i] = sdl.Point{c.center.x + v.x, c.center.y - v.y}
		imgp[i] = point{c.center.x + v.x, c.center.y - v.y}
	}
	renderer.DrawLines(pp)
	drawLines(imgp)

	for i, v := range p {
		pp[i] = sdl.Point{c.center.x + v.y, c.center.y - v.x}
		imgp[i] = point{c.center.x + v.y, c.center.y - v.x}
	}
	renderer.DrawLines(pp)
	drawLines(imgp)

	for i, v := range p {
		pp[i] = sdl.Point{c.center.x + v.y, c.center.y + v.x}
		imgp[i] = point{c.center.x + v.y, c.center.y + v.x}
	}
	renderer.DrawLines(pp)
	drawLines(imgp)
}

func (t *triangle) SortPoints() { //Сортирую по y координате
	if (*t).p1.y > (*t).p2.y {
		(*t).p1, (*t).p2 = (*t).p2, (*t).p1
	}
	if (*t).p2.y > (*t).p3.y {
		(*t).p2, (*t).p3 = (*t).p3, (*t).p2
	}
	if (*t).p1.y > (*t).p2.y {
		(*t).p1, (*t).p2 = (*t).p2, (*t).p1
	}
}

func (t triangle) drawTriangle(renderer *sdl.Renderer) { //Не моё
	t.SortPoints()

	if t.p2.y == t.p3.y {
		t.fillBottomFlatTriangle(renderer)
	} else if t.p1.y == t.p2.y {
		t.fillTopFlatTriangle(renderer)
	} else {
		p4 := point{x: t.p1.x + int32((float64(t.p2.y-t.p1.y)/float64(t.p3.y-t.p1.y)))*(t.p3.x-t.p1.x), y: t.p2.y}

		t.p3, p4 = p4, t.p3
		t.fillBottomFlatTriangle(renderer)
		t.p1, t.p2, t.p3 = t.p2, t.p3, p4 //В t.p3 находится 4-й пункт и наоборот
		t.fillTopFlatTriangle(renderer)
	}
}

func (t triangle) fillBottomFlatTriangle(renderer *sdl.Renderer) { //Не моё
	var (
		slope1 float64 = float64(t.p2.x-t.p1.x) / float64(t.p2.y-t.p1.y)
		slope2 float64 = float64(t.p3.x-t.p1.x) / float64(t.p3.y-t.p1.y)
		curx1  float64 = float64(t.p1.x)
		curx2  float64 = float64(t.p1.x)
	)

	for pointY := t.p1.y; pointY <= t.p2.y; pointY++ {
		renderer.DrawLine(int32(curx1), pointY, int32(curx2), pointY)
		drawImageLine(int32(curx1), pointY, int32(curx2), pointY)
		curx1 += slope1
		curx2 += slope2
	}
}

func (t triangle) fillTopFlatTriangle(renderer *sdl.Renderer) { //Не моё
	var (
		slope1 float64 = float64(t.p3.x-t.p1.x) / float64(t.p3.y-t.p1.y)
		slope2 float64 = float64(t.p3.x-t.p2.x) / float64(t.p3.y-t.p2.y)
		curx1  float64 = float64(t.p3.x)
		curx2  float64 = float64(t.p3.x)
	)

	for pointY := t.p3.y; pointY > t.p1.y; pointY-- {
		renderer.DrawLine(int32(curx1), pointY, int32(curx2), pointY)
		drawImageLine(int32(curx1), pointY, int32(curx2), pointY)
		curx1 -= slope1
		curx2 -= slope2
	}
}

func (hex hexagon) drawHexagon(renderer *sdl.Renderer) {
	var (
		x2, y2, x3, y3 int32
		tri            triangle
	)

	for angle := 60; angle <= 360; angle += 60 { //Получаю координаты углов и по ним рисую 6 треугольников, первая точка всегда центр
		x2 = int32(math.Cos(float64(angle-60)*(math.Pi/180.0))*float64(hex.size)) + hex.size + hex.center.x
		y2 = int32(math.Sin(float64(angle-60)*(math.Pi/180.0))*float64(hex.size)) + hex.size + hex.center.y
		x3 = int32(math.Cos(float64(angle)*(math.Pi/180.0))*float64(hex.size)) + hex.size + hex.center.x
		y3 = int32(math.Sin(float64(angle)*(math.Pi/180.0))*float64(hex.size)) + hex.size + hex.center.y
		tri = triangle{p1: point{hex.center.x + hex.size, hex.center.y + hex.size}, p2: point{x2, y2}, p3: point{x3, y3}, size: 0}
		tri.drawTriangle(renderer)
	}

}

func drawLines(p1 []point) { //Работает аналогично renderer.DrawLines() (Проверку этой функции можно найти по пути "Test files/Test.exe")
	for n := 1; n < len(p1); n++ {
		drawImageLine(p1[n].x, p1[n].y, p1[n-1].x, p1[n-1].y)
	}
}

func drawImageLineLow(x1, y1, x2, y2 int32) { //Функция взята с вики и изменена
	dx := x2 - x1
	dy := y2 - y1
	yi := int32(1)
	if dy < 0 {
		yi = -1
		dy = -dy
	}
	D := (2 * dy) - dx

	col := color.RGBA{cl.R, cl.G, cl.B, cl.A}
	for ; x1 < x2; x1++ {
		img.Set(int(x1), int(y1), col)
		if D > 0 {
			y1 += yi
			D += 2 * -dx
		}
		D += 2 * dy
	}
}

func drawImageLineHigh(x1, y1, x2, y2 int32) { //Функция взята с вики и изменена
	dx := x2 - x1
	dy := y2 - y1
	xi := int32(1)
	if dx < 0 {
		xi = -1
		dx = -dx
	}
	D := (2 * dx) - dy

	col := color.RGBA{cl.R, cl.G, cl.B, cl.A}
	for ; y1 < y2; y1++ {
		img.Set(int(x1), int(y1), col)
		if D > 0 {
			x1 += xi
			D += 2 * -dy
		}
		D += 2 * dx
	}
}

func drawImageLine(x1, y1, x2, y2 int32) { //Функция взята с вики и изменена
	y1 -= 100 //Отнимаю по 100, что-бы вместить объект в картинку (окно на 100 пикселей выше, чем картинка, так как в картинке панель с инструментами не сохраняется)
	y2 -= 100
	if y1 < 0 && y2 < 0 { //Если объект полностью находится на панели инструментов, он просто не учитывается
		return
	}

	if y1 < 0 { //Обрезаю объект, если он частично находится на панели инструнетов
		y1 = 0
	}
	if y2 < 0 {
		y2 = 0
	}

	if math.Abs(float64(y2-y1)) < math.Abs(float64(x2-x1)) {
		if x1 > x2 {
			drawImageLineLow(x2, y2, x1, y1)
		} else {
			drawImageLineLow(x1, y1, x2, y2)
		}
	} else {
		if y1 > y2 {
			drawImageLineHigh(x2, y2, x1, y1)
		} else {
			drawImageLineHigh(x1, y1, x2, y2)
		}
	}
}
