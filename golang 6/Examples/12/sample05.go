package main

import "fmt"

type AreaError struct {
	err    string  //error description
	Height float64 //length which caused the error
	Width  float64 //width which caused the error
}

func (e *AreaError) Error() string {
	return e.err
}

func (e *AreaError) HeightNegative() bool {
	return e.Height < 0
}

func (e *AreaError) WidthNegative() bool {
	return e.Width < 0
}

func rectArea(height, width float64) (float64, error) {
	err := ""
	if height < 0 {
		err += "length is less than zero"
	}
	if width < 0 {
		if err == "" {
			err = "width is less than zero"
		} else {
			err += ", width is less than zero"
		}
	}
	if err != "" {
		return 0, &AreaError{err, height, width}
	}
	return height * width, nil
}

func main() {
	length, width := -5.0, -9.0
	area, err := rectArea(length, width)
	if err != nil {
		if err, ok := err.(*AreaError); ok {
			if err.HeightNegative() {
				fmt.Printf("error: length %0.2f is less than zero\n", err.Height)
			}
			if err.WidthNegative() {
				fmt.Printf("error: width %0.2f is less than zero\n", err.Width)
			}
			return
		}
	}
	fmt.Println("area of rect", area)
}

// error: length -5.00 is less than zero
// error: width -9.00 is less than zero
