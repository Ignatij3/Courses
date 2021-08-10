package main

import Vse_Moi_Bratki_Usajut_FMT_Chekaite "fmt"

func main() {
	num := 186573
	mytext := "Text below is\n"
	var character rune = '*'
	Vse_Moi_Bratki_Usajut_FMT_Chekaite.Print(mytext)
	Vse_Moi_Bratki_Usajut_FMT_Chekaite.Println("HERE", num)
	Vse_Moi_Bratki_Usajut_FMT_Chekaite.Scan(&mytext)
	Vse_Moi_Bratki_Usajut_FMT_Chekaite.Println(mytext)
	Vse_Moi_Bratki_Usajut_FMT_Chekaite.Printf("Single character: %c\nType - %T", character, character)
}
