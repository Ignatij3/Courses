package main

import (
	"fmt"
	"bufio"
	"os"
	"encoding/json"
	"strings"
	"strconv"
)

type (
	tDate struct {
		Day string
		Month string
		Year string
	}
	tResults struct {
		Team string
		Score int
	}
	tInformation struct {
		Date tDate
		Results1 tResults
		Results2 tResults
		City string
		Draw string
	}
)

var Information []tInformation

func SportReading() {
	x := 0
	sport, err := os.Open("..//Files//Sport.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	scanner := bufio.NewScanner(sport)
	for scanner.Scan() {
		var inf tInformation
		Information = append(Information, inf)
		draw := strings.Split(scanner.Text(), ",")
		if len(draw) > 3 && draw[3] == "*" {
			Information[x].Draw = "+Draw"
		}
		date := strings.Split(scanner.Text(), "/")
			Information[x].Date.Day = date[0]
			Information[x].Date.Month = date[1]
		dateYear := strings.Split(date[2], ",")
			Information[x].Date.Year = dateYear[0]
		city := strings.Split(scanner.Text(), ",")
		cityExact := strings.Split(city[1], " ")
			Information[x].City = cityExact[0]
		res1 := strings.Split(scanner.Text(), ":")
		res1Exact := strings.Split(res1[0], " ")
			Information[x].Results1.Team = res1Exact[1]
		res2 := strings.Split(scanner.Text(), ":")
		res2Exact := strings.Split(res2[1], " ")
			Information[x].Results2.Team = res2Exact[1]
		score1 := strings.Split(scanner.Text(), ":")
		score1Exact := strings.Split(score1[1], ",")
		score1int, _ := strconv.Atoi(score1Exact[0])
			Information[x].Results1.Score = score1int
		score2 := strings.Split(scanner.Text(), ":")
		score2Exact := strings.Split(score2[2], ",")
		score2int, _ := strconv.Atoi(score2Exact[0])
			Information[x].Results2.Score = score2int
		x++
	}
	sport.Close()
}

func Transfering() (done bool) {
	Result, err := json.Marshal(Information)
	if err != nil {
		fmt.Println(err)
		return
	}
	sportS, _ := os.Create("..//Files//SportStructured.json")
	sportWrite := bufio.NewWriter(sportS)
	FWrite, _ := json.Marshal(Information)
	fmt.Fprintln(sportWrite, string(FWrite))
	sportWrite.Flush()
	fmt.Println(string(Result))
	sportS.Close()
	return true
}

func main () {
	SportReading()
	if Transfering() {
		fmt.Println("Файл успешно перебран и переписан")
	}
	fmt.Println(Information)
}
