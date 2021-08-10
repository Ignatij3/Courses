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
	dir := 0
	color := termbox.ColorYellow
	defer termbox.Close()
	w, h := termbox.Size()
	L, U, R, D := 0, 0, w-1, h-1
	for X, Y := 0, 0; ; {
	termbox.SetCell(X, Y, ' ', termbox.ColorRed, color)	
	color++
		if color > 7 { color = 1 }
		if dir == 0 {
			if X < R { X++ } else { U+=2; Y++; dir++ }
		} else if dir == 1 {
			if X < U { Y++ } else {X--; dir++ }
		} else if dir == 2 {
			if X > L { X-- } else { U+=2; Y++; dir++ }
		} else if dir == 3 {
			if Y > U { X++ } else {Y++; dir = 0 }
		}
		if X == D && X == R {break}
		time.Sleep(20*time.Millisecond)
		termbox.Flush()
	}
	fmt.Println("Done")
}
