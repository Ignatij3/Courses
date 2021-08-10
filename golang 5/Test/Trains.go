package main

import (
	//"strconv"
	"bufio"
	"fmt"
	"os"
)

type (
	data struct {
		platform int
		time     int
	}
	platforms []data
)

func ConvToInt(str string) (int, int) {
	var nums [2]int
	fmt.Sscanf(str, "%d %d", &nums[0], &nums[1])
	return nums[0], nums[1]
}

func Calculate(nFile, trains int, freePlat []int, takenPlat platforms) {
	var (
		output             []int
		arrival, departure int
		n                  int = 1
	)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		arrival, departure = ConvToInt(scanner.Text())
	start:
		if len(takenPlat) == 0 || (len(takenPlat) > 0 && arrival <= takenPlat[0].time) {
			if len(freePlat) > 0 {
				output = append(output, freePlat[0])
				takenPlat = append(takenPlat, data{platform: freePlat[0], time: departure})
				ExtractMin(&freePlat)
				takenPlat.PushUp(len(takenPlat) - 1)
			} else {
				output = []int{0, n}
				break
			}
		} else {
			for len(takenPlat) > 0 && arrival > takenPlat[0].time {
				freePlat = append(freePlat, takenPlat[0].platform)
				PushUp(freePlat, len(freePlat)-1)
				takenPlat.ExtractMin()
			}
			goto start
		}
		n++
	}
	WriteOutput(output)
}

func WriteOutput(output []int) {
	for _, c := range output {
		fmt.Println(c)
	}
}

func ExtractMin(a *[]int) {
	(*a)[0] = (*a)[len(*a)-1]
	*a = (*a)[:len(*a)-1]
	PushDown(a, 0)
}

func PushDown(a *[]int, place int) {
	if place >= len(*a) || place < 0 {
		return
	}
	x := (*a)[place]
	for {
		if 2*place+1 >= len(*a) {
			break
		}
		minson := 2*place + 1
		rson := minson + 1
		if rson < len(*a) && (*a)[rson] < (*a)[minson] {
			minson = rson
		}
		if (*a)[minson] >= x {
			break
		}
		(*a)[place] = (*a)[minson]
		place = minson
	}
	(*a)[place] = x
}

func PushUp(a []int, place int) {
	if place >= len(a) || place <= 0 {
		return
	}
	x := a[place]
	parent := (place - 1) / 2
	for place > 0 && a[parent] > x {
		a[place] = a[parent]
		place = parent
		parent = (place - 1) / 2
	}
	a[place] = x
}

func (a *platforms) ExtractMin() {
	(*a)[0] = (*a)[len(*a)-1]
	*a = (*a)[:len(*a)-1]
	a.PushDown(0)
}

func (a *platforms) PushDown(place int) {
	if place >= len(*a) || place < 0 {
		return
	}
	x := (*a)[place]
	for {
		if 2*place+1 >= len(*a) {
			break
		}
		minson := 2*place + 1
		rson := minson + 1
		if rson < len(*a) && (*a)[rson].time < (*a)[minson].time {
			minson = rson
		}
		if (*a)[minson].time >= x.time {
			break
		}
		(*a)[place] = (*a)[minson]
		place = minson
	}
	(*a)[place] = x
}

func (a platforms) PushUp(place int) {
	if place >= len(a) || place <= 0 {
		return
	}
	x := a[place]
	parent := (place - 1) / 2
	for place > 0 && a[parent].time > x.time {
		a[place] = a[parent]
		place = parent
		parent = (place - 1) / 2
	}
	a[place] = x
}

func main() {
	var (
		deadEnds, trains, fNum int
	)
	

	fmt.Scanf("%d %d\n", &deadEnds, &trains)
	freePlat := make([]int, deadEnds)
	takenPlat := make(platforms, 0)
	for n := range freePlat {
		freePlat[n] = n + 1
	}

	Calculate(fNum, trains, freePlat, takenPlat)
}
