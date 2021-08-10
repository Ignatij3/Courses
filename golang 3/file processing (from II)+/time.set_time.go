package main

import (
	"os"
	"fmt"
	"time"
)

func main() {
	finfo, _ := os.Stat(os.Args[0])
	fmt.Println(finfo.Size(), finfo.ModTime())			// 1

	os.Mkdir("test.dir", 0777)
	finfo, _ = os.Stat("test.dir")
	fmt.Println(finfo.Size(), finfo.ModTime())			// 2
	
    newFile, err := os.Create("test.txt")
    if err != nil {
        fmt.Println(err)
		return
    }
    newFile.Close()
	finfo, _ = os.Stat("test.txt")
	fmt.Println(finfo.Size(), finfo.ModTime())			// 3
	mtime := time.Date(2006, time.February, 1, 21, 42, 5, 0, time.UTC)
	atime := time.Date(2007, time.March, 8, 14, 15, 0, 0, time.UTC)
	if err := os.Chtimes("test.txt", atime, mtime); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	finfo, _ = os.Stat("test.txt")
	fmt.Println(finfo.Size(), finfo.ModTime())			// 4

	loc := time.FixedZone("UTC+2", +2*60*60)
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, loc)
	fmt.Println("The time is:", t)						// 5
	t = time.Now()
	fmt.Println("The time is:", t)						// 6
	fmt.Println(t.Date())								// 7
	fmt.Println(t.Clock())								// 8
	fmt.Println(t.YearDay())							// 9
	fmt.Println(t.Before(time.Now()))					//10
}

