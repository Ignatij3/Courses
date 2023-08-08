package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type point struct {
	x, y int
}

var (
	bridge_amount int
	bridge_len    int
)

func main() {
	coords := get_coordinates()
	length := calculate_length(coords)
	fmt.Printf("%.5f", length[len(length)-1][bridge_amount])
}

func get_coordinates() []point {
	var points_amount int

	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanf(reader, "%d %d %d\n", &points_amount, &bridge_amount, &bridge_len)
	coordinates := make([]point, points_amount)

	for i := 0; i < points_amount; i++ {
		fmt.Fscanf(reader, "%d %d\n", &coordinates[i].x, &coordinates[i].y)
	}

	return coordinates
}

func calculate_length(coords []point) [][]float32 {
	length := make([][]float32, len(coords))
	for i := 0; i < len(length); i++ {
		length[i] = make([]float32, bridge_amount+1)
	}

	for point_pos := 1; point_pos < len(length); point_pos++ {
		for bridge_amount := 0; bridge_amount < len(length[point_pos]); bridge_amount++ {
			s1, s2, s3 := dist(coords, length, point_pos, bridge_amount)
			if s1 <= s2 && s1 <= s3 {
				length[point_pos][bridge_amount] = s1
			} else if s2 <= s1 && s2 <= s3 {
				length[point_pos][bridge_amount] = s2
			} else {
				length[point_pos][bridge_amount] = s3
			}
		}
	}

	return length
}

func dist(coords []point, length [][]float32, point_pos, bridge_n int) (float32, float32, float32) {
	var sum1, sum2, sum3 float32 = 20001.0, 0.0, 20001.0
	var valid bool

	if bridge_n > 0 {
		s1, s2, s3 := dist(coords, length, point_pos, bridge_n-1)
		if s1 <= s2 && s1 <= s3 {
			sum1 = s1
		} else if s2 <= s1 && s2 <= s3 {
			sum1 = s2
		} else {
			sum1 = s3
		}

		var start int
		for ; start < point_pos-1; start++ {
			valid = true
			for _, c := range coords[start+1 : point_pos-1] {
				a, b := coords[start], coords[point_pos-1]
				if !((c.y-a.y)*(b.x-a.x) <= (c.x-a.x)*(b.y-a.y) && (a.x-b.x)*(a.x-b.x)+(a.y-b.y)*(a.y-b.y) <= bridge_len*bridge_len) {
					valid = false
				}
			}
			if valid {
				s1, s2, s3 := dist(coords, length, start, bridge_n-1)
				if s1 <= s2 && s1 <= s3 {
					sum3 = s1
				} else if s2 <= s1 && s2 <= s3 {
					sum3 = s2
				} else {
					sum3 = s3
				}

				sum3 += float32(math.Sqrt(float64((coords[start].x-coords[point_pos].x)*(coords[start].x-coords[point_pos].x) +
					(coords[start].y-coords[point_pos].y)*(coords[start].y-coords[point_pos].y))))
				break
			}
		}
	}

	if point_pos > 1 {
		s1, s2, s3 := dist(coords, length, point_pos-1, bridge_n)
		if s1 <= s2 && s1 <= s3 {
			sum2 = s1
		} else if s2 <= s1 && s2 <= s3 {
			sum2 = s2
		} else {
			sum2 = s3
		}

		sum2 += float32(math.Sqrt(float64((coords[point_pos-1].x-coords[point_pos].x)*(coords[point_pos-1].x-coords[point_pos].x) +
			(coords[point_pos-1].y-coords[point_pos].y)*(coords[point_pos-1].y-coords[point_pos].y))))
	} else if point_pos == 1 {
		sum2 = float32(math.Sqrt(float64((coords[point_pos-1].x-coords[point_pos].x)*(coords[point_pos-1].x-coords[point_pos].x) +
			(coords[point_pos-1].y-coords[point_pos].y)*(coords[point_pos-1].y-coords[point_pos].y))))
	}

	return sum1, sum2, sum3
}
