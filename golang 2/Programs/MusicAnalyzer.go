package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type  (
	tAlbum struct  {
		artist string
		genre string
		title string
		year string
		amount string
	}	
	tCollection []tAlbum
)

func (a *tAlbum) InitFromString(s string) {
	m := strings.Split(s, "[")
	for _, x := range m {    // x == 'artist]John Smith'
		p := strings.Split(x, "]")
		if len(p) == 2 {
			t := p[0]
			v := p[1]
			if t == "artist" {
				a.artist = v
			} else if t == "genre" {
				a.genre = v
			} else if t == "title" {
				a.title = v
			} else if t == "year" {
				a.year = v
			} else if t == "amount" {
				a.amount = v
			}
		}
	}
}

func main()  {
	var list tCollection
	var a tAlbum
	music, _ := os.Open("..\\Files\\music.txt")
	defer music.Close()
	mScan := bufio.NewScanner(music)
	for mScan.Scan() {
		s := mScan.Text()
		s = strings.ReplaceAll(s, "##", "[year]")
		s = strings.ReplaceAll(s, "&&&", "[amount]")
		a.InitFromString(s)
		list = append(list, a)
	}
	for i := 0; i < len(list); i++ {
		fmt.Printf("%v\n", list[i])
	}
}
