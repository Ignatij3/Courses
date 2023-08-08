package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"myrepo"
)

func main() {
	var (
		data []myrepo.Match
		rep  myrepo.MyRepository
	)

	fin, err := os.Open("hockey.json")
	if err != nil {
		return
	}
	defer fin.Close()

	b, err := rep.LoadAll(fin)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		row, err := b.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			continue
		}
		m := &myrepo.Match{}
		if err := m.ReadFrom(strings.NewReader(row)); err == nil { // игнорируем некорректные строки
			data = append(data, *m)
		}
	}
	// the data array is filled in

	over5 := filter(data, func(w myrepo.Match) bool {
		return w.Host.Goals >= w.Guest.Goals+5 || w.Guest.Goals >= w.Host.Goals+5
	})
	fmt.Println(over5)

	sort(data, func(x, y myrepo.Match) bool {
		if x.Host.Title != y.Host.Title {
			return x.Host.Title < y.Host.Title
		} else {
			return x.Guest.Title < y.Guest.Title
		}
	})

	for _, v := range data {
		fmt.Println(v)
	}

}

func filter(list []myrepo.Match, indicator func(v myrepo.Match) bool) (result []myrepo.Match) {
	for _, v := range list {
		if indicator(v) {
			result = append(result, v)
		}
	}
	return
}

func sort(list []myrepo.Match, correctorder func(x, y myrepo.Match) bool) {
	var i, j int
	for i = 1; i < len(list); i++ {
		v := list[i]
		for j = i - 1; j >= 0 && correctorder(v, list[j]); j-- {
			list[j+1] = list[j]
		}
		list[j+1] = v
	}
}
