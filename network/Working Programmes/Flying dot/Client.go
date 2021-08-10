package main

import (
	"fmt"
	"math"
	"net"
)

func initialData() (net.Conn, int, int, float64, float64, float64, error) {
	var (
		width, height int
		x, y, angle   float64
	)

	fmt.Print("Enter window width: ")
	fmt.Scan(&width)
	for width <= 0 {
		fmt.Print("Incorrect data, try again: ")
		fmt.Scan(&width)
	}

	fmt.Print("Enter window height: ")
	fmt.Scan(&height)
	for height <= 0 {
		fmt.Print("Incorrect data, try again: ")
		fmt.Scan(&height)
	}

	fmt.Print("Enter start point (horizontally): ")
	fmt.Scan(&x)
	for x < 0 || int(x) > width {
		fmt.Print("Incorrect data, try again: ")
		fmt.Scan(&x)
	}

	fmt.Print("Enter start point (vertically): ")
	fmt.Scan(&y)
	for y < 0 || int(y) > height {
		fmt.Print("Incorrect data, try again: ")
		fmt.Scan(&y)
	}

	fmt.Print("Enter flying angle 90 - right, 180 - down, 270 - left, 360 - up: ")
	fmt.Scan(&angle)
	for angle < 0 || angle > 360 {
		fmt.Print("Incorrect data, try again: ")
		fmt.Scan(&angle)
	}

	server, err := net.Dial("udp", "127.0.0.1:8081")
	if err != nil {
		fmt.Printf("Some error %v", err)
		return server, width, height, x, y, angle, err
	}

	fmt.Fprintf(server, "%d.%d", width, height)
	return server, width, height, x, y, angle, err
}

func calculate(server net.Conn, width, height int, angle, x, y float64) {
	fmt.Fprintf(server, "%d.%d", int32(x), int32(y))
	var radians float64 = (math.Pi / 180.0) * angle

	for {
		x = x + math.Cos(radians)
		y = y + math.Sin(radians)
		fmt.Fprintf(server, "%d.%d", int32(x), int32(y))

		if y <= 0 {
			angle = 360 - angle
			radians = (math.Pi / 180.0) * angle
		} else if y >= float64(height) {
			angle = 360 - angle
			radians = (math.Pi / 180.0) * angle
		} else if x <= 0 {
			angle = 180 - angle
			radians = (math.Pi / 180.0) * angle
		} else if x >= float64(width) {
			angle = 180 - angle
			radians = (math.Pi / 180.0) * angle
		}
	}
}

func main() {
	server, width, height, x, y, angle, err := initialData()
	if err != nil {
		server.Close()
		return
	}
	calculate(server, width, height, angle, x, y)
	server.Close()
}
