package main

import (
	"bufio"
	"fmt"
	"os"
)

type  (
	tAlbum struct  {
		artist string
		genre string
		title string
		year uint16
		amount byte
	}	
	tCollection []tAlbum
)

func (r *tAlbum) InitString(s string)  {
	// . . . 
	// . . . 
	// . . . 
}	
	
func (c *tCollection) Append(a tAlbum)  {
	*c = append(*c, a)
	return
}	

func (c tCollection) CountGenre(genre string) int  {
	var res int
	for a, _ := range c  {
		if a.genre == genre  { res++ }
	}	 	
	return res
}		

func main()  {

	fin, _ := os.Open("music.dat")
	defer fin.Close()

	// Create a new Scanner for the input
	scanner := bufio.NewScanner(fin)

	var  (
		al tAlbum
		collection tCollection
	)
	for scanner.Scan() {
		al.InitString(scanner.Text())
		collection.Append(al)
	}
	fmt.Println(collection.CountGenre("Blues")
}	
