package main
import "fmt"
func main() {
	var int string
	fmt.Scan(&int)
	
	if int == "hi" {
		fmt.Println("hello!")
	} else {
		fmt.Println("bye!")
	}
}
