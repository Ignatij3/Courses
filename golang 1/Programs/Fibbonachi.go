package main
import "fmt"
func main() {
var a, b, c int
	for a, b = 1, 1 ; a<=b; a, b = b, c {
		c=a+b
		fmt.Print(c," ")
		if c>100000 {
			break
		}
	}
}
