package main
import (
	"github.com/nsf/termbox-go"
	"time"
	"os"
	"fmt"
	)
func main() {
	err := termbox.Init()
	if err != nil {
		os.Exit(1)
	}
	defer termbox.Close()
	w, h := termbox.Size()
	w2 := w / 2
	LEFT, UP, RIGHT, DOWN := 0, 0, w2 - 1, h - 1
	dir := 0 // 0 - right, 1 - down, 2 - left, 3 - up
	color := termbox.ColorRed
	for x1, y1 := 0, 0; ; {
		termbox.SetCell(x1, y1, ' ', termbox.ColorMagenta, color)
		color++
		if color > 7 { color = 1 }
		x2 := w - x1 - 1
		y2 := y1
		termbox.SetCell(x2, y2, ' ', termbox.ColorMagenta, color)
		if dir == 0 {
			if x1 < RIGHT { x1++ } else { UP++; y1++; dir++ }
		} else if dir == 1 {
			if y1 < DOWN { y1++ } else { RIGHT--; x1--; dir++ }
		} else if dir == 2 {
			if x1 > LEFT { x1-- } else { DOWN--; y1--; dir++ }
		} else if dir == 3 {
			if y1 > UP { y1-- } else { LEFT++; x1++; dir = 0 }
		}
		if UP == DOWN {
			break
		}
		time.Sleep(40*time.Millisecond)
		termbox.Flush()
		fmt.Println("")
	}
	fmt.Println("DONE")
}
