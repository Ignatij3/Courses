package main
import "fmt"
func main() {
	var n int
	fmt.Scan(&n)
	if n%2==0 || n%3==0 || n%5==0 || n%7==0 || n%9==0 {
		fmt.Println("Это число НЕ простое")
		
	} else {
		fmt.Println("Это число простое")
		}
}
