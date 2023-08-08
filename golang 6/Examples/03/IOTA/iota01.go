package main

import "fmt"

func main() {
	// Basic sample
	const (
		Start = iota
		Run = iota
		End = iota
	)
	fmt.Printf("Start: %d\nRun: %d\nEnd: %d\n", Start, Run, End)
	fmt.Println()
/*
Start: 0
Run: 1
End: 2
*/ 

	// Simplified notation.
	//     Every time keyword const appears in the code, iota’s value is reseting
	const (
		Small = iota
		Medium
		Large
	)
	fmt.Printf("Small: %d %T\nMedium: %d %T\nLarge: %d %T\n", Small, Small, Medium, Medium, Large, Large)
	fmt.Println()
/*
Small: 0 int
Medium: 1 int
Large: 2 int
*/

const (
	Black = iota
	Blue
	Green
	Cyan
	Red
	Magenta
	Brown
	Gray
)
	fmt.Printf("Colors:\n Black = %d %T\n Blue = %d %T\n Green = %d %T\n Cyan = %d %T\n" +
	           " Red = %d %T\n Magenta = %d %T\n Brown = %d %T\n Gray = %d %T\n", 
		        Black, Black, Blue, Blue, Green, Green, Cyan, Cyan, 
		        Red, Red, Magenta, Magenta, Brown, Brown, Gray, Gray)
	fmt.Println()
/*
Colors:
 Black = 0 int
 Blue = 1 int
 Green = 2 int
 Cyan = 3 int 
 Red = 4 int
 Magenta = 5 int
 Brown = 6 int
 Gray = 7 int
*/

// Best practice: complete enumerated type (enum) with strings.
// See function ColorName definition below.
	for c:= 0; c < 8; c++ {
		fmt.Printf("%d %T %s\n", c, c, ColorName(c) )	
	}	
	fmt.Println()
/*
0 int Black
1 int Blue
2 int Green
3 int Cyan
4 int Red
5 int Magenta
6 int Brown
7 int Gray
*/

}

func ColorName(c int) string {
	names:= [...]string{"Black", "Blue", "Green", "Cyan", "Red", "Magenta", "Brown", "Gray"}
	return names[c]
}
