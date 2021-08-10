package main

import (
	"bytes"
	"flag"
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/gorilla/websocket"
	"github.com/veandco/go-sdl2/sdl"
)

type currentObj interface{} //Pass shape here

var (
	img *image.RGBA
	cl  sdl.Color
)

func establishConnection() *websocket.Conn {
	addr := flag.String("addr", "localhost:8080", "http service address")
	u := url.URL{Scheme: "ws", Host: *addr, Path: "/"}
	conn, resp, err := websocket.DefaultDialer.Dial(u.String(), nil)

	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if fmt.Sprint(string(body)) == "" { //Так как после апгрейда до вебсокета отправлять ответ нельзя, он будет пустой при успешном подключении
		fmt.Println("Connection established successfully")
	} else {
		fmt.Println(string(body))
	}
	return conn
}

func createImage() *image.RGBA {
	return image.NewRGBA(image.Rectangle{Min: image.Point{0, 0}, Max: image.Point{int(winWidth), int(winHeight - 100)}}) //winHeight - 100 потому-что первые 100 пикселей - "панель инструментов"
}

func checkToolOnClick(tools [14]toolHitbox, x, y int32) (bool, int) {
	for n, c := range tools {
		if x >= c.upperLeftCorner.x && x <= c.upperLeftCorner.x+80 &&
			y >= c.upperLeftCorner.y && y <= c.upperLeftCorner.y+80 {
			return true, n
		}
	}
	return false, 0
}

func changeTool(object interface{}, renderer *sdl.Renderer, n int, x, y int32) interface{} {
	var size int32

	switch n {
	case 0: //В первых 4-х кейсах меняется фигура, поэтому прошу пользователся ввести размер фигуры (в пикселях)
		var rect rectangle

		fmt.Print("Enter width: ")
		fmt.Scan(&size)
		rect.width = size

		fmt.Print("Enter height: ")
		fmt.Scan(&size)
		rect.height = size

		return rect
	case 1:
		var circ circle

		fmt.Print("Enter radius: ")
		fmt.Scan(&size)
		circ.radius = size

		return circ
	case 2:
		var tri triangle

		fmt.Print("Enter size: ")
		fmt.Scan(&size)
		tri.size = size

		return tri
	case 3:
		var hex hexagon

		fmt.Print("Enter size: ")
		fmt.Scan(&size)
		hex.size = size

		return hex
	case 4: //В остальных меняю цвет
		renderer.SetDrawColor(255, 0, 0, 255)
		cl = sdl.Color{255, 0, 0, 255}
	case 5:
		renderer.SetDrawColor(255, 145, 0, 255)
		cl = sdl.Color{255, 145, 0, 255}
	case 6:
		renderer.SetDrawColor(255, 255, 0, 255)
		cl = sdl.Color{255, 255, 0, 255}
	case 7:
		renderer.SetDrawColor(0, 255, 0, 255)
		cl = sdl.Color{0, 255, 0, 255}
	case 8:
		renderer.SetDrawColor(0, 255, 255, 255)
		cl = sdl.Color{0, 255, 255, 255}
	case 9:
		renderer.SetDrawColor(0, 0, 255, 255)
		cl = sdl.Color{0, 0, 255, 255}
	case 10:
		renderer.SetDrawColor(150, 0, 255, 255)
		cl = sdl.Color{150, 0, 255, 255}
	case 11:
		renderer.SetDrawColor(255, 255, 255, 255)
		cl = sdl.Color{255, 255, 255, 255}
	case 12:
		renderer.SetDrawColor(60, 30, 0, 255)
		cl = sdl.Color{60, 30, 0, 255}
	case 13:
		renderer.SetDrawColor(0, 0, 0, 255)
		cl = sdl.Color{0, 0, 0, 255}
	}
	return object
}

func waitForInput(renderer *sdl.Renderer, hitbox [14]toolHitbox, conn *websocket.Conn) {
	var (
		event                             sdl.Event
		object                            currentObj //Сюда я закидываю одну из возможных фигур
		lastCheckedTool, lastCheckedColor int
		lastX, lastY                      int32
		control                           bool //Нажат ли Ctrl или нет
	)

	for {
		event = sdl.WaitEvent()
		switch t := event.(type) {
		case *sdl.KeyboardEvent:
			if t.State == sdl.PRESSED {
				switch t.Keysym.Sym {
				case sdl.K_ESCAPE: //При нажатии escape программа закрывается
					return
				case sdl.K_LCTRL, sdl.K_RCTRL: //Нажатие Ctrl записывается, для реализации сочетания клавиш
					control = true
				case sdl.K_s:
					if control { //Чтобы произошло сохранение картинки Ctrl должен быть нажат
						sendImageToServer(conn)
						saveImageUser()
					}
				}
			} else { //Если пользователь отпустил какие-либо клавиши, считается, что Ctrl не нажат
				control = false
			}
		case *sdl.MouseButtonEvent:
			if t.State == sdl.PRESSED {
				switch t.Button {
				case sdl.BUTTON_LEFT: //ВАЖНО! Если мышью не двигать, то рисование объекта произойдёт после отпускания кнопки мыши, или при нажатии другой (линия 173)
					if t.Y <= 100 { //Если затронута панель инструментов
						if ok, n := checkToolOnClick(hitbox, t.X, t.Y); ok { //Если по координатам среди хитбоксов нашёлся тот, в пределах которого произошло нажатие, происходит выбор данного "инструмента"
							lastCheckedTool, lastCheckedColor = changeFrame(renderer, hitbox, n, lastCheckedTool, lastCheckedColor) //Меняю активную рамку
							object = changeTool(object, renderer, n, t.X, t.Y)                                                      //Меняю активный инструмент
						}
					} else { //В ином случае, происходит рисование
						for {
							event = sdl.WaitEvent()
							switch tp := event.(type) {
							case *sdl.MouseButtonEvent:
								drawShape(object, renderer, hitbox, tp.X, tp.Y, lastCheckedTool, lastCheckedColor)
								renderer.Present()
								if tp.Type == sdl.MOUSEBUTTONUP {
									goto end
								}
							case *sdl.MouseMotionEvent: //При движении мыши рисуется объект с новыми координатами
								drawBetween(object, renderer, hitbox, lastX, lastY, tp.X, tp.Y, lastCheckedTool, lastCheckedColor) //Передаю прошлое и нынешнее положение мыши
								lastX, lastY = drawShape(object, renderer, hitbox, tp.X, tp.Y, lastCheckedTool, lastCheckedColor)
								renderer.Present()
							}
						}
					end:
						lastX, lastY = 0, 0 //Обновляю переменные
					}
				}
			}
		}
	}
}

func saveImageUser() {
	var (
		fName string
		b     strings.Builder
	)

	fmt.Print("Enter file name: ")
	fmt.Scan(&fName)

	dirname, err := os.UserHomeDir() //Картинка сохраняется в папку "%USERPROFILE/%Downloads", для этого я достаю путь профиля пользователя
	if err != nil {
		log.Fatal(err)
	}

	b.Grow(50)
	b.WriteString(dirname)
	b.WriteString("/Downloads/")
	b.WriteString(fName)
	b.WriteString(".png")

	file, err := os.Create(b.String())
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}
	png.Encode(file, img) //Эта функция кодирует имеющуюся картинку в file, в формате .png
}

func sendImageToServer(conn *websocket.Conn) {
	buf := new(bytes.Buffer)
	err := png.Encode(buf, img)
	if err != nil {
		log.Fatal(err)
	}
	byteImage := buf.Bytes() //[]byte нужен для передачи на сервер

	conn.WriteMessage(websocket.TextMessage, byteImage)
}

func main() {
	fmt.Println("Waiting for server")
	conn := establishConnection() //Захожу сюда, что-бы установить связь с сервером через websocket
	defer conn.Close()

	//ВНИМАНИЕ! Все функции, связанные с выводом на экран и обработкой изображения находятся в файле Visual.go
	window, renderer := makeWindow() //Создаю окно через sdl
	defer window.Destroy()
	defer renderer.Destroy()

	img = createImage()                   //Инициализирую полотно картинки, что-бы в будущем наносить на него изменения и сохранять
	hitbox := initHitbox()                //Первые 100 пикселей находится панель инструментов, а hitbox это слайс с квадратными рамками фигур и цветов
	drawOptions(renderer, hitbox, -1, -1) //Рисую те самые "инструменты"
	renderer.Present()
	waitForInput(renderer, hitbox, conn) //Главная функция, в ней я ожидаю ввода со стороны пользователя
}
